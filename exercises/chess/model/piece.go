package model

import(
	"fmt"
)

//TODO Implement game pieces here
type Coordinate interface {
	ShowCoord() string
}

type CoordinateSquare struct {
	Letter string
	Number string
}

func (c CoordinateSquare) ShowCoord() string {
	return "[" + c.Letter + "," + c.Number + "]"
}

type IPiece interface{
	fmt.Stringer
}

type Piece struct {
	Repr string
}

func (e Piece) String() string{
	return e.Repr
}

// type Queen struct {
// 	Movments []string
// 	InitialPosition CoordinateSquare
// 	CurrentPosition CoordinateSquare
// }
