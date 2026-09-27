package parser

import (
	"fmt"
)

func Parse(tokens [][]Token) (Program, []Error) {
    program := Program{}
    errors := []Error{}

    for i, line := range tokens {
        check, operation, err := parseLine(line, i)

        if err != nil {
            errors = append(errors, *err)
            continue
        }

        if _, ok := program[check]; ok {
            errors = append(errors, Error{
            	Message: fmt.Sprintf("There is already a rule with state '%s' and symbol '%c'.", check.State, check.Symbol),
            	Line: i,
            })
            continue
        }

        program[check] = operation
    }

    return program, errors
}

// TODO: test whitespace as symbols (lexemes)
func parseLine(line []Token, lineNum int) (Check, Operation, *Error) {
    if len(line) != 5 {
        return Check{}, Operation{}, &Error{
        	Message: fmt.Sprintf("Invalid number of entries per line. Expected 5, got %d.", len(line)),
        	Line: lineNum,
        }
    }

    checkSymbol, err := GetSymbol(line[1].Lexeme, lineNum)
    if err != nil {
        return Check{}, Operation{}, err
    }

    check := Check{
    	State: line[0].Lexeme,
    	Symbol: checkSymbol,
    }

    opSymbol, err := GetSymbol(line[2].Lexeme, lineNum)
    if err != nil {
        return Check{}, Operation{}, err
    }

    opDir, err := GetDirection(line[3].Lexeme, lineNum)
    if err != nil {
        return Check{}, Operation{}, err
    }

    operation := Operation{
    	Symbol: opSymbol,
    	Direction: opDir,
    	State: line[4].Lexeme,
    }

    return check, operation, nil
}
