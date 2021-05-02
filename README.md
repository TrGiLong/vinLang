# VinLang

## Information

- Specification
  draft: [Google Docs](https://docs.google.com/document/d/1xPEmMNa3Uk88ojS0cZ2JavUKcekrdJztDJ9aVXjb9jc/edit#)
- Tutorial for writing lexer: [Link](https://www.aaronraff.dev/blog/how-to-write-a-lexer-in-go)
- Tokens
  example: [Link](https://softwareengineering.stackexchange.com/questions/328636/what-should-be-the-datatype-of-the-tokens-a-lexer-returns-to-its-parser)
- Version: v0.1

## How to use

```bash
go run  ./cmd/compiler/main.go ./vinlang_src/expression_2.vl
```

## Code guidelines

### Project structure

Follow a project structure from [here](https://github.com/golang-standards/project-layout).

- `cmd` is a main applications for this project.
- `internal` is a private application and library code.
- `pkg` is a library code that's ok to use by external applications
- `vinlang_src` contains VinLang programs.

## VinLang grammar
```
Program: SequenceStatement;

Function: ‘hàm’  ID ‘(‘ FunctionArgs ‘)’  ‘:’ <Type> ‘{‘ SequenceStatement ‘}’
FunctionArgs: FunctionArg (‘,’ FunctionArg)* 
		        | e
FunctionArg: ID ‘:’ Type

SequenceStatement: Statement (Statement)* ;

Statement: ID ‘=’ Expression ‘;’
        | ‘nếu’ Bool ‘khong_thi’ Statement ‘thi’ Statement # if
        | ‘lặp’ (VariableDeclaration; Bool; Expression)  	 # for
        | ‘dung’ ‘;’							                         # break
        | ‘tiep’ ‘;’						                           # continue
        | ‘in’ Expression ‘;’						                   # print
        | ‘{‘ SequenceStatement ‘}’					               # block code
        | VariableDeclaration ’;’					                 # declare variable
        | ‘tra’ Expression ‘;’ 						                 # return
;

VariableDeclaration: bien Id ‘:’ Type (‘=’ Expression);

Type: ‘so’ | ‘chuoi’		// number or string 

Expression: Number                                     # int
          | 'read'                                     # read
          | ID                                         # id
          | expression '*' expression                  # binOp
          | expression ('+'|'-') expression            # binOp
          | '(' expression ')'                         # expParen
          ;


Bool: (‘dung’|’sai’)                                 # boolean
    | Expression '=' Expression                      # relOp
    | Expression '<=' Expression                     # relOp
    | ‘khong’ Bool                                   # not
    | Bool 'var' Bool                                # and
    | '(' Bool ')'                                   # boolParen
    ;

Number: // Regex: ^\d*(.\d)*$

ID: ID_R ( ‘(‘ )*
ID_R// Regex: [a-z]([a-z]|\d)+

Text: ‘“‘ AnyCharacter ‘“‘;
AnyCharacter: // Regex: .+

```
