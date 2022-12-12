package models

import "github.com/julelang/jule/lex"

// Namespace is the AST model of namespace statements.
type Namespace struct {
	Token   lex.Token
	Id      string
	Tree    []Object
	Defines *Defmap
}
