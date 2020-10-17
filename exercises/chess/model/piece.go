package model

import(
	"fmt"
)

//TODO Implement game pieces here

type IPiece interface{
	fmt.Stringer
}

type Piece struct {
	Repr string // To represent on map
	Direcitons []string // line, diag, L
	Quantity int // Nomber of case the piece can move
}

func (e Piece) String() string{
	return e.Repr
}
