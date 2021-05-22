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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 83, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 6, 12, 58, 10,
	12, 13, 12, 14, 12, 59, 3, 12, 3, 12, 6, 12, 64, 10, 12, 13, 12, 14, 12,
	65, 5, 12, 68, 10, 12, 3, 13, 3, 13, 3, 14, 6, 14, 73, 10, 14, 13, 14,
	14, 14, 74, 3, 15, 5, 15, 78, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 2, 2,
	17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 3, 2, 5, 3, 2, 50, 59, 6, 2, 50,
	59, 67, 92, 99, 124, 194, 7891, 5, 2, 11, 12, 15, 15, 34, 34, 2, 86, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2,
	2, 2, 5, 35, 3, 2, 2, 2, 7, 37, 3, 2, 2, 2, 9, 42, 3, 2, 2, 2, 11, 44,
	3, 2, 2, 2, 13, 46, 3, 2, 2, 2, 15, 48, 3, 2, 2, 2, 17, 50, 3, 2, 2, 2,
	19, 52, 3, 2, 2, 2, 21, 54, 3, 2, 2, 2, 23, 57, 3, 2, 2, 2, 25, 69, 3,
	2, 2, 2, 27, 72, 3, 2, 2, 2, 29, 77, 3, 2, 2, 2, 31, 79, 3, 2, 2, 2, 33,
	34, 7, 42, 2, 2, 34, 4, 3, 2, 2, 2, 35, 36, 7, 43, 2, 2, 36, 6, 3, 2, 2,
	2, 37, 38, 7, 100, 2, 2, 38, 39, 7, 107, 2, 2, 39, 40, 7, 7873, 2, 2, 40,
	41, 7, 112, 2, 2, 41, 8, 3, 2, 2, 2, 42, 43, 7, 60, 2, 2, 43, 10, 3, 2,
	2, 2, 44, 45, 7, 61, 2, 2, 45, 12, 3, 2, 2, 2, 46, 47, 7, 63, 2, 2, 47,
	14, 3, 2, 2, 2, 48, 49, 7, 45, 2, 2, 49, 16, 3, 2, 2, 2, 50, 51, 7, 47,
	2, 2, 51, 18, 3, 2, 2, 2, 52, 53, 7, 44, 2, 2, 53, 20, 3, 2, 2, 2, 54,
	55, 7, 49, 2, 2, 55, 22, 3, 2, 2, 2, 56, 58, 5, 25, 13, 2, 57, 56, 3, 2,
	2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 67,
	3, 2, 2, 2, 61, 63, 7, 48, 2, 2, 62, 64, 5, 25, 13, 2, 63, 62, 3, 2, 2,
	2, 64, 65, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68,
	3, 2, 2, 2, 67, 61, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 24, 3, 2, 2, 2,
	69, 70, 9, 2, 2, 2, 70, 26, 3, 2, 2, 2, 71, 73, 5, 29, 15, 2, 72, 71, 3,
	2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75,
	28, 3, 2, 2, 2, 76, 78, 9, 3, 2, 2, 77, 76, 3, 2, 2, 2, 78, 30, 3, 2, 2,
	2, 79, 80, 9, 4, 2, 2, 80, 81, 3, 2, 2, 2, 81, 82, 8, 16, 2, 2, 82, 32,
	3, 2, 2, 2, 8, 2, 59, 65, 67, 74, 77, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'bi\u1EBFn'", "':'", "';'", "'='", "'+'", "'-'", "'*'",
	"'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "VAR", "COLON", "SEMICOLIN", "ASSIGN", "ADD", "SUB", "MUL",
	"DIV", "NUMBER", "DIGIT", "ID", "CHAR", "SPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "VAR", "COLON", "SEMICOLIN", "ASSIGN", "ADD", "SUB", "MUL",
	"DIV", "NUMBER", "DIGIT", "ID", "CHAR", "SPACE",
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
	vinLangLexerT__0      = 1
	vinLangLexerT__1      = 2
	vinLangLexerVAR       = 3
	vinLangLexerCOLON     = 4
	vinLangLexerSEMICOLIN = 5
	vinLangLexerASSIGN    = 6
	vinLangLexerADD       = 7
	vinLangLexerSUB       = 8
	vinLangLexerMUL       = 9
	vinLangLexerDIV       = 10
	vinLangLexerNUMBER    = 11
	vinLangLexerDIGIT     = 12
	vinLangLexerID        = 13
	vinLangLexerCHAR      = 14
	vinLangLexerSPACE     = 15
)
