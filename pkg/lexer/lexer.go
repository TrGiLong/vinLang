package lexer

import (
	"bufio"
	"io"
	"unicode"
)

type Lexer struct {
	position Position
	reader   *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		position: Position{Line: 1, Column: 0},
		reader:   bufio.NewReader(reader),
	}
}

func (lexer *Lexer) resetPosition() {
	lexer.position.Line++
	lexer.position.Column = 0
}

func (lexer *Lexer) back() {
	if err := lexer.reader.UnreadRune(); err != nil {
		return
	}

	lexer.position.Column--
}

func (lexer *Lexer) Next() Token {
	for {
		r, _, err := lexer.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return Token{Type: EOF}
			} else {
				return Token{Type: ILLEGAL}
			}
		}

		lexer.position.Column++

		switch r {
		case '\n':
			lexer.resetPosition()
		case ';':
			return Token{Type: SEMICOLON, Position: lexer.position}
		case '+':
			return Token{Type: ADD, Position: lexer.position}
		case '-':
			return Token{Type: SUB, Position: lexer.position}
		case '*':
			return Token{Type: MUL, Position: lexer.position}
		case '/':
			return Token{Type: DIV, Position: lexer.position}
		case '=':
			return Token{Type: ASSIGN, Position: lexer.position}

		default:
			if unicode.IsSpace(r) {
				continue // nothing to do here, just move on
			} else if unicode.IsDigit(r) {
				lexer.back()
				return lexer.nextNumber()
			} else if unicode.IsLetter(r) {
				lexer.back()
				token := lexer.nextIdentifier()
				if isKeyword(&token.Text) {
					return token.toKeyword()
				}
				return token
			} else {
				return Token{Type: ILLEGAL, Position: lexer.position, Text: string(r)}
			}
		}
	}
}

func (lexer *Lexer) nextNumber() Token {
	// TODO: return token with type INTEGER and string value if it's a Integer. Or Double if it's a double

	var number string
	var dotCounter = true

	for {
		r, _, err := lexer.reader.ReadRune()
		print(string(r))
		if err != nil && err == io.EOF {
			r = 0
		}

		lexer.position.Column++
		if unicode.IsDigit(r) {
			number = number + string(r)
		} else if string(r) == "." && dotCounter {
			//Double only dot
			number = number + string(r)
			dotCounter = false
		} else {
			// scanned something not in the identifier
			lexer.back()
			//Double recognised
			if !dotCounter {
				return Token{Type: DOUBLE, Position: lexer.position, Text: number}
			}
			return Token{Type: INTEGER, Position: lexer.position, Text: number}
		}
	}
}

func (lexer *Lexer) nextIdentifier() Token {
	var identifier string

	for {
		r, _, err := lexer.reader.ReadRune()
		if err != nil && err == io.EOF {
			r = 0
		}

		lexer.position.Column++
		if unicode.IsLetter(r) {
			identifier = identifier + string(r)
		} else if unicode.IsDigit(r) {
			identifier = identifier + string(r)
		} else {
			lexer.back()
			return Token{Type: IDENTIFIER, Position: lexer.position, Text: identifier}
		}
	}
}
