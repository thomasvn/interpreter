package repl

import (
	"bufio"
	"fmt"
	"io"
	"main/ast"
	"main/lexer"
	"main/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan() // Reads until encountering a newline
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		// 	fmt.Fprintf(out, "%+v\n", tok)
		// }

		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			fmt.Fprintln(out, "parser errors:")
			for _, err := range p.Errors() {
				fmt.Fprintln(out, "\t"+err)
			}
			continue
		}

		for _, stmt := range program.Statements {
			switch s := stmt.(type) {
			case *ast.LetStatement:
				if s.Value == nil {
					fmt.Fprintf(out, "LetStatement: Name=%s, Value=nil\n", s.Name.Value)
				} else {
					fmt.Fprintf(out, "LetStatement: Name=%s, Value=%s\n", s.Name.Value, s.Value.TokenLiteral())
				}
			default:
				fmt.Fprintf(out, "Unknown Statement: %T\n", stmt)
			}
		}
	}
}
