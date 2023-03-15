// Copyright 2023 The Jule Programming Language.
// Use of this source code is governed by a BSD 3-Clause
// license that can be found in the LICENSE file.

package sema

import (
	"github.com/julelang/jule/ast"
	"github.com/julelang/jule/lex"
)

// Return type.
type RetType struct {
	Kind   *Type
	Idents []lex.Token
}

// Parameter.
type Param struct {
	Token    lex.Token
	Mutable  bool
	Variadic bool
	Kind     *Type
	Ident    string
}

// Function.
type Fn struct {
	Token      lex.Token
	Unsafety   bool
	Public     bool
	Cpp_linked bool
	Ident      string
	Directives []*ast.Directive
	Doc        string
	Scope      *ast.Scope
	Generics   []*ast.Generic
	Result     *RetType
	Params     []*Param

	// Function instances for each unique type combination of function call.
	// Nil if function is never used.
	Combines [][]*Type
}

// Reports whether return type is void.
func (f *Fn) Is_void() bool { return f.Result == nil }

// Function instance.
type FnIns struct {
	Decl   *Fn
	Params []*TypeKind
	Result *TypeKind
	Scope  *ast.Scope
}
