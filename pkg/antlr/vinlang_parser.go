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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 123,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 7, 3, 23, 10, 3, 12, 3,
	14, 3, 26, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 5, 4, 39, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 51, 10, 6, 12, 6, 14, 6, 54, 11, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 98, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 110, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 7, 9, 118, 10, 9, 12, 9, 14, 9, 121, 11, 9, 3, 9, 2, 3, 16,
	10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 4, 3, 2, 9, 10, 3, 2, 7, 8, 2, 133,
	2, 18, 3, 2, 2, 2, 4, 20, 3, 2, 2, 2, 6, 38, 3, 2, 2, 2, 8, 40, 3, 2, 2,
	2, 10, 44, 3, 2, 2, 2, 12, 55, 3, 2, 2, 2, 14, 97, 3, 2, 2, 2, 16, 109,
	3, 2, 2, 2, 18, 19, 5, 4, 3, 2, 19, 3, 3, 2, 2, 2, 20, 24, 5, 6, 4, 2,
	21, 23, 5, 6, 4, 2, 22, 21, 3, 2, 2, 2, 23, 26, 3, 2, 2, 2, 24, 22, 3,
	2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 5, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2, 27,
	28, 5, 10, 6, 2, 28, 29, 7, 5, 2, 2, 29, 39, 3, 2, 2, 2, 30, 39, 5, 12,
	7, 2, 31, 32, 5, 8, 5, 2, 32, 33, 7, 5, 2, 2, 33, 39, 3, 2, 2, 2, 34, 35,
	7, 20, 2, 2, 35, 36, 5, 4, 3, 2, 36, 37, 7, 21, 2, 2, 37, 39, 3, 2, 2,
	2, 38, 27, 3, 2, 2, 2, 38, 30, 3, 2, 2, 2, 38, 31, 3, 2, 2, 2, 38, 34,
	3, 2, 2, 2, 39, 7, 3, 2, 2, 2, 40, 41, 7, 35, 2, 2, 41, 42, 7, 6, 2, 2,
	42, 43, 5, 16, 9, 2, 43, 9, 3, 2, 2, 2, 44, 45, 7, 3, 2, 2, 45, 46, 7,
	35, 2, 2, 46, 47, 7, 4, 2, 2, 47, 52, 7, 35, 2, 2, 48, 49, 7, 6, 2, 2,
	49, 51, 5, 16, 9, 2, 50, 48, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3,
	2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 11, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55,
	56, 7, 24, 2, 2, 56, 57, 7, 22, 2, 2, 57, 58, 5, 10, 6, 2, 58, 59, 7, 5,
	2, 2, 59, 60, 5, 14, 8, 2, 60, 61, 7, 5, 2, 2, 61, 62, 5, 8, 5, 2, 62,
	63, 7, 23, 2, 2, 63, 64, 5, 6, 4, 2, 64, 13, 3, 2, 2, 2, 65, 98, 7, 31,
	2, 2, 66, 98, 7, 32, 2, 2, 67, 68, 5, 16, 9, 2, 68, 69, 7, 16, 2, 2, 69,
	70, 5, 16, 9, 2, 70, 98, 3, 2, 2, 2, 71, 72, 5, 16, 9, 2, 72, 73, 7, 14,
	2, 2, 73, 74, 5, 16, 9, 2, 74, 98, 3, 2, 2, 2, 75, 76, 5, 16, 9, 2, 76,
	77, 7, 13, 2, 2, 77, 78, 5, 16, 9, 2, 78, 98, 3, 2, 2, 2, 79, 80, 5, 16,
	9, 2, 80, 81, 7, 12, 2, 2, 81, 82, 5, 16, 9, 2, 82, 98, 3, 2, 2, 2, 83,
	84, 5, 16, 9, 2, 84, 85, 7, 15, 2, 2, 85, 86, 5, 16, 9, 2, 86, 98, 3, 2,
	2, 2, 87, 88, 5, 16, 9, 2, 88, 89, 7, 17, 2, 2, 89, 90, 5, 16, 9, 2, 90,
	98, 3, 2, 2, 2, 91, 92, 7, 11, 2, 2, 92, 98, 5, 14, 8, 2, 93, 94, 7, 22,
	2, 2, 94, 95, 5, 14, 8, 2, 95, 96, 7, 23, 2, 2, 96, 98, 3, 2, 2, 2, 97,
	65, 3, 2, 2, 2, 97, 66, 3, 2, 2, 2, 97, 67, 3, 2, 2, 2, 97, 71, 3, 2, 2,
	2, 97, 75, 3, 2, 2, 2, 97, 79, 3, 2, 2, 2, 97, 83, 3, 2, 2, 2, 97, 87,
	3, 2, 2, 2, 97, 91, 3, 2, 2, 2, 97, 93, 3, 2, 2, 2, 98, 15, 3, 2, 2, 2,
	99, 100, 8, 9, 1, 2, 100, 110, 7, 33, 2, 2, 101, 110, 7, 35, 2, 2, 102,
	103, 7, 22, 2, 2, 103, 104, 5, 16, 9, 2, 104, 105, 7, 23, 2, 2, 105, 110,
	3, 2, 2, 2, 106, 107, 7, 20, 2, 2, 107, 108, 7, 38, 2, 2, 108, 110, 7,
	21, 2, 2, 109, 99, 3, 2, 2, 2, 109, 101, 3, 2, 2, 2, 109, 102, 3, 2, 2,
	2, 109, 106, 3, 2, 2, 2, 110, 119, 3, 2, 2, 2, 111, 112, 12, 6, 2, 2, 112,
	113, 9, 2, 2, 2, 113, 118, 5, 16, 9, 7, 114, 115, 12, 5, 2, 2, 115, 116,
	9, 3, 2, 2, 116, 118, 5, 16, 9, 6, 117, 111, 3, 2, 2, 2, 117, 114, 3, 2,
	2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2,
	120, 17, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 9, 24, 38, 52, 97, 109, 117,
	119,
}
var literalNames = []string{
	"", "'bi\u1EBFn'", "':'", "';'", "'='", "'+'", "'-'", "'*'", "", "'!'",
	"'<='", "'<'", "'>'", "'>='", "", "", "'++'", "'--'", "'{'", "'}'", "'('",
	"')'", "'l\u1EB7p'", "'ti\u1EBFp'", "'d\u1EEBng'", "'tr\u1EA3'", "'n\u1EBFu'",
	"'th\u00EC'", "'kh\u00F4ng_th\u00EC'", "'\u0111\u00FAng'", "'sai'",
}
var symbolicNames = []string{
	"", "VAR", "COLON", "SEMICOLIN", "ASSIGN", "ADD", "SUB", "MUL", "DIV",
	"NEG", "LE", "L", "G", "GE", "E", "NE", "INCREMENT", "DECREMENT", "LEFT_BRACE",
	"RIGHT_BRACE", "LEFT_PARENTHESE", "RIGHT_PARENTHESE", "FOR", "CONTINUE",
	"BREAK", "RETURN", "IF", "ELSE", "ELIF", "TRUE", "FALSE", "NUMBER", "DIGIT",
	"ID", "CHAR", "SPACE", "SequenceStatement",
}

var ruleNames = []string{
	"program", "sequenceStatement", "statement", "assign", "declaration", "forStatement",
	"boolExpression", "expression",
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
	vinLangParserEOF               = antlr.TokenEOF
	vinLangParserVAR               = 1
	vinLangParserCOLON             = 2
	vinLangParserSEMICOLIN         = 3
	vinLangParserASSIGN            = 4
	vinLangParserADD               = 5
	vinLangParserSUB               = 6
	vinLangParserMUL               = 7
	vinLangParserDIV               = 8
	vinLangParserNEG               = 9
	vinLangParserLE                = 10
	vinLangParserL                 = 11
	vinLangParserG                 = 12
	vinLangParserGE                = 13
	vinLangParserE                 = 14
	vinLangParserNE                = 15
	vinLangParserINCREMENT         = 16
	vinLangParserDECREMENT         = 17
	vinLangParserLEFT_BRACE        = 18
	vinLangParserRIGHT_BRACE       = 19
	vinLangParserLEFT_PARENTHESE   = 20
	vinLangParserRIGHT_PARENTHESE  = 21
	vinLangParserFOR               = 22
	vinLangParserCONTINUE          = 23
	vinLangParserBREAK             = 24
	vinLangParserRETURN            = 25
	vinLangParserIF                = 26
	vinLangParserELSE              = 27
	vinLangParserELIF              = 28
	vinLangParserTRUE              = 29
	vinLangParserFALSE             = 30
	vinLangParserNUMBER            = 31
	vinLangParserDIGIT             = 32
	vinLangParserID                = 33
	vinLangParserCHAR              = 34
	vinLangParserSPACE             = 35
	vinLangParserSequenceStatement = 36
)

// vinLangParser rules.
const (
	vinLangParserRULE_program           = 0
	vinLangParserRULE_sequenceStatement = 1
	vinLangParserRULE_statement         = 2
	vinLangParserRULE_assign            = 3
	vinLangParserRULE_declaration       = 4
	vinLangParserRULE_forStatement      = 5
	vinLangParserRULE_boolExpression    = 6
	vinLangParserRULE_expression        = 7
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
		p.SetState(16)
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
		p.SetState(18)
		p.Statement()
	}
	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<vinLangParserVAR)|(1<<vinLangParserLEFT_BRACE)|(1<<vinLangParserFOR))) != 0) || _la == vinLangParserID {
		{
			p.SetState(19)
			p.Statement()
		}

		p.SetState(24)
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

func (s *StatementContext) Assign() IAssignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
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

	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case vinLangParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)
			p.Declaration()
		}
		{
			p.SetState(26)
			p.Match(vinLangParserSEMICOLIN)
		}

	case vinLangParserFOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.ForStatement()
		}

	case vinLangParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)
			p.Assign()
		}
		{
			p.SetState(30)
			p.Match(vinLangParserSEMICOLIN)
		}

	case vinLangParserLEFT_BRACE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(32)
			p.Match(vinLangParserLEFT_BRACE)
		}
		{
			p.SetState(33)
			p.SequenceStatement()
		}
		{
			p.SetState(34)
			p.Match(vinLangParserRIGHT_BRACE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
		p.SetState(38)
		p.Match(vinLangParserID)
	}
	{
		p.SetState(39)
		p.Match(vinLangParserASSIGN)
	}
	{
		p.SetState(40)
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
		p.SetState(42)
		p.Match(vinLangParserVAR)
	}
	{
		p.SetState(43)
		p.Match(vinLangParserID)
	}
	{
		p.SetState(44)
		p.Match(vinLangParserCOLON)
	}
	{
		p.SetState(45)
		p.Match(vinLangParserID)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == vinLangParserASSIGN {
		{
			p.SetState(46)
			p.Match(vinLangParserASSIGN)
		}
		{
			p.SetState(47)
			p.expression(0)
		}

		p.SetState(52)
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
		p.SetState(53)
		p.Match(vinLangParserFOR)
	}
	{
		p.SetState(54)
		p.Match(vinLangParserLEFT_PARENTHESE)
	}
	{
		p.SetState(55)
		p.Declaration()
	}
	{
		p.SetState(56)
		p.Match(vinLangParserSEMICOLIN)
	}
	{
		p.SetState(57)
		p.BoolExpression()
	}
	{
		p.SetState(58)
		p.Match(vinLangParserSEMICOLIN)
	}
	{
		p.SetState(59)
		p.Assign()
	}
	{
		p.SetState(60)
		p.Match(vinLangParserRIGHT_PARENTHESE)
	}
	{
		p.SetState(61)
		p.Statement()
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
	p.EnterRule(localctx, 12, vinLangParserRULE_boolExpression)

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

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(vinLangParserTRUE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Match(vinLangParserFALSE)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(65)
			p.expression(0)
		}
		{
			p.SetState(66)
			p.Match(vinLangParserE)
		}
		{
			p.SetState(67)
			p.expression(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(69)
			p.expression(0)
		}
		{
			p.SetState(70)
			p.Match(vinLangParserG)
		}
		{
			p.SetState(71)
			p.expression(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(73)
			p.expression(0)
		}
		{
			p.SetState(74)
			p.Match(vinLangParserL)
		}
		{
			p.SetState(75)
			p.expression(0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(77)
			p.expression(0)
		}
		{
			p.SetState(78)
			p.Match(vinLangParserLE)
		}
		{
			p.SetState(79)
			p.expression(0)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(81)
			p.expression(0)
		}
		{
			p.SetState(82)
			p.Match(vinLangParserGE)
		}
		{
			p.SetState(83)
			p.expression(0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(85)
			p.expression(0)
		}
		{
			p.SetState(86)
			p.Match(vinLangParserNE)
		}
		{
			p.SetState(87)
			p.expression(0)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(89)
			p.Match(vinLangParserNEG)
		}
		{
			p.SetState(90)
			p.BoolExpression()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(91)
			p.Match(vinLangParserLEFT_PARENTHESE)
		}
		{
			p.SetState(92)
			p.BoolExpression()
		}
		{
			p.SetState(93)
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

func (s *ExpressionContext) SequenceStatement() antlr.TerminalNode {
	return s.GetToken(vinLangParserSequenceStatement, 0)
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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, vinLangParserRULE_expression, _p)
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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case vinLangParserNUMBER:
		{
			p.SetState(98)
			p.Match(vinLangParserNUMBER)
		}

	case vinLangParserID:
		{
			p.SetState(99)
			p.Match(vinLangParserID)
		}

	case vinLangParserLEFT_PARENTHESE:
		{
			p.SetState(100)
			p.Match(vinLangParserLEFT_PARENTHESE)
		}
		{
			p.SetState(101)
			p.expression(0)
		}
		{
			p.SetState(102)
			p.Match(vinLangParserRIGHT_PARENTHESE)
		}

	case vinLangParserLEFT_BRACE:
		{
			p.SetState(104)
			p.Match(vinLangParserLEFT_BRACE)
		}
		{
			p.SetState(105)
			p.Match(vinLangParserSequenceStatement)
		}
		{
			p.SetState(106)
			p.Match(vinLangParserRIGHT_BRACE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vinLangParserRULE_expression)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(110)
					_la = p.GetTokenStream().LA(1)

					if !(_la == vinLangParserMUL || _la == vinLangParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(111)
					p.expression(5)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vinLangParserRULE_expression)
				p.SetState(112)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(113)
					_la = p.GetTokenStream().LA(1)

					if !(_la == vinLangParserADD || _la == vinLangParserSUB) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(114)
					p.expression(4)
				}

			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

func (p *vinLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
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
