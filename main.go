package main

import (
	"fmt"
	"os"

	"github.com/ilya-tokar/interpreter/repl"
)

func main() {
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
