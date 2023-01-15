package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 50; i++ {
		fmt.Println("")
	}
	fmt.Printf("Hello to %s! This is the Monkey Programming Language from \"Writing An Interpreter in Go\"\n", user.Name)
	fmt.Println("Feel free to try!")
	repl.Start(os.Stdin, os.Stdout)
}
