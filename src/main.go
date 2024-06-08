package main

import (
	"fmt"
	"main/lexer"
	"main/token"
	"os"
)

func main() {
	readAndLex("./helloworld.monkey")
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
