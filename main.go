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

    program, errs := parser.Parse(tokens)
    for _, e := range errs {
    	fmt.Println(e)
    }

    fmt.Println(program)
}
