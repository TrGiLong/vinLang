package parser

import (
	lexer "github.com/TrGiLong/vinLang/pkg/lexer"
)

type (
	Node interface {
		Position() lexer.Position
	}

	Expr interface {
		Node
		exprNode()
	}

	Stml interface {
		Node
		stmtNode()
	}

	Decl interface {
		Node
		declNode()
	}

	Spec interface {
		Node
		specNode()
	}
)

type (
	AssignStmt struct {
		Pos lexer.Position
		Lhs []Expr
		Rhs []Expr
	}

	DeclStmt struct {
		Decl Decl
	}
)

func (x *AssignStmt) Position() lexer.Position { return x.Pos }
func (x *DeclStmt) Position() lexer.Position   { return x.Decl.Position() }

func (x *AssignStmt) stmtNode() {}
func (x *DeclStmt) stmtNode()   {}

type (
	BinaryExpr struct {
		Pos lexer.Position
		Op  lexer.TokenType
		Lhs Expr
		Rhs Expr
	}

	Identifier struct {
		Pos  lexer.Position
		Name string
		//TODO: need declaration
	}

	// A BasicLit node represents a literal of basic type.
	BasicLit struct {
		Pos   lexer.Position
		Kind  lexer.TokenType
		Value string
	}
)

func (x *BinaryExpr) Position() lexer.Position { return x.Pos }
func (x *Identifier) Position() lexer.Position { return x.Pos }
func (x *BasicLit) Position() lexer.Position   { return x.Pos }
func (x *BinaryExpr) exprNode()                {}
func (x *Identifier) exprNode()                {}
func (x *BasicLit) exprNode()                  {}

type (
	VarDecl struct {
		Pos  lexer.Position
		Spec Spec
	}
)

func (x *VarDecl) Position() lexer.Position { return x.Pos }
func (*VarDecl) declNode()                  {}

type (
	ValueSpec struct {
		Pos   lexer.Position
		Identifier *Identifier
		Value Expr
	}
)

func (x *ValueSpec) Position() lexer.Position { return x.Pos }
func (*ValueSpec) specNode()                  {}
