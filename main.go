package main

import (
	"fmt"
	"os"
	"turing/parser"
)

func main() {
    tokens, err := parser.Lex(os.Stdin)
    if err != nil {
    	panic(err)
    }

    fmt.Println(tokens)
}
