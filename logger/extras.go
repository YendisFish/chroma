package logger

import (
	"bytes"
	"os"
	"strings"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/fatih/color"
)

func SLogLine(filename string, line int) string {
	ret := "\n"

	content, err := os.ReadFile(filename)
	if err != nil {
		Error("Failed to read file for error display", []string{"Runtime Error", err.Error()})
	}

	lines := strings.Split(string(content), "\n")
	lines = syntaxHighlight(lines)

	if line-5 > 0 {
		for i := line - 5; i < line-1; i++ {
			ret = ret + " " + color.New(color.FgCyan).Sprint(i+1) + " " + lines[i] + ""
		}
	} else {
		for i := 0; i < line-1; i++ {
			ret = ret + " " + color.New(color.FgCyan).Sprint(i+1) + " " + lines[i] + ""
		}
	}

	//fmt.Println(lines[line-1])
	ret = ret + color.New(color.FgRed, color.Bold).Sprint("E") + color.New(color.FgCyan).Sprint(line) + " " + lines[line-1] + ""

	if line+5 < len(lines) {
		for i := line; i < line+5; i++ {
			ret = ret + " " + color.New(color.FgCyan).Sprint(i+1) + " " + lines[i] + ""
		}
	} else {
		for i := line; i < len(lines); i++ {
			ret = ret + " " + color.New(color.FgCyan).Sprint(i+1) + " " + lines[i] + ""
		}
	}

	return ret
}

func BLogLine(filename string, line int) string {
	ret := "\n"

	content, err := os.ReadFile(filename)
	if err != nil {
		Error("Failed to read file for error display", []string{"Runtime Error", err.Error()})
	}

	lines := strings.Split(string(content), "\n")
	lines = syntaxHighlight(lines)

	if line-5 > 0 {
		for i := line - 5; i < line-1; i++ {
			ret = ret + " " + color.New(color.FgCyan).Sprint(i+1) + " " + lines[i] + ""
		}
	} else {
		for i := 0; i < line-1; i++ {
			ret = ret + " " + color.New(color.FgCyan).Sprint(i+1) + " " + lines[i] + ""
		}
	}

	//fmt.Println(lines[line-1])
	ret = ret + color.New(color.FgRed, color.Bold).Sprint("B") + color.New(color.FgCyan).Sprint(line) + " " + lines[line-1] + ""

	if line+5 < len(lines) {
		for i := line; i < line+5; i++ {
			ret = ret + " " + color.New(color.FgCyan).Sprint(i+1) + " " + lines[i] + ""
		}
	} else {
		for i := line; i < len(lines); i++ {
			ret = ret + " " + color.New(color.FgCyan).Sprint(i+1) + " " + lines[i] + ""
		}
	}

	return ret
}

func syntaxHighlight(raw []string) []string {
	style := styles.Get("onedark")
	formatter := formatters.TTY256

	ret := []string{}
	for _, v := range raw {
		lexer := lexers.Get("go")
		iterator, err := lexer.Tokenise(nil, v)
		if err != nil {
			Error("Failed to read file for error display", []string{"Runtime Error", err.Error()})
		}

		var buf bytes.Buffer
		formatter.Format(&buf, style, iterator)

		ret = append(ret, buf.String())
	}

	return ret
}
