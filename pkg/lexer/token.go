package lexer

type TokenType int

const (
	EOF TokenType = iota
	ILLEGAL

	IDENTIFIER

	NUMBER
	STRING

	COLON
	SEMICOLON

	VARIABLE

	ADD // +
	SUB // -
	MUL // *
	DIV // /

	ASSIGN // =

	LEFT_BRACE  // (
	RIGHT_BRACE // )

	FUNCTION // function
	RETURN   // return
	IF       // if
	ELIF     // else if
	ELSE     // else
	TRUE     // true
	FALSE    // false
	FOR      // for
	BREAK    // break
	CONTINUE // continue

	//TODO: 1. () Keyword(eng)
)

var tokenTypes = []string{
	EOF:        "EOF",
	ILLEGAL:    "ILLEGAL",
	IDENTIFIER: "IDENTIFIER",
	COLON:      "COLON",
	NUMBER:     "NUMBER",
	STRING:     "STRING",
	SEMICOLON:  ";",

	VARIABLE: "VARIABLE",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	ASSIGN: "=",

	LEFT_BRACE:  "(",
	RIGHT_BRACE: ")",

	FUNCTION: "FUNCTION",
	RETURN:   "RETURN",
	IF:       "IF",
	ELIF:     "ELIF",
	ELSE:     "ELSE",
	TRUE:     "TRUE",
	FALSE:    "FALSE",
	FOR:      "FOR",
	BREAK:    "BREAK",
	CONTINUE: "CONTINUE",
}

func (t TokenType) String() string {
	return tokenTypes[t]
}

func toKeyword(str *string) TokenType {
	switch *str {
	//TODO: () Keyword(vn)
	case "biến":
		return VARIABLE
	case "hàm":
		return FUNCTION
	case "trả":
		return RETURN
	case "nếu":
		return IF
	case "nếu_thì":
		return ELIF
	case "thì":
		return ELSE
	case "đúng":
		return TRUE
	case "sai":
		return FALSE
	case "lặp":
		return FOR
	case "dừng":
		return BREAK
	case "tiếp":
		return CONTINUE
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
