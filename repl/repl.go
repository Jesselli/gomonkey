package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Jesselli/gomonkey/evaluator"
	"github.com/Jesselli/gomonkey/lexer"
	"github.com/Jesselli/gomonkey/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == ".exit" {
			fmt.Println("Bye!")
			return
		}

		lexer := lexer.New(line)
		parser := parser.New(lexer)
		program := parser.ParseProgram()
		if len(parser.Errors()) > 0 {
			printParserErrors(out, parser.Errors())
			continue
		}

		result := evaluator.Eval(program)

		io.WriteString(out, result.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, err := range errors {
		io.WriteString(out, err+"\n")
	}
}
