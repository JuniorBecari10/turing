package main

import (
	"fmt"
	"os"
	"turing/parser"
	"turing/run"
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

	var tape [run.TAPE_LENGTH]rune
	parser.FillSlice(tape[:], ' ')

	tape[0] = 'a'
	tape[1] = 'a'
	tape[2] = 'a'
	tape[3] = 'a'
	tape[4] = '-'
	tape[5] = 'a'
	tape[6] = 'a'
	tape[7] = 'a'
	tape[8] = '='

	runner := run.New(program, &tape, nil, nil)
	runner.Run()
}
