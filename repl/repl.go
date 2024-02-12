package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Jesselli/gomonkey/lexer"
	"github.com/Jesselli/gomonkey/token"
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

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
