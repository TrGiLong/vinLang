// Code generated from ./antlr/vinLang.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // vinLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasevinLangListener is a complete listener for a parse tree produced by vinLangParser.
type BasevinLangListener struct{}

var _ vinLangListener = &BasevinLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasevinLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasevinLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasevinLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasevinLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasevinLangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasevinLangListener) ExitProgram(ctx *ProgramContext) {}

// EnterSequenceStatement is called when production sequenceStatement is entered.
func (s *BasevinLangListener) EnterSequenceStatement(ctx *SequenceStatementContext) {}

// ExitSequenceStatement is called when production sequenceStatement is exited.
func (s *BasevinLangListener) ExitSequenceStatement(ctx *SequenceStatementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasevinLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasevinLangListener) ExitStatement(ctx *StatementContext) {}

// EnterAssign is called when production assign is entered.
func (s *BasevinLangListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BasevinLangListener) ExitAssign(ctx *AssignContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasevinLangListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasevinLangListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BasevinLangListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BasevinLangListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterBoolExpression is called when production boolExpression is entered.
func (s *BasevinLangListener) EnterBoolExpression(ctx *BoolExpressionContext) {}

// ExitBoolExpression is called when production boolExpression is exited.
func (s *BasevinLangListener) ExitBoolExpression(ctx *BoolExpressionContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasevinLangListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasevinLangListener) ExitExpression(ctx *ExpressionContext) {}
