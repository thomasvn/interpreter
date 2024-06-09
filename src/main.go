package main

import (
	"fmt"
	"main/lexer"
	"main/repl"
	"main/token"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language.\n", user.Username)
	fmt.Printf("Feel free to type in commands.\n")
	repl.Start(os.Stdin, os.Stdout)
}

// FOR TESTING ONLY
func readAndLex(filepath string) {
	input, _ := os.ReadFile(filepath)
	l := lexer.New(string(input))

	fmt.Printf("Lexing %s...\n", filepath)
	for {
		tok := l.NextToken()
		fmt.Printf("%s ", tok.Type)
		if tok.Type == token.EOF {
			break
		}
	}
	fmt.Printf("\n")
}
