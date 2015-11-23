package main

import (
	"flag"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	flags := make(map[string]*bool)

	flags["b"] = flag.Bool("b", false, "Number the non-blank output lines, starting at 1.")
	flags["e"] = flag.Bool("e", false, "Display non-printing characters, and display a dollar sign(`$') at the end of each line.")
	flags["n"] = flag.Bool("n", false, "Number the output lines, starting at 1.")
	flags["s"] = flag.Bool("s", false, "Squeeze multiple adjacent empty lines, causing the output to be single-spaced.")
	flags["t"] = flag.Bool("t", false, "Display non-printing characters, and display tab characters as `^I'.")
	flags["u"] = flag.Bool("u", false, "Disable output buffering.")

	// Parse 'em
	flag.Parse()

	// Real program stuff

	io.Copy(os.Stdout, os.Stdin)

}
