package main

import (
	"fmt"
	Lexer "github.com/TrGiLong/vinLang/pkg/lexer"
	"github.com/jessevdk/go-flags"
	"os"
)

type Option struct {
	Verbose []bool `short:"v" long:"verbose" description:"Verbose output"`
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
	for {
		token := lexer.Next()
		if token.Type == Lexer.EOF {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", token.Position.Line, token.Position.Column, token.Type, token.Text,
		)
	}

}
