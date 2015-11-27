package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

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

	// Read from Stdin, spit back out to Stdout
	if flag.NArg() == 0 || (flag.NArg() == 1 && flag.Arg(0) == "-") {
		io.Copy(os.Stdout, os.Stdin)
		os.Exit(0)
	}

	// Cat files
	outputString := ""
	lineNum := 1
	for _, arg := range flag.Args() {
		argBytes, err := ioutil.ReadFile(arg)
		if err != nil {
			os.Stderr.WriteString("File '" + arg + "' does not exist!\n")
			continue
		}

		argString := string(argBytes)

		// Print \t as ^I
		if *flags["t"] {
			argString = strings.Replace(argString, "\t", "^I", -1)
		}

		// Number each line
		if *flags["n"] {
			// Parse argString
			for _, line := range strings.Split(argString, "\n") {
				// Skip blank lines
				if *flags["b"] && line == "" {
					continue
				}
				// Add output
				outputString += fmt.Sprintf("\t%d  %s\n", lineNum, line)

				lineNum++
			}

		} else {
			// Just add to output
			outputString += argString
		}
	}

	// Print \n as $
	if *flags["e"] {
		outputString = strings.Replace(outputString, "\n", "$\n", -1)
	}

	fmt.Print(outputString)
}
