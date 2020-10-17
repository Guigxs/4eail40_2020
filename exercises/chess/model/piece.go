package model

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

type Piece interface{}

type Empty struct {
	Repr string
}

type Queen struct {
	Movments []string
	InitialPosition CoordinateSquare
	CurrentPosition CoordinateSquare
}
