package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/schafer14/monkey/lexer"
	"github.com/schafer14/monkey/token"
)

const PROMT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(strings.NewReader(line))

		for tok := l.Next(); tok.Type != token.EOF; tok = l.Next() {
			fmt.Printf("%v\n", tok)
		}
	}

}
