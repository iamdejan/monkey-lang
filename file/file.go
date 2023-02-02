package file

import (
	"bufio"
	"fmt"
	"monkey/lexer"
	"monkey/parser"
	"os"
)

func Start(f *os.File) {
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			for _, msg := range p.Errors() {
				fmt.Println("\t" + msg + "\t")
			}
			continue
		}
		fmt.Println(program.String())
	}
}
