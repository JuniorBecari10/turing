package parser

type Position struct {
    Line int
    Col int
}

type Token struct {
    Lexeme string
    Pos Position
}
