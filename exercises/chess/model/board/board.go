// Package board model contains the gameplay logic for the game of chess
package board

import (
	"fmt"

	coord "../coord"
	piece "../piece"
)

// IBoard is the board interface
type IBoard interface {
	fmt.Stringer
	CreateEmptyTable()
	Init()                                              // initialize the game
	Move(FromCoordinate, ToCoordinate coord.ICoord) int // Move a piece from X to Y
	GetTable() [][]piece.IPiece
}

// Board8x8 is the IBoard implementation with a 8x8 array
type Board8x8 struct {
	Table [][]piece.IPiece
}

// GetTable return the board array
func (b *Board8x8) GetTable() [][]piece.IPiece {
	return b.Table
}

// CreateEmptyTable create an empty table of 8x8
func (b *Board8x8) CreateEmptyTable() {
	for i := 0; i < 8; i++ {
		line := make([]piece.IPiece, 0)
		for j := 0; j < 8; j++ {
			line = append(line, &piece.Piece{Repr: "_"})
		}

		b.Table = append(b.Table, line)
	}
}

// Place place an IPiece on the board
func (b *Board8x8) Place(Position coord.ICoord, P piece.IPiece) int {
	if Position.GetCoord(0) < 8 && Position.GetCoord(1) < 8 {
		for i := range b.Table {
			for j := range b.Table[i] {
				if Position.GetCoord(0) == i && Position.GetCoord(1) == j {
					b.Table[i][j] = P
				}
			}
		}
		return 1
	} else {
		return -1
	}
}

// GetPieceAt return the piece at the given coordinate
func (b *Board8x8) GetPieceAt(Coordinate coord.ICoord) piece.IPiece {
	if Coordinate.GetCoord(0) < 8 && Coordinate.GetCoord(1) < 8 {
		return b.Table[Coordinate.GetCoord(0)][Coordinate.GetCoord(1)]
	} else {
		return &piece.Piece{Repr: ""}
	}
}

// PlaceTeam place a team on the board
func (b *Board8x8) PlaceTeam(pieces map[string][]string, team string) int {
	for key, value := range pieces {
		for _, i := range value {
			if key == "p" {
				var p piece.IPiece
				if team == "A" {
					p = &piece.Piece{InitPlace: &coord.CartesianCoord{i, 1}, Repr: key, Col: piece.Blue, Direcitons: []string{"L", "D"}, Quantity: 1} // Use IPiece with getter/setter
				} else {
					p = &piece.Piece{InitPlace: &coord.CartesianCoord{i, 6}, Repr: key, Col: piece.Red, Direcitons: []string{"L", "D"}, Quantity: 1} // Use IPiece with getter/setter
				}

				b.Place(p.GetInitPlace(), p)
			} else {
				var p piece.IPiece
				if team == "A" {
					p = &piece.Piece{InitPlace: &coord.CartesianCoord{i, 0}, Repr: key, Col: piece.Blue, Direcitons: []string{"L", "D"}, Quantity: 1}
				} else {
					p = &piece.Piece{InitPlace: &coord.CartesianCoord{i, 7}, Repr: key, Col: piece.Red, Direcitons: []string{"L", "D"}, Quantity: 1}
				}
				b.Place(p.GetInitPlace(), p)
			}
		}
	}
	return 1
}

// Init initialize the board by placing the 2 teams pieces
func (b *Board8x8) Init() {
	b.CreateEmptyTable()

	var pieces8x8 = map[string][]string{"Q": []string{"D"}, "K": []string{"E"}, "r": []string{"A", "H"}, "k": []string{"B", "G"}, "b": []string{"C", "F"}, "p": []string{"A", "B", "C", "D", "E", "F", "G", "H"}}

	b.PlaceTeam(pieces8x8, "A")
	b.PlaceTeam(pieces8x8, "B")
}

// Move move a IPiece from a ICoord to an ICoord
func (b *Board8x8) Move(FromCoordinate, ToCoordinate coord.ICoord) int {
	// Add 1 to State.ActionNumber
	pieceFrom := b.GetPieceAt(FromCoordinate)
	if pieceFrom.GetRepr() != "_" {
		// TODO : Check with rules
		b.Table[ToCoordinate.GetCoord(0)][ToCoordinate.GetCoord(1)] = pieceFrom
		b.Table[FromCoordinate.GetCoord(0)][FromCoordinate.GetCoord(1)] = &piece.Piece{Repr: "_"}
		return 1
	}

	return -1
}

func (b *Board8x8) String() string {
	a := ""
	for i := range b.Table {
		if i == 0 {
			a += "|   | A | B | C | D | E | F | G | H |\n"
		}
		for j := range b.Table[i] {
			if j == 0 {
				a += fmt.Sprintf("| %d | ", i+1)
			}
			a += fmt.Sprintf(b.Table[i][j].String())
			a += " | "
		}
		a += "\n"
	}
	return a
}
