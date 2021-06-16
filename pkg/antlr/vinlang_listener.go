// Code generated from ./antlr/vinLang.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // vinLang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// vinLangListener is a complete listener for a parse tree produced by vinLangParser.
type vinLangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSequenceStatement is called when entering the sequenceStatement production.
	EnterSequenceStatement(c *SequenceStatementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionArgs is called when entering the functionArgs production.
	EnterFunctionArgs(c *FunctionArgsContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterBoolExpression is called when entering the boolExpression production.
	EnterBoolExpression(c *BoolExpressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSequenceStatement is called when exiting the sequenceStatement production.
	ExitSequenceStatement(c *SequenceStatementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionArgs is called when exiting the functionArgs production.
	ExitFunctionArgs(c *FunctionArgsContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitBoolExpression is called when exiting the boolExpression production.
	ExitBoolExpression(c *BoolExpressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
