package parser

import (
	"fmt"
	lex "github.com/TrGiLong/vinLang/pkg/lexer"
)

type Parser struct {
	lexer        *lex.Lexer
	currentToken lex.Token
}

func NewParser(lexer *lex.Lexer) *Parser {
	_parser := &Parser{lexer: lexer}
	_parser.currentToken = _parser.lexer.Next()
	return _parser
}

func (parser *Parser) eat(tokenType lex.TokenType) {
	if parser.currentToken.Type == tokenType {
		parser.currentToken = parser.lexer.Next()
	} else {
		panic(fmt.Sprintf("Expect %s, but get %s at %+v",
			tokenType,
			parser.currentToken.Type,
			parser.currentToken.Position,
		))
	}
}

func (parser *Parser) Parse() Node {
	return parser.statement()
}

func (parser *Parser) sequenceStatement() Node {
	return nil
}

func (parser *Parser) statement() Node {
	token := parser.currentToken

	switch token.Type {
	case lex.VARIABLE:
		return parser.varDeclaration()
	}

	return nil
}

func (parser *Parser) expression() Expr {
	if parser.currentToken.Type == lex.NUMBER {
		return parser.number()
	}
	panic(fmt.Sprintf("Invalid expression at %v", parser.currentToken.Position))
}

func (parser *Parser) number() Expr {
	pos := parser.currentToken.Position
	value := parser.currentToken.Text
	valueType := parser.currentToken.Type
	parser.eat(lex.NUMBER)

	return &BasicLit{Pos: pos, Kind: valueType, Value: value}
}

func (parser *Parser) varDeclaration() Node {
	varPos := parser.currentToken.Position
	parser.eat(lex.VARIABLE)

	identifier := &Identifier{Name: parser.currentToken.Text, Pos: parser.currentToken.Position}
	identifierPos := parser.currentToken.Position
	parser.eat(lex.IDENTIFIER)

	parser.eat(lex.COLON)

	identifierType := &Type{Name: parser.currentToken.Text, Pos: parser.currentToken.Position}
	parser.eat(lex.IDENTIFIER)

	var value Expr = nil
	if parser.currentToken.Type == lex.ASSIGN {
		parser.eat(lex.ASSIGN)
		value = parser.expression()
	}

	test := &DeclStmt{
		Pos: varPos,
		Decl: &VarDecl{
			Spec: &ValueSpec{
				Identifier: identifier,
				Type:       identifierType,
				Pos:        identifierPos,
				Value:      value,
			},
		}}
	return test
}
