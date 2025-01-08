package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Error(text string, info ...[]string) {
	col := color.New(color.FgRed, color.Bold)

	col.Print("Error: ")
	fmt.Println(text)

	for _, v := range info {
		if len(v) != 2 {
			Error("Error info must have 2 elements in array!")
		}

		fmt.Print("    ")
		col.Print(v[0] + ": ")
		fmt.Println(v[1])
	}
}

func Exit(text string, info ...[]string) {
	col := color.New(color.FgRed, color.Bold)

	col.Print("Error: ")
	fmt.Println(text)

	for _, v := range info {
		if len(v) != 2 {
			Error("Error info must have 2 elements in array!")
		}

		fmt.Print("    ")
		col.Print(v[0] + ": ")
		fmt.Println(v[1])
	}

	os.Exit(1)
}

func Warn(text string, info ...[]string) {
	col := color.New(color.FgYellow, color.Bold)

	col.Print("Warning: ")
	fmt.Println(text)

	for _, v := range info {
		if len(v) != 2 {
			Error("Warn info must have 2 elements in array!")
		}

		fmt.Print("    ")
		col.Print(v[0] + ": ")
		fmt.Println(v[1])
	}
}

func Log(text string, info ...[]string) {
	col := color.New(color.FgGreen, color.Bold)

	col.Print("Info: ")
	fmt.Println(text)

	for _, v := range info {
		if len(v) != 2 {
			Error("Log info must have 2 elements in array!")
		}

		fmt.Print("    ")
		col.Print(v[0] + ": ")
		fmt.Println(v[1])
	}
}

func Custom(text string, clr color.Attribute, info ...[]string) {
	col := color.New(clr, color.Bold)

	col.Print("Info: ")
	fmt.Println(text)

	for _, v := range info {
		if len(v) != 2 {
			Error("Log info must have 2 elements in array!")
		}

		fmt.Print("    ")
		col.Print(v[0] + ": ")
		fmt.Println(v[1])
	}
}
