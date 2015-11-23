package main

import (
	"fmt"
	"flag"
)

func main() {
	var noNewLine bool
	flag.BoolVar(&noNewLine, "n", false, "Do not print the trailing newline character.")
	flag.Parse()

	for i := len(flag.Args()) - flag.NArg(); i < flag.NArg(); i++ {
		fmt.Print(flag.Arg(i))
		if i != flag.NArg() - 1 {
			fmt.Print(" ")
		}
	}

	if !noNewLine {
		fmt.Println()
	}
}
