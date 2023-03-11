package sema

import (
	"github.com/julelang/jule/ast"
	"github.com/julelang/jule/lex"
)

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
	Result     *ast.RetType
	Params     []*ast.Param

	// Type combinations of generic function.
	// Nil or len() = 0 if never invoked.
	Combines [][]*ast.Type
}
