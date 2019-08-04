package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/schafer14/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to the Monkey REPL.\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
