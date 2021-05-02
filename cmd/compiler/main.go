package main

import (
	"fmt"
	Lexer "github.com/TrGiLong/vinLang/pkg/lexer"
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
	lexer := Lexer.NewLexer(file)

	// Analyse
	var hasError bool = false
	for {
		token := lexer.Next()
		if token.Type == Lexer.EOF {
			break
		}

		if options.Verbose {
			fmt.Printf("%d:%d\t%s\t%s\n", token.Position.Line, token.Position.Column, token.Type, token.Text)
		}

		if token.Type == Lexer.ILLEGAL {
			hasError = true
			if options.WFatalErrors {
				break
			}
		}
	}

	if hasError {
		os.Exit(1)
	}
	os.Exit(0)
}
