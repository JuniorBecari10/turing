package parser

import (
	"bufio"
	"io"
	"strings"
)

func Lex(reader io.Reader) ([]Token, error) {
    tokens := []Token{}
  
    scanner := bufio.NewScanner(reader)
    lineNum := 0
    
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        if err := scanner.Err(); err != nil {
            return nil, err
        }

        if line == "" {
            // empty / only whitespace
            lineNum += 1
            continue
        }

        // remove comments; len will never be 0.
        line = strings.SplitN(line, ";", 2)[0]

        lineTokens := lexLine(line, lineNum)
        tokens = append(tokens, lineTokens...)

        lineNum += 1
    }

    return tokens, nil
}

func lexLine(line string, lineNum int) []Token {
    split := strings.Split(line, " ")

    tokens := []Token{}
    for _, lexeme := range split {
        tokens = append(tokens, Token{
        	Lexeme: lexeme,
        	Line: lineNum,
        })
    }
    
    return tokens
}
