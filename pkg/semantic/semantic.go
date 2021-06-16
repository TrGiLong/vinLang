package semantic

import (
	"fmt"
	vinLang "github.com/TrGiLong/vinLang/pkg/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Symbol struct {
	Id         string
	Type       string
	IsFunction bool
}

type SymbolTable struct {
	Symbols []*Symbol
}

type Scope struct {
	Tables []*SymbolTable
}

type Semantic struct {
	Scopes *Scope
	tree   *vinLang.IProgramContext
}

func (t *SymbolTable) PushSymbol(symbol *Symbol) {
	t.Symbols = append(t.Symbols, symbol)
}

func (s *Scope) GetLastTable() *SymbolTable {
	return s.Tables[len(s.Tables)-1]
}

func (s *Scope) PushTable() *SymbolTable {
	s.Tables = append(s.Tables, &SymbolTable{})
	return s.GetLastTable()
}

func (s *Scope) PopTable() *SymbolTable {
	s.Tables = s.Tables[:len(s.Tables)-1]
	return s.GetLastTable()
}

func (s *Scope) GetSymbol(id string) *Symbol {
	scope := s.GetLastTable()

	for i := len(scope.Symbols) - 1; i >= 0; i-- {
		scope := s.Tables[i]
		for _, v := range scope.Symbols {
			if v.Id == id {
				return v
			}
		}
	}

	return nil
}

func (s *Scope) GetSymbolInLastTable(id string) *Symbol {
	scope := s.GetLastTable()

	for _, v := range scope.Symbols {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func (s *Scope) CheckSymbol(id string, line int, position int) {
	symbol := s.GetSymbol(id)
	if symbol == nil {
		panic(fmt.Sprintf(" [%d:%d]: The variable %s wasn't already declareted.",
			line,
			position,
			id,
		))
	}
}

func (s *Scope) pushSymbol(id string, _type string, isFunction bool) *Symbol {
	symbol := &Symbol{
		Id:         id,
		Type:       _type,
		IsFunction: isFunction,
	}

	table := s.GetLastTable()
	table.PushSymbol(symbol)

	return symbol
}

func (s *Semantic) Analyse() {
	antlr.ParseTreeWalkerDefault.Walk(&VinLangListener{Scopes: s.Scopes}, *s.tree)
}

func NewSemantic(tree *vinLang.IProgramContext) *Semantic {
	_analyse := &Semantic{tree: tree, Scopes: &Scope{Tables: []*SymbolTable{}}}
	_analyse.Scopes.PushTable()
	return _analyse
}

type VinLangListener struct {
	*vinLang.BasevinLangListener
	Scopes *Scope
}


func (l *VinLangListener) EnterDeclaration(ctx *vinLang.DeclarationContext) {
	symbol := l.Scopes.GetSymbolInLastTable(ctx.ID(0).GetText())
	if symbol != nil {
		panic(fmt.Sprintf(" %+v: The variable %s was declareted.",
			ctx.GetStart().GetLine(),
			symbol.Id,
		))
	}

	l.Scopes.pushSymbol(ctx.ID(0).GetText(), ctx.ID(1).GetText(), false)

	fmt.Printf("%s %s\n", ctx.ID(0).GetText(), ctx.ID(1).GetText())
}

func (l *VinLangListener) EnterAssign(ctx *vinLang.AssignContext) {
	l.Scopes.CheckSymbol(ctx.ID().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetStart())
}

func (l *VinLangListener) EnterExpression(ctx *vinLang.ExpressionContext) {
	if ctx.ID() != nil {
		fmt.Printf("EXP: %s\n", ctx.GetText())
		l.Scopes.CheckSymbol(ctx.ID().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetStart())
	}
}

func (l *VinLangListener) EnterFunctionDeclaration(ctx *vinLang.FunctionDeclarationContext) {
	l.Scopes.PushTable()
}

func (l *VinLangListener) ExitFunctionDeclaration(ctx *vinLang.FunctionDeclarationContext) {
	l.Scopes.PopTable()
}
