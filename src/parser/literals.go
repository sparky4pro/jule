package parser

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/julelang/jule/ast"
	"github.com/julelang/jule/build"
	"github.com/julelang/jule/lex"
	"github.com/julelang/jule/types"
)

type literal_eval struct {
	token lex.Token
	model *exprModel
	p     *Parser
}

func strModel(v value) ast.ExprModel {
	content := v.expr.(string)
	if lex.IsRawStr(content) {
		return exprNode{ToRawStrLiteral([]byte(content))}
	}
	return exprNode{ToStrLiteral([]byte(content))}
}

func boolModel(v value) ast.ExprModel {
	if v.expr.(bool) {
		return exprNode{lex.KND_TRUE}
	}
	return exprNode{lex.KND_FALSE}
}

func getModel(v value) ast.ExprModel {
	switch v.expr.(type) {
	case string:
		return strModel(v)
	case bool:
		return boolModel(v)
	default:
		return numericModel(v)
	}
}

func numericModel(v value) ast.ExprModel {
	switch t := v.expr.(type) {
	case uint64:
		fmt := strconv.FormatUint(t, 10)
		return exprNode{fmt + "LLU"}
	case int64:
		fmt := strconv.FormatInt(t, 10)
		return exprNode{fmt + "LL"}
	case float64:
		switch {
		case normalize(&v):
			return numericModel(v)
		case v.data.DataType.Id == types.F32:
			return exprNode{fmt.Sprint(t) + "f"}
		case v.data.DataType.Id == types.F64:
			return exprNode{fmt.Sprint(t)}
		}
	}
	return nil
}

func (ve *literal_eval) str() value {
	var v value
	v.constant = true
	v.data.Value = ve.token.Kind
	v.data.DataType.Id = types.STR
	v.data.DataType.Kind = types.TYPE_MAP[v.data.DataType.Id]
	content := ve.token.Kind[1 : len(ve.token.Kind)-1]
	v.expr = content
	v.model = strModel(v)
	ve.model.append_sub(v.model)
	return v
}

func toByteLiteral(kind string) (string, bool) {
	kind = kind[1 : len(kind)-1]
	isByte := false
	switch {
	case len(kind) == 1 && kind[0] <= 255:
		isByte = true
	case kind[0] == '\\' && kind[1] == 'x':
		isByte = true
	case kind[0] == '\\' && kind[1] >= '0' && kind[1] <= '7':
		isByte = true
	}
	return kind, isByte
}

func (ve *literal_eval) rune() value {
	var v value
	v.constant = true
	v.data.Value = ve.token.Kind
	content, isByte := toByteLiteral(ve.token.Kind)
	if isByte {
		v.data.DataType.Id = types.U8
	} else { // rune
		v.data.DataType.Id = types.I32
	}
	content = ToRuneLiteral([]byte(content))
	v.data.DataType.Kind = types.TYPE_MAP[v.data.DataType.Id]
	v.expr, _ = strconv.ParseInt(content[2:], 16, 64)
	v.model = exprNode{content}
	ve.model.append_sub(v.model)
	return v
}

func (ve *literal_eval) bool() value {
	var v value
	v.constant = true
	v.data.Value = ve.token.Kind
	v.data.DataType.Id = types.BOOL
	v.data.DataType.Kind = types.TYPE_MAP[v.data.DataType.Id]
	v.expr = ve.token.Kind == lex.KND_TRUE
	v.model = boolModel(v)
	ve.model.append_sub(v.model)
	return v
}

func (ve *literal_eval) nil() value {
	var v value
	v.constant = true
	v.data.Value = ve.token.Kind
	v.data.DataType.Id = types.NIL
	v.data.DataType.Kind = types.TYPE_MAP[v.data.DataType.Id]
	v.expr = nil
	v.model = exprNode{ve.token.Kind}
	ve.model.append_sub(v.model)
	return v
}

func normalize(v *value) (normalized bool) {
	switch {
	case !v.constant:
		return
	case int_assignable(types.U64, *v):
		v.data.DataType.Id = types.U64
		v.data.DataType.Kind = types.TYPE_MAP[v.data.DataType.Id]
		v.expr = tonumu(v.expr)
		bitize(v)
		return true
	case int_assignable(types.I64, *v):
		v.data.DataType.Id = types.I64
		v.data.DataType.Kind = types.TYPE_MAP[v.data.DataType.Id]
		v.expr = tonums(v.expr)
		bitize(v)
		return true
	}
	return
}

func (ve *literal_eval) float() value {
	var v value
	v.data.Value = ve.token.Kind
	v.data.DataType.Id = types.F64
	v.data.DataType.Kind = types.TYPE_MAP[v.data.DataType.Id]
	v.expr, _ = strconv.ParseFloat(v.data.Value, 64)
	return v
}

func (ve *literal_eval) integer() value {
	var v value
	v.data.Value = ve.token.Kind
	var bigint big.Int
	switch {
	case strings.HasPrefix(ve.token.Kind, "0x"):
		_, _ = bigint.SetString(ve.token.Kind[2:], 16)
	case strings.HasPrefix(ve.token.Kind, "0b"):
		_, _ = bigint.SetString(ve.token.Kind[2:], 2)
	case ve.token.Kind[0] == '0':
		_, _ = bigint.SetString(ve.token.Kind[1:], 8)
	default:
		_, _ = bigint.SetString(ve.token.Kind, 10)
	}
	if bigint.IsInt64() {
		v.expr = bigint.Int64()
	} else {
		v.expr = bigint.Uint64()
	}
	bitize(&v)
	return v
}

func (ve *literal_eval) numeric() value {
	var v value
	if lex.IsFloat(ve.token.Kind) {
		v = ve.float()
	} else {
		v = ve.integer()
	}
	v.constant = true
	v.model = numericModel(v)
	ve.model.append_sub(v.model)
	return v
}

func make_value_from_var(v *Var) (val value) {
	val.data.Value = v.Id
	val.data.DataType = v.DataType
	val.constant = v.Constant
	val.data.Token = v.Token
	val.lvalue = !val.constant
	val.mutable = v.Mutable
	if val.constant {
		val.expr = v.ExprTag
		val.model = v.Expr.Model
	}
	return
}

func (ve *literal_eval) var_id(id string, variable *Var, global bool) (v value) {
	variable.Used = true
	v = make_value_from_var(variable)
	if v.constant {
		ve.model.append_sub(v.model)
	} else {
		if variable.Id == lex.KND_SELF && !types.IsRef(variable.DataType) {
			ve.model.append_sub(exprNode{"(*this)"})
		} else {
			ve.model.append_sub(exprNode{variable.OutId()})
		}
		ve.p.eval.has_error = ve.p.eval.has_error || types.IsVoid(v.data.DataType)
	}
	return
}

func make_value_from_fn(f *ast.Fn) (v value) {
	v.data.Value = f.Id
	v.data.DataType.Id = types.FN
	v.data.DataType.Tag = f
	v.data.DataType.Kind = f.TypeKind()
	v.data.Token = f.Token
	return
}

func (ve *literal_eval) fn_id(id string, f *Fn) (v value) {
	f.Used = true
	v = make_value_from_fn(f)
	ve.model.append_sub(exprNode{f.OutId()})
	return
}

func (ve *literal_eval) enumId(id string, e *Enum) (v value) {
	e.Used = true
	v.data.Value = id
	v.data.DataType.Id = types.ENUM
	v.data.DataType.Kind = e.Id
	v.data.DataType.Tag = e
	v.data.Token = e.Token
	v.constant = true
	v.is_type = true
	// If built-in.
	if e.Token.Id == lex.ID_NA {
		ve.model.append_sub(exprNode{build.OutId(id, 0)})
	} else {
		ve.model.append_sub(exprNode{build.OutId(id, e.Token.File.Addr())})
	}
	return
}

func make_value_from_struct(s *Struct) (v value) {
	v.data.Value = s.Id
	v.data.DataType.Id = types.STRUCT
	v.data.DataType.Tag = s
	v.data.DataType.Kind = s.Id
	v.data.DataType.Token = s.Token
	v.data.Token = s.Token
	v.is_type = true
	return
}

func (ve *literal_eval) struct_id(id string, s *Struct) (v value) {
	s.Used = true
	v = make_value_from_struct(s)
	// If builtin.
	if s.Token.Id == lex.ID_NA {
		ve.model.append_sub(exprNode{build.OutId(id, 0)})
	} else {
		ve.model.append_sub(exprNode{build.OutId(id, s.Token.File.Addr())})
	}
	return
}

func (ve *literal_eval) type_id(id string, t *TypeAlias) (_ value, _ bool) {
	dt, ok := ve.p.realType(t.TargetType, true)
	if !ok {
		return
	}
	if types.IsStruct(dt) {
		return ve.struct_id(id, dt.Tag.(*Struct)), true
	}
	return
}

func (ve *literal_eval) id() (_ value, ok bool) {
	id := ve.token.Kind

	v, _ := ve.p.block_var_by_id(id)
	if v != nil {
		return ve.var_id(id, v, false), true
	} else {
		v, _, _ := ve.p.global_by_id(id)
		if v != nil {
			return ve.var_id(id, v, true), true
		}
	}

	f, _, _ := ve.p.fn_by_id(id)
	if f != nil {
		return ve.fn_id(id, f), true
	}

	e, _, _ := ve.p.enum_by_id(id)
	if e != nil {
		return ve.enumId(id, e), true
	}

	s, _, _ := ve.p.struct_by_id(id)
	if s != nil {
		return ve.struct_id(id, s), true
	}

	t, _, _ := ve.p.type_by_id(id)
	if t != nil {
		return ve.type_id(id, t)
	}

	ve.p.eval.pusherrtok(ve.token, "id_not_exist", id)
	return
}
