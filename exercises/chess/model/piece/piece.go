package piece

import (
	"fmt"

	coord "../coord"
)

//TODO Implement game pieces here

// Color struct for setting a color to a player
type Color struct {
	Code string
}

// Reset color
var Reset Color = Color{Code: "\033[0m"}

// Red color
var Red Color = Color{Code: "\033[31m"}

// Green color
var Green Color = Color{Code: "\033[32m"}

// Yellow color
var Yellow Color = Color{Code: "\033[33m"}

// Blue color
var Blue Color = Color{Code: "\033[34m"}

// Purple color
var Purple Color = Color{Code: "\033[35m"}

// Cyan color
var Cyan Color = Color{Code: "\033[36m"}

// Gray color
var Gray Color = Color{Code: "\033[37m"}

// White color
var White Color = Color{Code: "\033[38m"}

// IPiece interface for pieces
type IPiece interface {
	fmt.Stringer
	GetInitPlace() coord.ICoord
	GetRepr() string
}

// Piece is a IPiece implementation
type Piece struct {
	InitPlace  coord.ICoord
	Repr       string // To represent on map
	Col        Color
	Direcitons []string // line, diag, L
	Quantity   int      // Nomber of case the piece can move
}

// GetInitPlace return the initial place of a piece
func (e *Piece) GetInitPlace() coord.ICoord {
	return e.InitPlace
}

// GetRepr return the string of the piece
func (e *Piece) GetRepr() string {
	return e.Repr
}

func (e *Piece) String() string {
	return e.Col.Code + e.Repr + Reset.Code
}
