package parser

// TODO: add pedantic flag to not include these sugars
type Direction int
const (
    DIR_LEFT Direction = iota
    DIR_RIGHT
    DIR_NOT_MOVE
)

type Program = map[Check]Operation
type Symbol = uint32 // any UTF-32 character

type Check struct {
    State string
    Symbol Symbol
}

type Operation struct {
    Symbol Symbol
    Direction Direction
    State string
}
