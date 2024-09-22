package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	var scanner *bufio.Scanner = bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		var scanned bool = scanner.Scan()

		if !scanned {
			return
		}

		var line string = scanner.Text()
		var lexer *lexer.Lexer = lexer.New(line)

		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
