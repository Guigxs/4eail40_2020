package piece

import (
	"fmt"
)

//TODO Implement game pieces here

type Color struct {
	Code string
}

var Reset Color = Color{Code: "\033[0m"}
var Red Color = Color{Code: "\033[31m"}
var Green Color = Color{Code: "\033[32m"}
var Yellow Color = Color{Code: "\033[33m"}
var Blue Color = Color{Code: "\033[34m"}
var Purple Color = Color{Code: "\033[35m"}
var Cyan Color = Color{Code: "\033[36m"}
var Gray Color = Color{Code: "\033[37m"}
var White Color = Color{Code: "\033[38m"}

type IPiece interface {
	fmt.Stringer
	GetInitPlace() []int
	GetRepr() string
}

type Piece struct {
	InitPlace  []int
	Repr       string // To represent on map
	Col        Color
	Direcitons []string // line, diag, L
	Quantity   int      // Nomber of case the piece can move
}

func (e *Piece) GetInitPlace() []int {
	return e.InitPlace
}

func (e *Piece) GetRepr() string {
	return e.Repr
}

func (e *Piece) String() string {
	return e.Col.Code + e.Repr + Reset.Code
}
