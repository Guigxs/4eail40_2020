// Package model contains the gameplay logic for the game of chess
package model

import (
	"fmt"
)

type Board interface {
	fmt.Stringer
	Init()                                       // initialize the game
	Move(FromCoordinate, ToCoordinate []int) int // Move a piece from X to Y
	GetTable() [][]IPiece
}

type Board8x8 struct {
	Table [][]IPiece
}

func (b *Board8x8) GetTable() [][]IPiece {
	return b.Table
}

func (b *Board8x8) Init() {
	for i := 0; i < 8; i++ {
		line := make([]IPiece, 0)
		for j := 0; j < 8; j++ {
			line = append(line, &Piece{Repr: "_"})
		}

		b.Table = append(b.Table, line)
	}

	var pieces8x8 = map[string][]int{"Q": []int{3}, "K": []int{4}, "r": []int{0, 7}, "k": []int{1, 6}, "b": []int{2, 5}, "p": []int{0, 1, 2, 3, 4, 5, 6, 7}}

	b.PlaceTeam(pieces8x8, "A")
	b.PlaceTeam(pieces8x8, "B")
}
func (b *Board8x8) PlaceTeam(pieces map[string][]int, team string) int {
	for key, value := range pieces {
		for _, i := range value {
			if key == "p" {
				var p IPiece
				if team == "A" {
					p = &Piece{InitPlace: []int{1, i}, Repr: key, Col: Blue, Direcitons: []string{"L", "D"}, Quantity: 1} // Use IPiece with getter/setter
				} else {
					p = &Piece{InitPlace: []int{6, i}, Repr: key, Col: Red, Direcitons: []string{"L", "D"}, Quantity: 1} // Use IPiece with getter/setter
				}

				b.Place(p.GetInitPlace(), p)
			} else {
				var p IPiece
				if team == "A" {
					p = &Piece{InitPlace: []int{0, i}, Repr: key, Col: Blue, Direcitons: []string{"L", "D"}, Quantity: 1}
				} else {
					p = &Piece{InitPlace: []int{7, i}, Repr: key, Col: Red, Direcitons: []string{"L", "D"}, Quantity: 1}
				}
				b.Place(p.GetInitPlace(), p)
			}
		}
	}
	return 1
}
func (b *Board8x8) Place(Position []int, P IPiece) int {
	for i := range b.Table {
		for j := range b.Table[i] {
			if Position[0] == i && Position[1] == j {
				b.Table[i][j] = P
			}
		}
	}
	return 1
}

func (b *Board8x8) Move(FromCoordinate, ToCoordinate []int) int {
	// Add 1 to State.ActionNumber
	fmt.Println(FromCoordinate)
	piece := b.GetPieceAt(FromCoordinate)
	fmt.Println(piece)
	if piece.GetRepr() != "_" {
		// TODO : Check with rules
		b.Table[ToCoordinate[0]][ToCoordinate[1]] = piece
		b.Table[FromCoordinate[0]][FromCoordinate[1]] = &Piece{Repr: "_"}
		return 1
	}
	
	return 0
}

func (b *Board8x8) GetPieceAt(Coordinate []int) IPiece {
	return b.Table[Coordinate[0]][Coordinate[1]]
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
