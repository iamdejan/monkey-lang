package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"monkey/util"
	"os"
	"os/user"
)

const PROMPT = ">> "

func Start() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 50; i++ {
		fmt.Println("")
	}
	fmt.Printf("Hello to %s! This is the Monkey Programming Language from \"Writing An Interpreter in Go\"\n", user.Name)
	fmt.Println("Feel free to try!")
	start(os.Stdin, os.Stdout)
}

func start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		program := p.ParseProgram()
		if len(p.Errors()) > 0 {
			util.PrintParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}
