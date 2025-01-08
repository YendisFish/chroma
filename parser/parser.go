package parser

import (
	"chroma/debugger"
	"chroma/lexer"
	"chroma/logger"
	"strconv"

	"github.com/k0kubun/pp"
)

type ParserState int

const (
	Func ParserState = iota
	Strct
)

type Parser struct {
	Toks     []lexer.Token
	Index    int
	Line     int
	Column   int
	Filename string
	Ast      Ast
	node     Node
}

func Create(toks []lexer.Token, file string) *Parser {
	ast := Ast{&RootNode{nil, []Node{}}}
	return &Parser{toks, 0, 1, 1, file, ast, ast.Root}
}

func (p *Parser) CurrentNode() Node { return p.node }

func (p *Parser) Advance() {
	p.Index++
	p.Line = p.Toks[p.Index].Line
	p.Column = p.Toks[p.Index].Col

	for _, v := range debugger.Lines {
		if p.Line == v {
			debugger.Break(p.node, p.Filename, p.Line)
		}
	}
}

func (p *Parser) Enter(n Node) {
	n.CreateParent(p.node)
	p.node.Append(n)
	p.node = n
}

func (p *Parser) Peek() lexer.Token {
	return p.Toks[p.Index+1]
}

func (p *Parser) Exit() {
	if p.node.Parent() == nil {
		logger.Exit("(Parsing) Compiler tried to escape root node of AST",
			[]string{"Line", strconv.Itoa(p.Line)},
			[]string{"Col", strconv.Itoa(p.Column)},
			[]string{"Ast", "\n" + pp.Sprintln(p.node) + "\n"},
			[]string{"File", p.Filename + logger.SLogLine(p.Filename, p.Line) + "\n"})
	}

	p.node = p.node.Parent()
}

func (p *Parser) Current() lexer.Token { return p.Toks[p.Index] }

func (p *Parser) Parse() {
	reading := true
	for ; p.Index < len(p.Toks); p.Advance() {
		switch p.Current().Type {
		case lexer.Word:
			p.ParseWord()
		case lexer.LBrace:
			p.Exit()

		case lexer.Eof:
			reading = false

		default:
			logger.Exit("(Parsing) Unrecognized token",
				[]string{"Line", strconv.Itoa(p.Line)},
				[]string{"Col", strconv.Itoa(p.Column)},
				[]string{"Token", p.Current().Raw},
				[]string{"Ast", "\n" + pp.Sprintln(p.node) + "\n"},
				[]string{"File", p.Filename + logger.SLogLine(p.Filename, p.Line) + "\n"})
		}

		if !reading {
			break
		}
	}
}

func (p *Parser) ParseWord() {
	switch p.Current().Raw {
	case "func":
		var fnc *Function = p.ParseFunction()
		p.Enter(fnc)
	}
}
