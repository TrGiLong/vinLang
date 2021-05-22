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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 57, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	3, 3, 3, 7, 3, 17, 10, 3, 12, 3, 14, 3, 20, 11, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 30, 10, 5, 12, 5, 14, 5, 33, 11, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 44, 10, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 52, 10, 6, 12, 6, 14, 6, 55, 11,
	6, 3, 6, 2, 3, 10, 7, 2, 4, 6, 8, 10, 2, 4, 3, 2, 11, 12, 3, 2, 9, 10,
	2, 57, 2, 12, 3, 2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 21, 3, 2, 2, 2, 8, 23,
	3, 2, 2, 2, 10, 43, 3, 2, 2, 2, 12, 13, 5, 4, 3, 2, 13, 3, 3, 2, 2, 2,
	14, 18, 5, 6, 4, 2, 15, 17, 5, 6, 4, 2, 16, 15, 3, 2, 2, 2, 17, 20, 3,
	2, 2, 2, 18, 16, 3, 2, 2, 2, 18, 19, 3, 2, 2, 2, 19, 5, 3, 2, 2, 2, 20,
	18, 3, 2, 2, 2, 21, 22, 5, 8, 5, 2, 22, 7, 3, 2, 2, 2, 23, 24, 7, 5, 2,
	2, 24, 25, 7, 15, 2, 2, 25, 26, 7, 6, 2, 2, 26, 31, 7, 15, 2, 2, 27, 28,
	7, 8, 2, 2, 28, 30, 5, 10, 6, 2, 29, 27, 3, 2, 2, 2, 30, 33, 3, 2, 2, 2,
	31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2, 33, 31, 3,
	2, 2, 2, 34, 35, 7, 7, 2, 2, 35, 9, 3, 2, 2, 2, 36, 37, 8, 6, 1, 2, 37,
	44, 7, 13, 2, 2, 38, 44, 7, 15, 2, 2, 39, 40, 7, 3, 2, 2, 40, 41, 5, 10,
	6, 2, 41, 42, 7, 4, 2, 2, 42, 44, 3, 2, 2, 2, 43, 36, 3, 2, 2, 2, 43, 38,
	3, 2, 2, 2, 43, 39, 3, 2, 2, 2, 44, 53, 3, 2, 2, 2, 45, 46, 12, 5, 2, 2,
	46, 47, 9, 2, 2, 2, 47, 52, 5, 10, 6, 6, 48, 49, 12, 4, 2, 2, 49, 50, 9,
	3, 2, 2, 50, 52, 5, 10, 6, 5, 51, 45, 3, 2, 2, 2, 51, 48, 3, 2, 2, 2, 52,
	55, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 11, 3, 2, 2,
	2, 55, 53, 3, 2, 2, 2, 7, 18, 31, 43, 51, 53,
}
var literalNames = []string{
	"", "'('", "')'", "'bi\u1EBFn'", "':'", "';'", "'='", "'+'", "'-'", "'*'",
	"'/'",
}
var symbolicNames = []string{
	"", "", "", "VAR", "COLON", "SEMICOLIN", "ASSIGN", "ADD", "SUB", "MUL",
	"DIV", "NUMBER", "DIGIT", "ID", "CHAR", "SPACE",
}

var ruleNames = []string{
	"program", "sequenceStatement", "statement", "declaration", "expression",
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
	vinLangParserEOF       = antlr.TokenEOF
	vinLangParserT__0      = 1
	vinLangParserT__1      = 2
	vinLangParserVAR       = 3
	vinLangParserCOLON     = 4
	vinLangParserSEMICOLIN = 5
	vinLangParserASSIGN    = 6
	vinLangParserADD       = 7
	vinLangParserSUB       = 8
	vinLangParserMUL       = 9
	vinLangParserDIV       = 10
	vinLangParserNUMBER    = 11
	vinLangParserDIGIT     = 12
	vinLangParserID        = 13
	vinLangParserCHAR      = 14
	vinLangParserSPACE     = 15
)

// vinLangParser rules.
const (
	vinLangParserRULE_program           = 0
	vinLangParserRULE_sequenceStatement = 1
	vinLangParserRULE_statement         = 2
	vinLangParserRULE_declaration       = 3
	vinLangParserRULE_expression        = 4
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
		p.SetState(10)
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
		p.SetState(12)
		p.Statement()
	}
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == vinLangParserVAR {
		{
			p.SetState(13)
			p.Statement()
		}

		p.SetState(18)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Declaration()
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

func (s *DeclarationContext) SEMICOLIN() antlr.TerminalNode {
	return s.GetToken(vinLangParserSEMICOLIN, 0)
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
	p.EnterRule(localctx, 6, vinLangParserRULE_declaration)
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
		p.SetState(21)
		p.Match(vinLangParserVAR)
	}
	{
		p.SetState(22)
		p.Match(vinLangParserID)
	}
	{
		p.SetState(23)
		p.Match(vinLangParserCOLON)
	}
	{
		p.SetState(24)
		p.Match(vinLangParserID)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == vinLangParserASSIGN {
		{
			p.SetState(25)
			p.Match(vinLangParserASSIGN)
		}
		{
			p.SetState(26)
			p.expression(0)
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(vinLangParserSEMICOLIN)
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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, vinLangParserRULE_expression, _p)
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case vinLangParserNUMBER:
		{
			p.SetState(35)
			p.Match(vinLangParserNUMBER)
		}

	case vinLangParserID:
		{
			p.SetState(36)
			p.Match(vinLangParserID)
		}

	case vinLangParserT__0:
		{
			p.SetState(37)
			p.Match(vinLangParserT__0)
		}
		{
			p.SetState(38)
			p.expression(0)
		}
		{
			p.SetState(39)
			p.Match(vinLangParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vinLangParserRULE_expression)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(44)
					_la = p.GetTokenStream().LA(1)

					if !(_la == vinLangParserMUL || _la == vinLangParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(45)
					p.expression(4)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vinLangParserRULE_expression)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(47)
					_la = p.GetTokenStream().LA(1)

					if !(_la == vinLangParserADD || _la == vinLangParserSUB) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(48)
					p.expression(3)
				}

			}

		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

func (p *vinLangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 4:
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
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
