package main

import (
	"fmt"
	"os"
	"turing/parser"
)

func main() {
    tokens, err := parser.Lex(os.Stdin)
    if err != nil {
    	fmt.Fprintln(os.Stderr, err.Error())
    	os.Exit(1)
    }

    program, errs := parser.Parse(tokens)
    if len(errs) > 0 {
	    for _, e := range errs {
	    	fmt.Printf("Error in line %d: %s\n", e.Line + 1, e.Message)
	    }
	    os.Exit(1)
    }

    fmt.Println(program)
}
