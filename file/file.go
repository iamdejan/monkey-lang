package file

import (
	"bufio"
	"monkey/lexer"
	"monkey/parser"
	"monkey/util"
	"os"
)

func Start(f *os.File) bool {
	defer f.Close()

	out := os.Stdout
	scanner := bufio.NewScanner(f)
	code := ""
	for scanner.Scan() {
		line := scanner.Text()
		code = code + line + "\n"
	}

	l := lexer.NewLexer(code)
	p := parser.NewParser(l)

	program := p.ParseProgram()
	if len(p.Errors()) > 0 {
		util.PrintParserErrors(out, p.Errors())
		return false
	}

	out.WriteString(program.String() + "\n")


	return true
}
