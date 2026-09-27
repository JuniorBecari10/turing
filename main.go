package main

import (
	"fmt"
	"os"

	"turing/cli"
	"turing/parser"
	"turing/run"
)

func main() {
	args := cli.ParseFlags()
	os.Exit(perform(args))
}

func perform(args cli.CliArgs) int {
	reader, err := getReader(args.File)
	if err != nil {
		return fail(err)
	}

	tokens, err := parser.Lex(reader)
	if err != nil {
		return fail(err)
	}

	program, errs := parser.Parse(tokens)
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Printf("Error in line %d: %s\n", e.Line+1, e.Message)
		}
		return 1
	}

	runner := run.New(program, buildInitialTape(args.Tape), args.Head, args.State, args.FullSpeed)
	runner.Run()

	return 0
}
