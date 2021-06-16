// Code generated from ./antlr/vinLang.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // vinLang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 39, 214,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 31, 10, 3, 12, 3, 14, 3, 34, 11, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 62, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 74, 10, 6, 12, 6, 14, 6, 77, 11, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 7, 9, 107, 10, 9, 12, 9, 14, 9, 110, 11, 9, 5, 9, 112, 10,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 119, 10, 10, 12, 10, 14, 10,
	122, 11, 10, 5, 10, 124, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 7, 11, 143, 10, 11, 12, 11, 14, 11, 146, 11, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 5, 11, 153, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 187, 10, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	5, 13, 201, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 209,
	10, 13, 12, 13, 14, 13, 212, 11, 13, 3, 13, 2, 3, 24, 14, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 2, 4, 3, 2, 10, 11, 3, 2, 8, 9, 2, 234, 2,
	26, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 61, 3, 2, 2, 2, 8, 63, 3, 2, 2, 2,
	10, 67, 3, 2, 2, 2, 12, 78, 3, 2, 2, 2, 14, 88, 3, 2, 2, 2, 16, 111, 3,
	2, 2, 2, 18, 113, 3, 2, 2, 2, 20, 127, 3, 2, 2, 2, 22, 186, 3, 2, 2, 2,
	24, 200, 3, 2, 2, 2, 26, 27, 5, 4, 3, 2, 27, 3, 3, 2, 2, 2, 28, 32, 5,
	6, 4, 2, 29, 31, 5, 6, 4, 2, 30, 29, 3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32,
	30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 5, 3, 2, 2, 2, 34, 32, 3, 2, 2,
	2, 35, 36, 5, 10, 6, 2, 36, 37, 7, 6, 2, 2, 37, 62, 3, 2, 2, 2, 38, 62,
	5, 12, 7, 2, 39, 62, 5, 20, 11, 2, 40, 41, 5, 18, 10, 2, 41, 42, 7, 6,
	2, 2, 42, 62, 3, 2, 2, 2, 43, 62, 5, 14, 8, 2, 44, 45, 5, 8, 5, 2, 45,
	46, 7, 6, 2, 2, 46, 62, 3, 2, 2, 2, 47, 48, 7, 37, 2, 2, 48, 62, 7, 6,
	2, 2, 49, 50, 7, 28, 2, 2, 50, 51, 5, 24, 13, 2, 51, 52, 7, 6, 2, 2, 52,
	62, 3, 2, 2, 2, 53, 54, 7, 26, 2, 2, 54, 62, 7, 6, 2, 2, 55, 56, 7, 27,
	2, 2, 56, 62, 7, 6, 2, 2, 57, 58, 7, 21, 2, 2, 58, 59, 5, 4, 3, 2, 59,
	60, 7, 22, 2, 2, 60, 62, 3, 2, 2, 2, 61, 35, 3, 2, 2, 2, 61, 38, 3, 2,
	2, 2, 61, 39, 3, 2, 2, 2, 61, 40, 3, 2, 2, 2, 61, 43, 3, 2, 2, 2, 61, 44,
	3, 2, 2, 2, 61, 47, 3, 2, 2, 2, 61, 49, 3, 2, 2, 2, 61, 53, 3, 2, 2, 2,
	61, 55, 3, 2, 2, 2, 61, 57, 3, 2, 2, 2, 62, 7, 3, 2, 2, 2, 63, 64, 7, 37,
	2, 2, 64, 65, 7, 7, 2, 2, 65, 66, 5, 24, 13, 2, 66, 9, 3, 2, 2, 2, 67,
	68, 7, 4, 2, 2, 68, 69, 7, 37, 2, 2, 69, 70, 7, 5, 2, 2, 70, 75, 7, 37,
	2, 2, 71, 72, 7, 7, 2, 2, 72, 74, 5, 24, 13, 2, 73, 71, 3, 2, 2, 2, 74,
	77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 11, 3, 2, 2,
	2, 77, 75, 3, 2, 2, 2, 78, 79, 7, 25, 2, 2, 79, 80, 7, 23, 2, 2, 80, 81,
	5, 10, 6, 2, 81, 82, 7, 6, 2, 2, 82, 83, 5, 22, 12, 2, 83, 84, 7, 6, 2,
	2, 84, 85, 5, 8, 5, 2, 85, 86, 7, 24, 2, 2, 86, 87, 5, 6, 4, 2, 87, 13,
	3, 2, 2, 2, 88, 89, 7, 34, 2, 2, 89, 90, 7, 37, 2, 2, 90, 91, 7, 23, 2,
	2, 91, 92, 5, 16, 9, 2, 92, 93, 7, 24, 2, 2, 93, 94, 7, 5, 2, 2, 94, 95,
	7, 37, 2, 2, 95, 96, 7, 21, 2, 2, 96, 97, 5, 4, 3, 2, 97, 98, 7, 22, 2,
	2, 98, 15, 3, 2, 2, 2, 99, 100, 7, 37, 2, 2, 100, 101, 7, 5, 2, 2, 101,
	108, 7, 37, 2, 2, 102, 103, 7, 3, 2, 2, 103, 104, 7, 37, 2, 2, 104, 105,
	7, 5, 2, 2, 105, 107, 7, 37, 2, 2, 106, 102, 3, 2, 2, 2, 107, 110, 3, 2,
	2, 2, 108, 106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2,
	110, 108, 3, 2, 2, 2, 111, 99, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 17,
	3, 2, 2, 2, 113, 114, 7, 37, 2, 2, 114, 123, 7, 23, 2, 2, 115, 120, 5,
	24, 13, 2, 116, 117, 7, 3, 2, 2, 117, 119, 5, 24, 13, 2, 118, 116, 3, 2,
	2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2,
	121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 115, 3, 2, 2, 2, 123,
	124, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 126, 7, 24, 2, 2, 126, 19,
	3, 2, 2, 2, 127, 128, 7, 29, 2, 2, 128, 129, 7, 23, 2, 2, 129, 130, 5,
	22, 12, 2, 130, 131, 7, 24, 2, 2, 131, 132, 7, 21, 2, 2, 132, 133, 5, 4,
	3, 2, 133, 144, 7, 22, 2, 2, 134, 135, 7, 31, 2, 2, 135, 136, 7, 23, 2,
	2, 136, 137, 5, 22, 12, 2, 137, 138, 7, 24, 2, 2, 138, 139, 7, 21, 2, 2,
	139, 140, 5, 4, 3, 2, 140, 141, 7, 22, 2, 2, 141, 143, 3, 2, 2, 2, 142,
	134, 3, 2, 2, 2, 143, 146, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 145,
	3, 2, 2, 2, 145, 152, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 147, 148, 7, 30,
	2, 2, 148, 149, 7, 21, 2, 2, 149, 150, 5, 4, 3, 2, 150, 151, 7, 22, 2,
	2, 151, 153, 3, 2, 2, 2, 152, 147, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153,
	21, 3, 2, 2, 2, 154, 187, 7, 32, 2, 2, 155, 187, 7, 33, 2, 2, 156, 157,
	5, 24, 13, 2, 157, 158, 7, 17, 2, 2, 158, 159, 5, 24, 13, 2, 159, 187,
	3, 2, 2, 2, 160, 161, 5, 24, 13, 2, 161, 162, 7, 15, 2, 2, 162, 163, 5,
	24, 13, 2, 163, 187, 3, 2, 2, 2, 164, 165, 5, 24, 13, 2, 165, 166, 7, 14,
	2, 2, 166, 167, 5, 24, 13, 2, 167, 187, 3, 2, 2, 2, 168, 169, 5, 24, 13,
	2, 169, 170, 7, 13, 2, 2, 170, 171, 5, 24, 13, 2, 171, 187, 3, 2, 2, 2,
	172, 173, 5, 24, 13, 2, 173, 174, 7, 16, 2, 2, 174, 175, 5, 24, 13, 2,
	175, 187, 3, 2, 2, 2, 176, 177, 5, 24, 13, 2, 177, 178, 7, 18, 2, 2, 178,
	179, 5, 24, 13, 2, 179, 187, 3, 2, 2, 2, 180, 181, 7, 12, 2, 2, 181, 187,
	5, 22, 12, 2, 182, 183, 7, 23, 2, 2, 183, 184, 5, 22, 12, 2, 184, 185,
	7, 24, 2, 2, 185, 187, 3, 2, 2, 2, 186, 154, 3, 2, 2, 2, 186, 155, 3, 2,
	2, 2, 186, 156, 3, 2, 2, 2, 186, 160, 3, 2, 2, 2, 186, 164, 3, 2, 2, 2,
	186, 168, 3, 2, 2, 2, 186, 172, 3, 2, 2, 2, 186, 176, 3, 2, 2, 2, 186,
	180, 3, 2, 2, 2, 186, 182, 3, 2, 2, 2, 187, 23, 3, 2, 2, 2, 188, 189, 8,
	13, 1, 2, 189, 201, 7, 35, 2, 2, 190, 201, 5, 18, 10, 2, 191, 201, 7, 37,
	2, 2, 192, 193, 7, 23, 2, 2, 193, 194, 5, 24, 13, 2, 194, 195, 7, 24, 2,
	2, 195, 201, 3, 2, 2, 2, 196, 197, 7, 21, 2, 2, 197, 198, 5, 24, 13, 2,
	198, 199, 7, 22, 2, 2, 199, 201, 3, 2, 2, 2, 200, 188, 3, 2, 2, 2, 200,
	190, 3, 2, 2, 2, 200, 191, 3, 2, 2, 2, 200, 192, 3, 2, 2, 2, 200, 196,
	3, 2, 2, 2, 201, 210, 3, 2, 2, 2, 202, 203, 12, 6, 2, 2, 203, 204, 9, 2,
	2, 2, 204, 209, 5, 24, 13, 7, 205, 206, 12, 5, 2, 2, 206, 207, 9, 3, 2,
	2, 207, 209, 5, 24, 13, 6, 208, 202, 3, 2, 2, 2, 208, 205, 3, 2, 2, 2,
	209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211,
	25, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 15, 32, 61, 75, 108, 111, 120, 123,
	144, 152, 186, 200, 208, 210,
}
var literalNames = []string{
	"", "','", "'bi\u1EBFn'", "':'", "';'", "'='", "'+'", "'-'", "'*'", "",
	"'!'", "'<='", "'<'", "'>'", "'>='", "'=='", "'!='", "'++'", "'--'", "'{'",
	"'}'", "'('", "')'", "'l\u1EB7p'", "'ti\u1EBFp'", "'d\u1EEBng'", "'tr\u1EA3'",
	"'n\u1EBFu'", "'th\u00EC'", "'kh\u00F4ng_th\u00EC'", "'\u0111\u00FAng'",
	"'sai'", "'h\u00E0m'",
}
var symbolicNames = []string{
	"", "", "VAR", "COLON", "SEMICOLIN", "ASSIGN", "ADD", "SUB", "MUL", "DIV",
	"NEG", "LE", "L", "G", "GE", "E", "NE", "INCREMENT", "DECREMENT", "LEFT_BRACE",
	"RIGHT_BRACE", "LEFT_PARENTHESE", "RIGHT_PARENTHESE", "FOR", "CONTINUE",
	"BREAK", "RETURN", "IF", "ELSE", "ELIF", "TRUE", "FALSE", "FUNCTION", "NUMBER",
	"DIGIT", "ID", "CHAR", "SPACE",
}

var ruleNames = []string{
	"program", "sequenceStatement", "statement", "assign", "declaration", "forStatement",
	"functionDeclaration", "functionArgs", "functionCall", "ifStatement", "boolExpression",
	"expression",
}

type vinLangParser struct {
	*antlr.BaseParser
}

// NewvinLangParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *vinLangParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewvinLangParser(input antlr.TokenStream) *vinLangParser {
	this := new(vinLangParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "vinLang.g4"

	return this
}

// vinLangParser tokens.
const (
	vinLangParserEOF              = antlr.TokenEOF
	vinLangParserT__0             = 1
	vinLangParserVAR              = 2
	vinLangParserCOLON            = 3
	vinLangParserSEMICOLIN        = 4
	vinLangParserASSIGN           = 5
	vinLangParserADD              = 6
	vinLangParserSUB              = 7
	vinLangParserMUL              = 8
	vinLangParserDIV              = 9
	vinLangParserNEG              = 10
	vinLangParserLE               = 11
	vinLangParserL                = 12
	vinLangParserG                = 13
	vinLangParserGE               = 14
	vinLangParserE                = 15
	vinLangParserNE               = 16
	vinLangParserINCREMENT        = 17
	vinLangParserDECREMENT        = 18
	vinLangParserLEFT_BRACE       = 19
	vinLangParserRIGHT_BRACE      = 20
	vinLangParserLEFT_PARENTHESE  = 21
	vinLangParserRIGHT_PARENTHESE = 22
	vinLangParserFOR              = 23
	vinLangParserCONTINUE         = 24
	vinLangParserBREAK            = 25
	vinLangParserRETURN           = 26
	vinLangParserIF               = 27
	vinLangParserELSE             = 28
	vinLangParserELIF             = 29
	vinLangParserTRUE             = 30
	vinLangParserFALSE            = 31
	vinLangParserFUNCTION         = 32
	vinLangParserNUMBER           = 33
	vinLangParserDIGIT            = 34
	vinLangParserID               = 35
	vinLangParserCHAR             = 36
	vinLangParserSPACE            = 37
)

// vinLangParser rules.
const (
	vinLangParserRULE_program             = 0
	vinLangParserRULE_sequenceStatement   = 1
	vinLangParserRULE_statement           = 2
	vinLangParserRULE_assign              = 3
	vinLangParserRULE_declaration         = 4
	vinLangParserRULE_forStatement        = 5
	vinLangParserRULE_functionDeclaration = 6
	vinLangParserRULE_functionArgs        = 7
	vinLangParserRULE_functionCall        = 8
	vinLangParserRULE_ifStatement         = 9
	vinLangParserRULE_boolExpression      = 10
	vinLangParserRULE_expression          = 11
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) SequenceStatement() ISequenceStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *vinLangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, vinLangParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.SequenceStatement()
	}

	return localctx
}

// ISequenceStatementContext is an interface to support dynamic dispatch.
type ISequenceStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequenceStatementContext differentiates from other interfaces.
	IsSequenceStatementContext()
}

type SequenceStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceStatementContext() *SequenceStatementContext {
	var p = new(SequenceStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_sequenceStatement
	return p
}

func (*SequenceStatementContext) IsSequenceStatementContext() {}

func NewSequenceStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceStatementContext {
	var p = new(SequenceStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_sequenceStatement

	return p
}

func (s *SequenceStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SequenceStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SequenceStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterSequenceStatement(s)
	}
}

func (s *SequenceStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitSequenceStatement(s)
	}
}

func (p *vinLangParser) SequenceStatement() (localctx ISequenceStatementContext) {
	localctx = NewSequenceStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, vinLangParserRULE_sequenceStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Statement()
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<vinLangParserVAR)|(1<<vinLangParserLEFT_BRACE)|(1<<vinLangParserFOR)|(1<<vinLangParserCONTINUE)|(1<<vinLangParserBREAK)|(1<<vinLangParserRETURN)|(1<<vinLangParserIF))) != 0) || _la == vinLangParserFUNCTION || _la == vinLangParserID {
		{
			p.SetState(27)
			p.Statement()
		}

		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) SEMICOLIN() antlr.TerminalNode {
	return s.GetToken(vinLangParserSEMICOLIN, 0)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StatementContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *StatementContext) ID() antlr.TerminalNode {
	return s.GetToken(vinLangParserID, 0)
}

func (s *StatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(vinLangParserRETURN, 0)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(vinLangParserCONTINUE, 0)
}

func (s *StatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(vinLangParserBREAK, 0)
}

func (s *StatementContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_BRACE, 0)
}

func (s *StatementContext) SequenceStatement() ISequenceStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceStatementContext)
}

func (s *StatementContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_BRACE, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *vinLangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, vinLangParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Declaration()
		}
		{
			p.SetState(34)
			p.Match(vinLangParserSEMICOLIN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.ForStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)
			p.IfStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(38)
			p.FunctionCall()
		}
		{
			p.SetState(39)
			p.Match(vinLangParserSEMICOLIN)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(41)
			p.FunctionDeclaration()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(42)
			p.Assign()
		}
		{
			p.SetState(43)
			p.Match(vinLangParserSEMICOLIN)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(45)
			p.Match(vinLangParserID)
		}
		{
			p.SetState(46)
			p.Match(vinLangParserSEMICOLIN)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(47)
			p.Match(vinLangParserRETURN)
		}
		{
			p.SetState(48)
			p.expression(0)
		}
		{
			p.SetState(49)
			p.Match(vinLangParserSEMICOLIN)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(51)
			p.Match(vinLangParserCONTINUE)
		}
		{
			p.SetState(52)
			p.Match(vinLangParserSEMICOLIN)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(53)
			p.Match(vinLangParserBREAK)
		}
		{
			p.SetState(54)
			p.Match(vinLangParserSEMICOLIN)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(55)
			p.Match(vinLangParserLEFT_BRACE)
		}
		{
			p.SetState(56)
			p.SequenceStatement()
		}
		{
			p.SetState(57)
			p.Match(vinLangParserRIGHT_BRACE)
		}

	}

	return localctx
}

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_assign
	return p
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(vinLangParserID, 0)
}

func (s *AssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(vinLangParserASSIGN, 0)
}

func (s *AssignContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *vinLangParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, vinLangParserRULE_assign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(vinLangParserID)
	}
	{
		p.SetState(62)
		p.Match(vinLangParserASSIGN)
	}
	{
		p.SetState(63)
		p.expression(0)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(vinLangParserVAR, 0)
}

func (s *DeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserID)
}

func (s *DeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserID, i)
}

func (s *DeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(vinLangParserCOLON, 0)
}

func (s *DeclarationContext) AllASSIGN() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserASSIGN)
}

func (s *DeclarationContext) ASSIGN(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserASSIGN, i)
}

func (s *DeclarationContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DeclarationContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *vinLangParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, vinLangParserRULE_declaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(vinLangParserVAR)
	}
	{
		p.SetState(66)
		p.Match(vinLangParserID)
	}
	{
		p.SetState(67)
		p.Match(vinLangParserCOLON)
	}
	{
		p.SetState(68)
		p.Match(vinLangParserID)
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == vinLangParserASSIGN {
		{
			p.SetState(69)
			p.Match(vinLangParserASSIGN)
		}
		{
			p.SetState(70)
			p.expression(0)
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_forStatement
	return p
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(vinLangParserFOR, 0)
}

func (s *ForStatementContext) LEFT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_PARENTHESE, 0)
}

func (s *ForStatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *ForStatementContext) AllSEMICOLIN() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserSEMICOLIN)
}

func (s *ForStatementContext) SEMICOLIN(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserSEMICOLIN, i)
}

func (s *ForStatementContext) BoolExpression() IBoolExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *ForStatementContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ForStatementContext) RIGHT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_PARENTHESE, 0)
}

func (s *ForStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (p *vinLangParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, vinLangParserRULE_forStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(vinLangParserFOR)
	}
	{
		p.SetState(77)
		p.Match(vinLangParserLEFT_PARENTHESE)
	}
	{
		p.SetState(78)
		p.Declaration()
	}
	{
		p.SetState(79)
		p.Match(vinLangParserSEMICOLIN)
	}
	{
		p.SetState(80)
		p.BoolExpression()
	}
	{
		p.SetState(81)
		p.Match(vinLangParserSEMICOLIN)
	}
	{
		p.SetState(82)
		p.Assign()
	}
	{
		p.SetState(83)
		p.Match(vinLangParserRIGHT_PARENTHESE)
	}
	{
		p.SetState(84)
		p.Statement()
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(vinLangParserFUNCTION, 0)
}

func (s *FunctionDeclarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserID)
}

func (s *FunctionDeclarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserID, i)
}

func (s *FunctionDeclarationContext) LEFT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_PARENTHESE, 0)
}

func (s *FunctionDeclarationContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionDeclarationContext) RIGHT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_PARENTHESE, 0)
}

func (s *FunctionDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(vinLangParserCOLON, 0)
}

func (s *FunctionDeclarationContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_BRACE, 0)
}

func (s *FunctionDeclarationContext) SequenceStatement() ISequenceStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceStatementContext)
}

func (s *FunctionDeclarationContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_BRACE, 0)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *vinLangParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, vinLangParserRULE_functionDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(vinLangParserFUNCTION)
	}
	{
		p.SetState(87)
		p.Match(vinLangParserID)
	}
	{
		p.SetState(88)
		p.Match(vinLangParserLEFT_PARENTHESE)
	}
	{
		p.SetState(89)
		p.FunctionArgs()
	}
	{
		p.SetState(90)
		p.Match(vinLangParserRIGHT_PARENTHESE)
	}
	{
		p.SetState(91)
		p.Match(vinLangParserCOLON)
	}
	{
		p.SetState(92)
		p.Match(vinLangParserID)
	}
	{
		p.SetState(93)
		p.Match(vinLangParserLEFT_BRACE)
	}
	{
		p.SetState(94)
		p.SequenceStatement()
	}
	{
		p.SetState(95)
		p.Match(vinLangParserRIGHT_BRACE)
	}

	return localctx
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserID)
}

func (s *FunctionArgsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserID, i)
}

func (s *FunctionArgsContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserCOLON)
}

func (s *FunctionArgsContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserCOLON, i)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (p *vinLangParser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, vinLangParserRULE_functionArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == vinLangParserID {
		{
			p.SetState(97)
			p.Match(vinLangParserID)
		}
		{
			p.SetState(98)
			p.Match(vinLangParserCOLON)
		}
		{
			p.SetState(99)
			p.Match(vinLangParserID)
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == vinLangParserT__0 {
			{
				p.SetState(100)
				p.Match(vinLangParserT__0)
			}
			{
				p.SetState(101)
				p.Match(vinLangParserID)
			}
			{
				p.SetState(102)
				p.Match(vinLangParserCOLON)
			}
			{
				p.SetState(103)
				p.Match(vinLangParserID)
			}

			p.SetState(108)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(vinLangParserID, 0)
}

func (s *FunctionCallContext) LEFT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_PARENTHESE, 0)
}

func (s *FunctionCallContext) RIGHT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_PARENTHESE, 0)
}

func (s *FunctionCallContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FunctionCallContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *vinLangParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, vinLangParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(vinLangParserID)
	}
	{
		p.SetState(112)
		p.Match(vinLangParserLEFT_PARENTHESE)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(vinLangParserLEFT_BRACE-19))|(1<<(vinLangParserLEFT_PARENTHESE-19))|(1<<(vinLangParserNUMBER-19))|(1<<(vinLangParserID-19)))) != 0 {
		{
			p.SetState(113)
			p.expression(0)
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == vinLangParserT__0 {
			{
				p.SetState(114)
				p.Match(vinLangParserT__0)
			}
			{
				p.SetState(115)
				p.expression(0)
			}

			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(123)
		p.Match(vinLangParserRIGHT_PARENTHESE)
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(vinLangParserIF, 0)
}

func (s *IfStatementContext) AllLEFT_PARENTHESE() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserLEFT_PARENTHESE)
}

func (s *IfStatementContext) LEFT_PARENTHESE(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_PARENTHESE, i)
}

func (s *IfStatementContext) AllBoolExpression() []IBoolExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExpressionContext)(nil)).Elem())
	var tst = make([]IBoolExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExpressionContext)
		}
	}

	return tst
}

func (s *IfStatementContext) BoolExpression(i int) IBoolExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *IfStatementContext) AllRIGHT_PARENTHESE() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserRIGHT_PARENTHESE)
}

func (s *IfStatementContext) RIGHT_PARENTHESE(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_PARENTHESE, i)
}

func (s *IfStatementContext) AllLEFT_BRACE() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserLEFT_BRACE)
}

func (s *IfStatementContext) LEFT_BRACE(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_BRACE, i)
}

func (s *IfStatementContext) AllSequenceStatement() []ISequenceStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISequenceStatementContext)(nil)).Elem())
	var tst = make([]ISequenceStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISequenceStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) SequenceStatement(i int) ISequenceStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISequenceStatementContext)
}

func (s *IfStatementContext) AllRIGHT_BRACE() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserRIGHT_BRACE)
}

func (s *IfStatementContext) RIGHT_BRACE(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_BRACE, i)
}

func (s *IfStatementContext) AllELIF() []antlr.TerminalNode {
	return s.GetTokens(vinLangParserELIF)
}

func (s *IfStatementContext) ELIF(i int) antlr.TerminalNode {
	return s.GetToken(vinLangParserELIF, i)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(vinLangParserELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *vinLangParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, vinLangParserRULE_ifStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(vinLangParserIF)
	}
	{
		p.SetState(126)
		p.Match(vinLangParserLEFT_PARENTHESE)
	}
	{
		p.SetState(127)
		p.BoolExpression()
	}
	{
		p.SetState(128)
		p.Match(vinLangParserRIGHT_PARENTHESE)
	}
	{
		p.SetState(129)
		p.Match(vinLangParserLEFT_BRACE)
	}
	{
		p.SetState(130)
		p.SequenceStatement()
	}
	{
		p.SetState(131)
		p.Match(vinLangParserRIGHT_BRACE)
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == vinLangParserELIF {
		{
			p.SetState(132)
			p.Match(vinLangParserELIF)
		}
		{
			p.SetState(133)
			p.Match(vinLangParserLEFT_PARENTHESE)
		}
		{
			p.SetState(134)
			p.BoolExpression()
		}
		{
			p.SetState(135)
			p.Match(vinLangParserRIGHT_PARENTHESE)
		}
		{
			p.SetState(136)
			p.Match(vinLangParserLEFT_BRACE)
		}
		{
			p.SetState(137)
			p.SequenceStatement()
		}
		{
			p.SetState(138)
			p.Match(vinLangParserRIGHT_BRACE)
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == vinLangParserELSE {
		{
			p.SetState(145)
			p.Match(vinLangParserELSE)
		}
		{
			p.SetState(146)
			p.Match(vinLangParserLEFT_BRACE)
		}
		{
			p.SetState(147)
			p.SequenceStatement()
		}
		{
			p.SetState(148)
			p.Match(vinLangParserRIGHT_BRACE)
		}

	}

	return localctx
}

// IBoolExpressionContext is an interface to support dynamic dispatch.
type IBoolExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExpressionContext differentiates from other interfaces.
	IsBoolExpressionContext()
}

type BoolExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExpressionContext() *BoolExpressionContext {
	var p = new(BoolExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_boolExpression
	return p
}

func (*BoolExpressionContext) IsBoolExpressionContext() {}

func NewBoolExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExpressionContext {
	var p = new(BoolExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_boolExpression

	return p
}

func (s *BoolExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExpressionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(vinLangParserTRUE, 0)
}

func (s *BoolExpressionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(vinLangParserFALSE, 0)
}

func (s *BoolExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BoolExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BoolExpressionContext) E() antlr.TerminalNode {
	return s.GetToken(vinLangParserE, 0)
}

func (s *BoolExpressionContext) G() antlr.TerminalNode {
	return s.GetToken(vinLangParserG, 0)
}

func (s *BoolExpressionContext) L() antlr.TerminalNode {
	return s.GetToken(vinLangParserL, 0)
}

func (s *BoolExpressionContext) LE() antlr.TerminalNode {
	return s.GetToken(vinLangParserLE, 0)
}

func (s *BoolExpressionContext) GE() antlr.TerminalNode {
	return s.GetToken(vinLangParserGE, 0)
}

func (s *BoolExpressionContext) NE() antlr.TerminalNode {
	return s.GetToken(vinLangParserNE, 0)
}

func (s *BoolExpressionContext) NEG() antlr.TerminalNode {
	return s.GetToken(vinLangParserNEG, 0)
}

func (s *BoolExpressionContext) BoolExpression() IBoolExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *BoolExpressionContext) LEFT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_PARENTHESE, 0)
}

func (s *BoolExpressionContext) RIGHT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_PARENTHESE, 0)
}

func (s *BoolExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterBoolExpression(s)
	}
}

func (s *BoolExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitBoolExpression(s)
	}
}

func (p *vinLangParser) BoolExpression() (localctx IBoolExpressionContext) {
	localctx = NewBoolExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, vinLangParserRULE_boolExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(vinLangParserTRUE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Match(vinLangParserFALSE)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.expression(0)
		}
		{
			p.SetState(155)
			p.Match(vinLangParserE)
		}
		{
			p.SetState(156)
			p.expression(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.expression(0)
		}
		{
			p.SetState(159)
			p.Match(vinLangParserG)
		}
		{
			p.SetState(160)
			p.expression(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(162)
			p.expression(0)
		}
		{
			p.SetState(163)
			p.Match(vinLangParserL)
		}
		{
			p.SetState(164)
			p.expression(0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(166)
			p.expression(0)
		}
		{
			p.SetState(167)
			p.Match(vinLangParserLE)
		}
		{
			p.SetState(168)
			p.expression(0)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(170)
			p.expression(0)
		}
		{
			p.SetState(171)
			p.Match(vinLangParserGE)
		}
		{
			p.SetState(172)
			p.expression(0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(174)
			p.expression(0)
		}
		{
			p.SetState(175)
			p.Match(vinLangParserNE)
		}
		{
			p.SetState(176)
			p.expression(0)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(178)
			p.Match(vinLangParserNEG)
		}
		{
			p.SetState(179)
			p.BoolExpression()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(180)
			p.Match(vinLangParserLEFT_PARENTHESE)
		}
		{
			p.SetState(181)
			p.BoolExpression()
		}
		{
			p.SetState(182)
			p.Match(vinLangParserRIGHT_PARENTHESE)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = vinLangParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = vinLangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(vinLangParserNUMBER, 0)
}

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(vinLangParserID, 0)
}

func (s *ExpressionContext) LEFT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_PARENTHESE, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RIGHT_PARENTHESE() antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_PARENTHESE, 0)
}

func (s *ExpressionContext) LEFT_BRACE() antlr.TerminalNode {
	return s.GetToken(vinLangParserLEFT_BRACE, 0)
}

func (s *ExpressionContext) RIGHT_BRACE() antlr.TerminalNode {
	return s.GetToken(vinLangParserRIGHT_BRACE, 0)
}

func (s *ExpressionContext) MUL() antlr.TerminalNode {
	return s.GetToken(vinLangParserMUL, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(vinLangParserDIV, 0)
}

func (s *ExpressionContext) ADD() antlr.TerminalNode {
	return s.GetToken(vinLangParserADD, 0)
}

func (s *ExpressionContext) SUB() antlr.TerminalNode {
	return s.GetToken(vinLangParserSUB, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vinLangListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *vinLangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *vinLangParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, vinLangParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(187)
			p.Match(vinLangParserNUMBER)
		}

	case 2:
		{
			p.SetState(188)
			p.FunctionCall()
		}

	case 3:
		{
			p.SetState(189)
			p.Match(vinLangParserID)
		}

	case 4:
		{
			p.SetState(190)
			p.Match(vinLangParserLEFT_PARENTHESE)
		}
		{
			p.SetState(191)
			p.expression(0)
		}
		{
			p.SetState(192)
			p.Match(vinLangParserRIGHT_PARENTHESE)
		}

	case 5:
		{
			p.SetState(194)
			p.Match(vinLangParserLEFT_BRACE)
		}
		{
			p.SetState(195)
			p.expression(0)
		}
		{
			p.SetState(196)
			p.Match(vinLangParserRIGHT_BRACE)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(206)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vinLangParserRULE_expression)
				p.SetState(200)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(201)
					_la = p.GetTokenStream().LA(1)

					if !(_la == vinLangParserMUL || _la == vinLangParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(202)
					p.expression(5)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vinLangParserRULE_expression)
				p.SetState(203)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(204)
					_la = p.GetTokenStream().LA(1)

					if !(_la == vinLangParserADD || _la == vinLangParserSUB) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(205)
					p.expression(4)
				}

			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

func (p *vinLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *vinLangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
