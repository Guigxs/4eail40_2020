// Package model contains the gameplay logic for the game of chess
package model

import (
	"fmt"
)

type Board interface {
	fmt.Stringer
	Init()
	Move(FromCoordinate, ToCoordinate []int) int
}

type Board8x8 struct {
	Table [][]Piece
}

func (b *Board8x8) Init() {
	for i := 0; i < 8; i++ {
		line := make([]Piece, 0)
		for j := 0; j < 8; j++ {
			line = append(line, Piece{Repr:"_"})
		}

		b.Table = append(b.Table, line)
	}

	// TODO : place the 2 teams with root
}

func (b *Board8x8) Move(FromCoordinate, ToCoordinate []int) int{
	// Add 1 to State.ActionNumber
	return 0
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
