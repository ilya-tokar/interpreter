package main

import (
	"fmt"
	"github.com/ilya-tokar/interpreter/repl"
	"os"
)

func main() {
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
