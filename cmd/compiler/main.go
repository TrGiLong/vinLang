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
	if args, err := parser.Parse(); err != nil {
		print(err)
	} else {
		file, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}

		lexer := Lexer.NewLexer(file)

		for {
			token := lexer.Next()
			if token.Type == Lexer.EOF {
				break
			}

			fmt.Printf("%d:%d\t%s\t%s\n", token.Position.Line, token.Position.Column, token.Type, token.Text,
			)
		}
	}
}
