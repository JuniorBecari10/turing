package parser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Lex(reader io.Reader) ([]Token, []Error, error) {
    tokens := []Token{}
    errors := []Error{}
    
    scanner := bufio.NewScanner(reader)
    
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        if err := scanner.Err(); err != nil {
            return nil, nil, err
        }

        if line == "" {
            // empty / only whitespace
            continue
        }

        token, err := lexLine(line)

        if err != nil {
            errors = append(errors, *err)
            continue
        }
       
        tokens = append(tokens, token)
    }

    return tokens, errors, nil
}

func lexLine(line string) (Token, *Error) {
    // remove comments; len will never be 0.
    line = strings.SplitN(line, ";", 2)[0]
    split := strings.Split(line, " ")
    fmt.Println(line, split)
    return Token{}, nil
}
