// Package model contains the gameplay logic for the game of chess
package model

import (
	"strconv"
	"fmt"
)

type Board interface{
	Init()
	Place()
	Move(FromCoordinate, ToCoordinate Coordinate) int
	Show() string
}

type Board8x8 struct{
	Table map[Coordinate]Piece
}

func (b Board8x8) Init(){
	b.Table = make(map[Coordinate]Piece)
	for _, elem := range([]string{"A", "B", "C", "D", "E", "F", "G", "H"}){
		for i := range([]string{"A", "B", "C", "D", "E", "F", "G", "H"}){
			var coord Coordinate = CoordinateSquare{Letter:elem, Number:strconv.Itoa(i+1)}
			b.Table[coord] = Empty{Repr:"_"}
		}
	}
	
	fmt.Println(b.Table)
}

func (b Board8x8) Place(pieces map[Coordinate]Piece){
	for key, elem := range(pieces){
		b.Table[key] = elem
	}
}

func (b Board8x8) Move(FromCoordinate, ToCoordinate Coordinate) int{
	return 0
}

func (b Board8x8) Show() string{
	return "BBBB"
}
