package main

import (
	"fmt"
	"io"
	"os"
	"turing/run"
)

func getReader(file *string) (io.Reader, error) {
	if file == nil {
		return os.Stdin, nil
	}

	return os.Open(*file)
}

func buildInitialTape(tape string) *[run.TAPE_LENGTH]rune {
	if tape == "" {
		return nil
	}

	initialTape := &[run.TAPE_LENGTH]rune{}
	run.FillSlice(initialTape[:], ' ')
	writeString(initialTape, tape)

	return initialTape
}

func writeString(buf *[run.TAPE_LENGTH]rune, s string) {
	i := 0

	for _, r := range s {
		if i >= len(buf) {
			break
		}

		buf[i] = r
		i++
	}
}

func fail(err error) int {
	fmt.Fprintln(os.Stderr, err.Error())
	return 1
}
