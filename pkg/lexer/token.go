package lexer

type TokenType int

const (
	EOF TokenType = iota
	ILLEGAL

	IDENTIFIER

	INTEGER
	DOUBLE

	SEMICOLON

	VARIABLE

	ADD // +
	SUB // -
	MUL // *
	DIV // /

	ASSIGN // =

)

var tokenTypes = []string{
	EOF:        "EOF",
	ILLEGAL:    "ILLEGAL",
	IDENTIFIER: "IDENTIFIER",
	INTEGER:    "INTEGER",
	DOUBLE:     "DOUBLE",
	SEMICOLON:  ";",

	VARIABLE: "VARIABLE",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	ASSIGN: "=",
}

func (t TokenType) String() string {
	return tokenTypes[t]
}

func toKeyword(str *string) TokenType {
	switch *str {
	case "biến":
		return VARIABLE
	default:
		return ILLEGAL
	}
}

func isKeyword(str *string) bool {
	return toKeyword(str) != ILLEGAL
}

type Token struct {
	Type     TokenType
	Position Position
	Text     string
}

func (token *Token) toKeyword() Token {
	return Token{Type: toKeyword(&token.Text)}
}
