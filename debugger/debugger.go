package debugger

import (
	"bufio"
	"chroma/logger"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/k0kubun/pp"
)

//YES IK THIS IS A HORRIBLE WAY TO DO A DEBUGGER!!!!!

var Lines = []int{5}

func Break(ast any, flename string, line int) {
	c := color.New(color.FgRed, color.Bold)

	c.Println("Breakpoint: ")
	pp.Println(ast)

	c.Println("File: ")
	str := logger.BLogLine(flename, line)
	fmt.Println(str)

	logger.Log("Press enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
