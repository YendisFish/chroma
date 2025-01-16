package main

import (
	"chroma/lexer"
	"chroma/logger"
	"chroma/parser"
	"os"

	"github.com/k0kubun/pp/v3"
)

var debug = false

func main() {
	handleArgs()

	content, err := os.ReadFile("./example.ch")
	if err != nil {
		logger.Error("Could not read example.ch", []string{"Runtime Info", err.Error()})
	}

	toks := lexer.Tokenize(string(content))
	//pp.Println(toks)

	p := parser.Create(toks, "./example.ch", debug)
	p.Parse()

	pp.Println(p.Ast)
	/*for _, v := range toks {
		pp.Println(v)
	}*/
}

func handleArgs() {
	args := os.Args[1:]
	for i, v := range args {
		switch v {
		case "debug":
			debug = true

			_ = i
		}
	}
}
