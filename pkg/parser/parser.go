package parser

import (
	"fmt"
	lex "github.com/TrGiLong/vinLang/pkg/lexer"
)

type Parser struct {
	lexer        lex.Lexer
	currentToken lex.Token
}

func (parser *Parser) eat(tokenType lex.TokenType) {
	if parser.currentToken.Type == tokenType {
		parser.currentToken = parser.lexer.Next()
	} else {
		panic(fmt.Sprintf("Expect %s, but get %s", tokenType, parser.currentToken.Type))
	}
	parser.currentToken = parser.lexer.Next()
}

func (parser *Parser) parse() Node {
	parser.currentToken = parser.lexer.Next()
	return parser.statement()
}

func (parser *Parser) sequenceStatement() Node {
	return nil
}

func (parser *Parser) statement() Node {
	token := parser.currentToken

	switch token.Type {
	case lex.VARIABLE:
		return parser.declaration()
	}

	return nil
}

func (parser *Parser) declaration() Node {
	varPos := parser.currentToken.Position
	parser.eat(lex.VARIABLE)

	identifier := Identifier{Name: parser.currentToken.Text, Pos: parser.currentToken.Position}
	parser.eat(lex.IDENTIFIER)

	type :=

	test := &DeclStmt{Decl: &VarDecl{Spec: &ValueSpec{Identifier: &identifier, Pos: varPos}}}
	return test
}
