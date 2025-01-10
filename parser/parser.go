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
	ast := Ast{&RootNode{nil, []Node{}, 0, 0, file}}
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

// sets the current node to the given node, basically entering a new scope
func (p *Parser) Enter(n Node) {
	n.CreateParent(p.node)
	p.node.Append(n)
	p.node = n
}

// appends the given node to the current node, not entering it (for vars and such)
func (p *Parser) Append(n Node) {
	n.CreateParent(p.node)
	p.node.Append(n)
}

// peeks at the next token
func (p *Parser) Peek() lexer.Token {
	return p.Toks[p.Index+1]
}

// exits a node/scope
func (p *Parser) Exit() {
	if p.node.Parent() == nil {
		logger.Exit("(Parsing) Compiler tried to escape root node of AST",
			[]string{"Line", strconv.Itoa(p.Line)},
			[]string{"Col", strconv.Itoa(p.Column)},
			[]string{"Ast", "\n" + pp.Sprintln(p.node) + "\n"},
			[]string{"File", p.Filename + logger.SLogLine(p.Filename, p.Line, "(Parsing) Compiler tried to escape root node of AST") + "\n"})
	}

	p.node = p.node.Parent()
}

// retrieves current token
func (p *Parser) Current() lexer.Token { return p.Toks[p.Index] }

// parses literally everything
func (p *Parser) Parse() {
	reading := true
	for p.Index < len(p.Toks) {
		switch p.Current().Type {
		case lexer.Word:
			p.ParseWord()
		case lexer.LBrace:
			p.Exit()
			p.Advance()
		case lexer.Eof:
			reading = false
		default:
			logger.Exit("(Parsing) Unrecognized token",
				[]string{"Line", strconv.Itoa(p.Line)},
				[]string{"Col", strconv.Itoa(p.Column)},
				[]string{"Token", p.Current().Raw},
				[]string{"Ast", "\n" + pp.Sprintln(p.node) + "\n"},
				[]string{"File", p.Filename + logger.SLogLine(p.Filename, p.Line, "(Parsing) Unrecognized token") + "\n"})
		}

		if !reading {
			break
		}
	}
}

// parses keywords, variables, and functions
func (p *Parser) ParseWord() {
	switch p.Current().Raw {
	case "func":
		var fnc *Function = p.ParseFunction()
		p.Enter(fnc)
		p.Advance()
	case "var":
		var v *Variable = p.ParseVariable()
		p.Append(v)
	case "if":
		var ifstat *If = p.ParseIf()
		p.Enter(ifstat)
		p.Advance()
	case "while":
		var wstat *While = p.ParseWhile()
		p.Enter(wstat)
		p.Advance()
	case "package":
		p.Advance()
		pkg := &Package{nil, nil, p.Line, p.Column, p.Filename, p.Current().Raw}
		p.Enter(pkg)
		p.Advance()
	default:
		var expr Expression
		p.ParseExpression(&expr)

		p.Advance()
	}
}
