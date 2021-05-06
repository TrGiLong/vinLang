package main

import (
	vinLexer "github.com/TrGiLong/vinLang/pkg/lexer"
	vinParser "github.com/TrGiLong/vinLang/pkg/parser"
	"github.com/davecgh/go-spew/spew"
	"github.com/jessevdk/go-flags"
	"os"
)

type Option struct {
	Verbose      bool `short:"v" long:"verbose" description:"Verbose output"`
	WFatalErrors bool `long:"wfatal-errors" description:"This option causes the compiler to abort compilation on the first error occurred rather than trying to keep going and printing further error messages."`
}

var options Option
var parser = flags.NewParser(&options, flags.Default)

func main() {
	// Parse arguments
	args, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	// Open vl source file
	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}

	// Create lexer
	lexer := vinLexer.NewLexer(file)
	parser := vinParser.NewParser(lexer)

	ast := parser.Parse()
	spew.Dump(ast)

	os.Exit(0)
}
