package run

import (
	"fmt"
	"strings"
	"time"
	"turing/parser"
)

/*
TODO:

0 * a r 1
0 a a r 1

The second should shadow the first
*/

const TAPE_LENGTH int = 2000
const TAPE_VIEW int = 80

type Runner struct {
	program   parser.Program
	tape      [TAPE_LENGTH]rune
	head      int
	state     string
	drawn     bool
	fullSpeed bool
}

func New(program parser.Program, initialTape *[TAPE_LENGTH]rune, initialHead int, initialState string, fullSpeed bool) Runner {
	var tape [TAPE_LENGTH]rune
	if initialTape == nil {
		FillSlice(tape[:], ' ')
	} else {
		tape = *initialTape
	}

	var state string = "0"
	if initialState != "" {
		state = initialState
	}

	return Runner{
		program:   program,
		tape:      tape,
		head:      initialHead,
		state:     state,
		drawn:     false,
		fullSpeed: fullSpeed,
	}
}

func (r *Runner) Run() {
	for {
		r.draw()

		if strings.Index(r.state, "halt") == 0 {
			fmt.Println("Halted.")
			break
		}

		check := r.getCheck()
		op, ok := r.program[check]

		if !ok {
			fmt.Printf("Halted. No rule for state '%s' and symbol '%c'.\n", check.State, check.Symbol)
			break
		}

		r.doOperation(op)

		if !r.fullSpeed {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (r *Runner) getCheck() parser.Check {
	symbol := r.tape[r.head]
	if symbol == ' ' {
		symbol = '_'
	}

	return parser.Check{
		State:  r.state,
		Symbol: symbol,
	}
}

func (r *Runner) doOperation(op parser.Operation) {
	if op.Symbol == '_' {
		r.tape[r.head] = ' '
	} else if op.Symbol != '*' {
		r.tape[r.head] = op.Symbol
	}

	if op.State != "*" {
		r.state = op.State
	}

	switch op.Direction {
	case parser.DIR_LEFT:
		r.head--
		if r.head < 0 {
			r.head = TAPE_LENGTH - 1
		}

	case parser.DIR_RIGHT:
		r.head++
		if r.head == TAPE_LENGTH {
			r.head = 0
		}

	case parser.DIR_NOT_MOVE:
		{
		} // no-op
	default:
		panic("unreachable")
	}
}

func (r *Runner) draw() {
	if r.drawn {
		fmt.Printf("\033[%dA\033[J", 4) // tape line, pointer line, blank line, state line — always shown now
	}
	r.drawn = true

	// Find which TAPE_VIEW-sized window the head currently falls in.
	offset := (r.head / TAPE_VIEW) * TAPE_VIEW

	prefix := fmt.Sprintf("%4d | ", offset)
	fmt.Printf("\033[2m%s\033[0m", prefix)

	for i := 0; i < TAPE_VIEW; i++ {
		idx := (offset + i) % TAPE_LENGTH
		fmt.Printf("%c", r.tape[idx])
	}

	fmt.Println()

	pointerPos := r.head - offset
	fmt.Print(strings.Repeat(" ", len(prefix)+pointerPos))
	fmt.Println("^")

	fmt.Printf("\nState: %s\n", r.state)
}
