// Code generated from ./antlr/vinLang.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 39, 212,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 7, 34, 181, 10, 34, 12, 34, 14,
	34, 184, 11, 34, 3, 34, 6, 34, 187, 10, 34, 13, 34, 14, 34, 188, 3, 34,
	3, 34, 6, 34, 193, 10, 34, 13, 34, 14, 34, 194, 5, 34, 197, 10, 34, 3,
	35, 3, 35, 3, 36, 6, 36, 202, 10, 36, 13, 36, 14, 36, 203, 3, 37, 5, 37,
	207, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 2, 2, 39, 3, 3, 5, 4, 7, 5, 9,
	6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15,
	29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24,
	47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33,
	65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 3, 2, 6, 4, 2, 39, 39,
	49, 49, 3, 2, 50, 59, 24, 2, 50, 59, 67, 92, 99, 124, 194, 197, 202, 204,
	206, 207, 212, 215, 219, 220, 223, 223, 226, 229, 234, 236, 238, 239, 244,
	247, 251, 252, 255, 255, 260, 261, 274, 275, 298, 299, 362, 363, 418, 419,
	433, 434, 7842, 7931, 5, 2, 11, 12, 15, 15, 34, 34, 2, 216, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2,
	2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2,
	2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3,
	2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73,
	3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 3, 77, 3, 2, 2, 2, 5, 79, 3, 2, 2, 2, 7,
	84, 3, 2, 2, 2, 9, 86, 3, 2, 2, 2, 11, 88, 3, 2, 2, 2, 13, 90, 3, 2, 2,
	2, 15, 92, 3, 2, 2, 2, 17, 94, 3, 2, 2, 2, 19, 96, 3, 2, 2, 2, 21, 98,
	3, 2, 2, 2, 23, 100, 3, 2, 2, 2, 25, 103, 3, 2, 2, 2, 27, 105, 3, 2, 2,
	2, 29, 107, 3, 2, 2, 2, 31, 110, 3, 2, 2, 2, 33, 113, 3, 2, 2, 2, 35, 116,
	3, 2, 2, 2, 37, 119, 3, 2, 2, 2, 39, 122, 3, 2, 2, 2, 41, 124, 3, 2, 2,
	2, 43, 126, 3, 2, 2, 2, 45, 128, 3, 2, 2, 2, 47, 130, 3, 2, 2, 2, 49, 134,
	3, 2, 2, 2, 51, 139, 3, 2, 2, 2, 53, 144, 3, 2, 2, 2, 55, 148, 3, 2, 2,
	2, 57, 152, 3, 2, 2, 2, 59, 156, 3, 2, 2, 2, 61, 166, 3, 2, 2, 2, 63, 171,
	3, 2, 2, 2, 65, 175, 3, 2, 2, 2, 67, 182, 3, 2, 2, 2, 69, 198, 3, 2, 2,
	2, 71, 201, 3, 2, 2, 2, 73, 206, 3, 2, 2, 2, 75, 208, 3, 2, 2, 2, 77, 78,
	7, 46, 2, 2, 78, 4, 3, 2, 2, 2, 79, 80, 7, 100, 2, 2, 80, 81, 7, 107, 2,
	2, 81, 82, 7, 7873, 2, 2, 82, 83, 7, 112, 2, 2, 83, 6, 3, 2, 2, 2, 84,
	85, 7, 60, 2, 2, 85, 8, 3, 2, 2, 2, 86, 87, 7, 61, 2, 2, 87, 10, 3, 2,
	2, 2, 88, 89, 7, 63, 2, 2, 89, 12, 3, 2, 2, 2, 90, 91, 7, 45, 2, 2, 91,
	14, 3, 2, 2, 2, 92, 93, 7, 47, 2, 2, 93, 16, 3, 2, 2, 2, 94, 95, 7, 44,
	2, 2, 95, 18, 3, 2, 2, 2, 96, 97, 9, 2, 2, 2, 97, 20, 3, 2, 2, 2, 98, 99,
	7, 35, 2, 2, 99, 22, 3, 2, 2, 2, 100, 101, 7, 62, 2, 2, 101, 102, 7, 63,
	2, 2, 102, 24, 3, 2, 2, 2, 103, 104, 7, 62, 2, 2, 104, 26, 3, 2, 2, 2,
	105, 106, 7, 64, 2, 2, 106, 28, 3, 2, 2, 2, 107, 108, 7, 64, 2, 2, 108,
	109, 7, 63, 2, 2, 109, 30, 3, 2, 2, 2, 110, 111, 7, 63, 2, 2, 111, 112,
	7, 63, 2, 2, 112, 32, 3, 2, 2, 2, 113, 114, 7, 35, 2, 2, 114, 115, 7, 63,
	2, 2, 115, 34, 3, 2, 2, 2, 116, 117, 7, 45, 2, 2, 117, 118, 7, 45, 2, 2,
	118, 36, 3, 2, 2, 2, 119, 120, 7, 47, 2, 2, 120, 121, 7, 47, 2, 2, 121,
	38, 3, 2, 2, 2, 122, 123, 7, 125, 2, 2, 123, 40, 3, 2, 2, 2, 124, 125,
	7, 127, 2, 2, 125, 42, 3, 2, 2, 2, 126, 127, 7, 42, 2, 2, 127, 44, 3, 2,
	2, 2, 128, 129, 7, 43, 2, 2, 129, 46, 3, 2, 2, 2, 130, 131, 7, 110, 2,
	2, 131, 132, 7, 7865, 2, 2, 132, 133, 7, 114, 2, 2, 133, 48, 3, 2, 2, 2,
	134, 135, 7, 118, 2, 2, 135, 136, 7, 107, 2, 2, 136, 137, 7, 7873, 2, 2,
	137, 138, 7, 114, 2, 2, 138, 50, 3, 2, 2, 2, 139, 140, 7, 102, 2, 2, 140,
	141, 7, 7917, 2, 2, 141, 142, 7, 112, 2, 2, 142, 143, 7, 105, 2, 2, 143,
	52, 3, 2, 2, 2, 144, 145, 7, 118, 2, 2, 145, 146, 7, 116, 2, 2, 146, 147,
	7, 7845, 2, 2, 147, 54, 3, 2, 2, 2, 148, 149, 7, 112, 2, 2, 149, 150, 7,
	7873, 2, 2, 150, 151, 7, 119, 2, 2, 151, 56, 3, 2, 2, 2, 152, 153, 7, 118,
	2, 2, 153, 154, 7, 106, 2, 2, 154, 155, 7, 238, 2, 2, 155, 58, 3, 2, 2,
	2, 156, 157, 7, 109, 2, 2, 157, 158, 7, 106, 2, 2, 158, 159, 7, 246, 2,
	2, 159, 160, 7, 112, 2, 2, 160, 161, 7, 105, 2, 2, 161, 162, 7, 97, 2,
	2, 162, 163, 7, 118, 2, 2, 163, 164, 7, 106, 2, 2, 164, 165, 7, 238, 2,
	2, 165, 60, 3, 2, 2, 2, 166, 167, 7, 275, 2, 2, 167, 168, 7, 252, 2, 2,
	168, 169, 7, 112, 2, 2, 169, 170, 7, 105, 2, 2, 170, 62, 3, 2, 2, 2, 171,
	172, 7, 117, 2, 2, 172, 173, 7, 99, 2, 2, 173, 174, 7, 107, 2, 2, 174,
	64, 3, 2, 2, 2, 175, 176, 7, 106, 2, 2, 176, 177, 7, 226, 2, 2, 177, 178,
	7, 111, 2, 2, 178, 66, 3, 2, 2, 2, 179, 181, 7, 47, 2, 2, 180, 179, 3,
	2, 2, 2, 181, 184, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2,
	2, 183, 186, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 185, 187, 5, 69, 35, 2,
	186, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 188,
	189, 3, 2, 2, 2, 189, 196, 3, 2, 2, 2, 190, 192, 7, 48, 2, 2, 191, 193,
	5, 69, 35, 2, 192, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 192, 3,
	2, 2, 2, 194, 195, 3, 2, 2, 2, 195, 197, 3, 2, 2, 2, 196, 190, 3, 2, 2,
	2, 196, 197, 3, 2, 2, 2, 197, 68, 3, 2, 2, 2, 198, 199, 9, 3, 2, 2, 199,
	70, 3, 2, 2, 2, 200, 202, 5, 73, 37, 2, 201, 200, 3, 2, 2, 2, 202, 203,
	3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 72, 3, 2,
	2, 2, 205, 207, 9, 4, 2, 2, 206, 205, 3, 2, 2, 2, 207, 74, 3, 2, 2, 2,
	208, 209, 9, 5, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 8, 38, 2, 2, 211,
	76, 3, 2, 2, 2, 9, 2, 182, 188, 194, 196, 203, 206, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'bi\u1EBFn'", "':'", "';'", "'='", "'+'", "'-'", "'*'", "",
	"'!'", "'<='", "'<'", "'>'", "'>='", "'=='", "'!='", "'++'", "'--'", "'{'",
	"'}'", "'('", "')'", "'l\u1EB7p'", "'ti\u1EBFp'", "'d\u1EEBng'", "'tr\u1EA3'",
	"'n\u1EBFu'", "'th\u00EC'", "'kh\u00F4ng_th\u00EC'", "'\u0111\u00FAng'",
	"'sai'", "'h\u00E0m'",
}

var lexerSymbolicNames = []string{
	"", "", "VAR", "COLON", "SEMICOLIN", "ASSIGN", "ADD", "SUB", "MUL", "DIV",
	"NEG", "LE", "L", "G", "GE", "E", "NE", "INCREMENT", "DECREMENT", "LEFT_BRACE",
	"RIGHT_BRACE", "LEFT_PARENTHESE", "RIGHT_PARENTHESE", "FOR", "CONTINUE",
	"BREAK", "RETURN", "IF", "ELSE", "ELIF", "TRUE", "FALSE", "FUNCTION", "NUMBER",
	"DIGIT", "ID", "CHAR", "SPACE",
}

var lexerRuleNames = []string{
	"T__0", "VAR", "COLON", "SEMICOLIN", "ASSIGN", "ADD", "SUB", "MUL", "DIV",
	"NEG", "LE", "L", "G", "GE", "E", "NE", "INCREMENT", "DECREMENT", "LEFT_BRACE",
	"RIGHT_BRACE", "LEFT_PARENTHESE", "RIGHT_PARENTHESE", "FOR", "CONTINUE",
	"BREAK", "RETURN", "IF", "ELSE", "ELIF", "TRUE", "FALSE", "FUNCTION", "NUMBER",
	"DIGIT", "ID", "CHAR", "SPACE",
}

type vinLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewvinLangLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *vinLangLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewvinLangLexer(input antlr.CharStream) *vinLangLexer {
	l := new(vinLangLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "vinLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// vinLangLexer tokens.
const (
	vinLangLexerT__0             = 1
	vinLangLexerVAR              = 2
	vinLangLexerCOLON            = 3
	vinLangLexerSEMICOLIN        = 4
	vinLangLexerASSIGN           = 5
	vinLangLexerADD              = 6
	vinLangLexerSUB              = 7
	vinLangLexerMUL              = 8
	vinLangLexerDIV              = 9
	vinLangLexerNEG              = 10
	vinLangLexerLE               = 11
	vinLangLexerL                = 12
	vinLangLexerG                = 13
	vinLangLexerGE               = 14
	vinLangLexerE                = 15
	vinLangLexerNE               = 16
	vinLangLexerINCREMENT        = 17
	vinLangLexerDECREMENT        = 18
	vinLangLexerLEFT_BRACE       = 19
	vinLangLexerRIGHT_BRACE      = 20
	vinLangLexerLEFT_PARENTHESE  = 21
	vinLangLexerRIGHT_PARENTHESE = 22
	vinLangLexerFOR              = 23
	vinLangLexerCONTINUE         = 24
	vinLangLexerBREAK            = 25
	vinLangLexerRETURN           = 26
	vinLangLexerIF               = 27
	vinLangLexerELSE             = 28
	vinLangLexerELIF             = 29
	vinLangLexerTRUE             = 30
	vinLangLexerFALSE            = 31
	vinLangLexerFUNCTION         = 32
	vinLangLexerNUMBER           = 33
	vinLangLexerDIGIT            = 34
	vinLangLexerID               = 35
	vinLangLexerCHAR             = 36
	vinLangLexerSPACE            = 37
)
