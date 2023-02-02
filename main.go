package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"monkey/file"
	"github.com/akamensky/argparse"
)

func main() {
	argparser := argparse.NewParser("monkey", "Monkey Language")
	var f *os.File = argparser.File("i", "input-file", os.O_RDONLY, 0444, &argparse.Options{Required: false, Help: "Source code of monkey language. It must be *.monkey"})
	// Parse input
	err := argparser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(argparser.Usage(err))
		os.Exit(1)
	}

	if !argparser.GetArgs()[1].GetParsed() {
		repl.Start()
		os.Exit(0)
	}

	file.Start(f)
}
