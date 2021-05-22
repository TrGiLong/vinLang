package main

import (
	"fmt"
	vinLang "github.com/TrGiLong/vinLang/pkg/antlr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/jessevdk/go-flags"
	"os"
)

type Option struct {
	Verbose      bool `short:"v" long:"verbose" description:"Verbose output"`
	WFatalErrors bool `long:"wfatal-errors" description:"This option causes the compiler to abort compilation on the first error occurred rather than trying to keep going and printing further error messages."`
}

var options Option
var parser = flags.NewParser(&options, flags.Default)

type VinLangListener struct {
	*vinLang.BasevinLangListener
}

func (l *VinLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

func (l *VinLangListener) EnterDeclaration(ctx *vinLang.DeclarationContext) {
fmt.Printf("%+v\n", ctx.GetText())
}

func
main() {
// Parse arguments
args, err := parser.Parse()
if err != nil {
panic(err)
}

input, err := antlr.NewFileStream(args[0])
if err != nil {
panic(err)
}

lexer := vinLang.NewvinLangLexer(input)
stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
parser := vinLang.NewvinLangParser(stream)
parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
parser.BuildParseTrees = true
tree := parser.Program()

antlr.ParseTreeWalkerDefault.Walk(&VinLangListener{}, tree)

os.Exit(0)
}
