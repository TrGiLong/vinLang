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
go run  ./cmd/compiler/main.go ./vinlang_src/sum.vl
```

## Code guidelines

### Project structure

Follow a project structure from [here](https://github.com/golang-standards/project-layout).

- `cmd` is a main applications for this project.
- `internal` is a private application and library code.
- `pkg` is a library code that's ok to use by external applications
- `vinlang_src` contains VinLang programs.

