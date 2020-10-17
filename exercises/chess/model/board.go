// Package model contains the gameplay logic for the game of chess
package model

import (
	"fmt"
	// "strconv"
)

type Board interface {
	fmt.Stringer
	Init()
	Move(FromCoordinate, ToCoordinate Coordinate) int
	Show() string
}

type Board8x8 struct {
	Table [][]Piece
}

func (b *Board8x8) Init() {
	for i := 0; i < 8; i++ {
		line := make([]Piece, 0)
		for j := 0; j < 8; j++ {
			line = append(line, Piece{"_"})
		}

		b.Table = append(b.Table, line)

	}
}

func (b *Board8x8) Move(FromCoordinate, ToCoordinate Coordinate) int {
	return 0
}

func (b *Board8x8) Show() string {
	return "BBBB"
}

func (b *Board8x8) String() string {
	a := ""
	for i := range b.Table {
		if i == 0{
			a+="|   | A | B | C | D | E | F | G | H |\n"
		}
		for j := range b.Table[i] {
			if j == 0{
				a += fmt.Sprintf("| %d | ", i+1)
			}
			fmt.Println(j)
			a += fmt.Sprintf(b.Table[i][j].String())
			a += " | "
		}
		a += "\n"
	}
	return a
}
