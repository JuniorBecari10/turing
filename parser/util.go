package parser

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func getSymbol(s string, lineNum int) (rune, *Error) {
	newErr := func(msg string) *Error { return &Error{msg, lineNum} }

	if s == "" {
		return 0, newErr("Empty string")
	}

	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError && size <= 1 {
		return 0, newErr("Invalid UTF-8 encoding")
	}

	return r, nil
}

func getDirection(s string, lineNum int) (Direction, *Error) {
	newErr := func(msg string) *Error { return &Error{msg, lineNum} }
	s = strings.ToLower(s)

	switch s {
	case "l":
		return DIR_LEFT, nil
	case "r":
		return DIR_RIGHT, nil
	case "*":
		return DIR_NOT_MOVE, nil

	default:
		return 0, newErr(fmt.Sprintf("Invalid direction: '%s'.", s))
	}
}
