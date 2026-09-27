package main

import (
	"os"
	"turing/parser"
)

func main() {
    parser.Lex(os.Stdin)
}
