package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"monkey/file"
	"monkey/repl"
	"os"
)

const (
	CodeFile      = 0
	InputFileName = 1
)

func main() {
	argparser := argparse.NewParser("monkey", "Monkey Language")
	f := argparser.File("i", "input-file", os.O_RDONLY, 0444, &argparse.Options{Required: false, Help: "Source code of monkey language. It must be *.monkey"})

	err := argparser.Parse(os.Args)
	if err != nil {
		fmt.Print(argparser.Usage(err))
		os.Exit(1)
	}

	if argparser.GetArgs()[InputFileName].GetParsed() {
		fs := file.Start(f)
		if !fs {
			os.Exit(1)
		}
		os.Exit(0)
	}

	repl.Start()
	os.Exit(0)
}
