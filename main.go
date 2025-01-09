package main

import (
	"chroma/lexer"
	"chroma/logger"
	"chroma/parser"
	"os"

	"github.com/k0kubun/pp/v3"
)

func main() {
	content, err := os.ReadFile("./example.ch")
	if err != nil {
		logger.Error("Could not read example.ch", []string{"Runtime Info", err.Error()})
	}

	toks := lexer.Tokenize(string(content))
	//pp.Println(toks)

	p := parser.Create(toks, "./example.ch")
	p.Parse()

	pp.Println(p.Ast)

	/*for _, v := range toks {
		pp.Println(v)
	}*/
}
