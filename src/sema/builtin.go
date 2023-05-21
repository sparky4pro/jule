package sema

import (
	"github.com/julelang/jule/ast"
	"github.com/julelang/jule/types"
)

// Type alias for built-in function callers.
//
// Parameters;
//  e: Caller owner Eval instance.
//  fc: Function call expression.
//  d: Data instance for evaluated expression of function.
type _BuiltinCaller = func(e *_Eval, fc *ast.FnCallExpr, d *Data) *Data

var builtin_fn_out = &FnIns{}
var builtin_fn_outln = &FnIns{}
var builtin_fn_new = &FnIns{}
var builtin_fn_drop = &FnIns{}
var builtin_fn_panic = &FnIns{}
var builtin_fn_make = &FnIns{}
var builtin_fn_append = &FnIns{}

var builtin_fn_real = &FnIns{
	Result: &TypeKind{kind: build_prim_type(types.TypeKind_BOOL)},
}

var builtin_fn_copy = &FnIns{
	Result: &TypeKind{kind: build_prim_type(types.TypeKind_INT)},
}

func init() {
	builtin_fn_out.Caller = builtin_caller_out
	builtin_fn_outln.Caller = builtin_caller_outln
	builtin_fn_new.Caller = builtin_caller_new
	builtin_fn_real.Caller = builtin_caller_real
	builtin_fn_drop.Caller = builtin_caller_drop
	builtin_fn_panic.Caller = builtin_caller_panic
	builtin_fn_make.Caller = builtin_caller_make
	builtin_fn_append.Caller = builtin_caller_append
	builtin_fn_copy.Caller = builtin_caller_copy
}

func get_builtin_def(ident string) any {
	switch ident {
	case "out":
		return builtin_fn_out

	case "outln":
		return builtin_fn_outln

	case "new":
		return builtin_fn_new

	case "real":
		return builtin_fn_real

	case "drop":
		return builtin_fn_drop

	case "panic":
		return builtin_fn_panic

	case "make":
		return builtin_fn_make

	case "append":
		return builtin_fn_append

	case "copy":
		return builtin_fn_copy

	default:
		return nil
	}
}

func builtin_caller_common(e *_Eval, fc *ast.FnCallExpr, d *Data) *Data {
	f := d.Kind.Fnc()

	fcac := _FnCallArgChecker{
		e:                  e,
		f:                  f,
		args:               fc.Args,
		dynamic_annotation: false,
		error_token:        fc.Token,
	}
	_ = fcac.check()

	model := &FnCallExprModel{
		Func: f,
		IsCo: fc.Concurrent,
		Expr: d.Model,
		Args: fcac.arg_models,
	}

	if f.Result == nil {
		d = build_void_data()
	} else {
		d = &Data{
			Kind: f.Result,
		}
	}

	d.Model = model
	return d
}

func builtin_caller_out(e *_Eval, fc *ast.FnCallExpr, _ *Data) *Data {
	if len(fc.Args) < 1 {
		e.push_err(fc.Token, "missing_expr_for", "v")
		return nil
	}
	if len(fc.Args) > 1 {
		e.push_err(fc.Args[2].Token, "argument_overflow")
	}

	expr := e.eval_expr(fc.Args[0])
	if expr == nil {
		return nil
	}

	if expr.Kind.Fnc() != nil {
		e.push_err(fc.Args[0].Token, "invalid_expr")
		return nil
	}

	d := build_void_data()
	d.Model = &BuiltinOutCallExprModel{Expr: expr.Model}
	return d
}

func builtin_caller_outln(e *_Eval, fc *ast.FnCallExpr, _ *Data) *Data {
	d := builtin_caller_out(e, fc, nil)
	if d == nil {
		return nil
	}

	d.Model = &BuiltinOutlnCallExprModel{
		Expr: d.Model.(*BuiltinOutCallExprModel).Expr,
	}
	return d
}

func builtin_caller_new(e *_Eval, fc *ast.FnCallExpr, d *Data) *Data {
	if len(fc.Args) < 1 {
		e.push_err(fc.Token, "missing_expr_for", "type")
		return nil
	}
	if len(fc.Args) > 2 {
		e.push_err(fc.Args[2].Token, "argument_overflow")
	}

	t := e.eval_expr_kind(fc.Args[0].Kind)
	if t == nil {
		return nil
	}

	if !t.Decl {
		e.push_err(fc.Args[0].Token, "invalid_type")
		return nil
	}

	if !is_valid_for_ref(t.Kind) {
		e.push_err(fc.Args[0].Token, "invalid_type")
		return nil
	}

	d.Kind = &TypeKind{kind: &Ref{Elem: t.Kind}}


	if len(fc.Args) == 2 { // Initialize expression.
		init := e.s.evalp(fc.Args[1], e.lookup, &TypeSymbol{Kind: t.Kind})
		if init != nil {
			e.s.check_assign_type(t.Kind, init, fc.Args[1].Token, false)
			d.Model = &BuiltinNewCallExprModel{
				Kind: t.Kind,
				Init: init.Model,
			}
		}
	} else {
		d.Model = &BuiltinNewCallExprModel{Kind: t.Kind}
	}

	return d
}

func builtin_caller_real(e *_Eval, fc *ast.FnCallExpr, d *Data) *Data {
	if len(fc.Args) < 1 {
		e.push_err(fc.Token, "missing_expr_for", "ref")
		return nil
	}
	if len(fc.Args) > 1 {
		e.push_err(fc.Args[2].Token, "argument_overflow")
	}

	ref := e.eval_expr(fc.Args[0])
	if ref == nil {
		return nil
	}

	if ref.Kind.Ref() == nil {
		e.push_err(fc.Args[0].Token, "invalid_expr")
		return nil
	}

	d.Kind = builtin_fn_real.Result
	d.Model = &BuiltinRealCallExprModel{Expr: ref.Model}
	return d
}

func builtin_caller_drop(e *_Eval, fc *ast.FnCallExpr, _ *Data) *Data {
	if len(fc.Args) < 1 {
		e.push_err(fc.Token, "missing_expr_for", "ref")
		return nil
	}
	if len(fc.Args) > 1 {
		e.push_err(fc.Args[2].Token, "argument_overflow")
	}

	ref := e.eval_expr(fc.Args[0])
	if ref == nil {
		return nil
	}

	if ref.Kind.Ref() == nil {
		e.push_err(fc.Args[0].Token, "invalid_expr")
		return nil
	}

	d := build_void_data()
	d.Model = &BuiltinDropCallExprModel{Expr: ref.Model}
	return d
}

func builtin_caller_panic(e *_Eval, fc *ast.FnCallExpr, _ *Data) *Data {
	if len(fc.Args) < 1 {
		e.push_err(fc.Token, "missing_expr_for", "error")
		return nil
	}
	if len(fc.Args) > 1 {
		e.push_err(fc.Args[2].Token, "argument_overflow")
	}

	expr := e.eval_expr(fc.Args[0])
	if expr == nil {
		return nil
	}

	d := build_void_data()
	d.Model = &BuiltinPanicCallExprModel{Expr: expr.Model}
	return d
}

func builtin_caller_make(e *_Eval, fc *ast.FnCallExpr, d *Data) *Data {
	if len(fc.Args) < 2 {
		if len(fc.Args) == 1 {
			e.push_err(fc.Token, "missing_expr_for", "size")
			return nil
		}
		e.push_err(fc.Token, "missing_expr_for", "type, and size")
		return nil
	}
	if len(fc.Args) > 2 {
		e.push_err(fc.Args[2].Token, "argument_overflow")
	}

	t := e.eval_expr_kind(fc.Args[0].Kind)
	if t == nil {
		return nil
	}

	if !t.Decl || t.Kind.Slc() == nil {
		e.push_err(fc.Args[0].Token, "invalid_type")
		return nil
	}

	d.Kind = t.Kind
	
	size := e.s.evalp(fc.Args[1], e.lookup, &TypeSymbol{Kind: t.Kind})
	if size == nil {
		return d
	}
	
	e.check_integer_indexing_by_data(size, fc.Args[1].Token)

	// Ignore size expression if size is constant zero.
	if size.Is_const() && size.Constant.As_i64() == 0 {
		size.Model = nil
	}

	d.Model = &BuiltinMakeCallExprModel{
		Kind: t.Kind,
		Size: size.Model,
	}

	return d
}

func builtin_caller_append(e *_Eval, fc *ast.FnCallExpr, d *Data) *Data {
	if len(fc.Args) < 2 {
		if len(fc.Args) == 1 {
			e.push_err(fc.Token, "missing_expr_for", "src")
			return nil
		}
		e.push_err(fc.Token, "missing_expr_for", "src, and values")
		return nil
	}

	t := e.eval_expr(fc.Args[0])
	if t == nil {
		return nil
	}

	if t.Kind.Slc() == nil {
		e.push_err(fc.Args[0].Token, "invalid_expr")
		return nil
	}

	f := &FnIns{
		Params: []*ParamIns{
			{
				Decl: &Param{},
				Kind: t.Kind,
			},
			{
				Decl: &Param{
					Variadic: true,
				},
				Kind: t.Kind.Slc().Elem,
			},
		},
		Result: t.Kind,
	}
	d.Kind = &TypeKind{kind: f}
	d.Model = &CommonIdentExprModel{Ident: "_append"}

	d = builtin_caller_common(e, fc, d)
	return d
}

func builtin_caller_copy(e *_Eval, fc *ast.FnCallExpr, d *Data) *Data {
	if len(fc.Args) < 2 {
		if len(fc.Args) == 1 {
			e.push_err(fc.Token, "missing_expr_for", "src")
			return nil
		}
		e.push_err(fc.Token, "missing_expr_for", "src, and values")
		return nil
	}

	t := e.eval_expr(fc.Args[0])
	if t == nil {
		return nil
	}

	if t.Kind.Slc() == nil {
		e.push_err(fc.Args[0].Token, "invalid_expr")
		return nil
	}

	f := &FnIns{
		Params: []*ParamIns{
			{
				Decl: &Param{},
				Kind: t.Kind,
			},
			{
				Decl: &Param{},
				Kind: t.Kind,
			},
		},
		Result: builtin_fn_copy.Result,
	}

	d.Kind = &TypeKind{kind: f}
	d.Model = &CommonIdentExprModel{Ident: "_copy"}

	d = builtin_caller_common(e, fc, d)
	return d
}
