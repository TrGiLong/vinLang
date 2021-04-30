package lexer

type TokenType int

const (
	EOF TokenType = iota
	ILLEGAL
	IDENTIFIER
	INTEGER
	DOUBLE
	SEMICOLON

	ADD // +
	SUB // -
	MUL // *
	DIV // /

	ASSIGN // =

)

var tokenTypes = []string{
	EOF:       "EOF",
	ILLEGAL:   "ILLEGAL",
	IDENTIFIER: "IDENTIFIER",
	INTEGER:   "INTEGER",
	DOUBLE:    "DOUBLE",
	SEMICOLON: ";",

	// Infix ops
	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	ASSIGN: "=",
}

func (t TokenType) String() string {
	return tokenTypes[t]
}

type Token struct {
	Type     TokenType
	Position Position
	Text     string
}
