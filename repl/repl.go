package repl

import (
	"bufio"
	"fmt"
	"io"
	"vore/lexer"
	"vore/token"
)

const Prompt = ">> "

func Start(in io.Reader, out io.Reader) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}