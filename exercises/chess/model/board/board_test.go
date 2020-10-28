package board

import (
	"testing"

	coord "../coord"
	piece "../piece"
)

func TestBoard8x8_CreateEmptyTable(t *testing.T) {
	b := Board8x8{}
	b.CreateEmptyTable()

	tests := []struct {
		name string
		b    Board8x8
		want []int
	}{
		{
			"Create empty table",
			b,
			[]int{8, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Table; len(got) != tt.want[0] && len(got[8]) != tt.want[1] {
				t.Errorf("Board8x8.CreateEmptyTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoard8x8_Place(t *testing.T) {
	b := Board8x8{}
	b.CreateEmptyTable()

	type args struct {
		Position coord.ICoord
		P        piece.IPiece
	}
	tests := []struct {
		name string
		b    Board8x8
		args args
		want string
	}{
		{
			"Place r",
			b,
			args{
				&coord.CartesianCoord{X: "A", Y: 0},
				&piece.Piece{Repr: "r"},
			},
			"r",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b.Place(tt.args.Position, tt.args.P)
			if got := tt.b.Table[tt.args.Position.GetCoord(0)][tt.args.Position.GetCoord(1)].GetRepr(); got != tt.want {
				t.Errorf("Board8x8.Place() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPieceAt(t *testing.T) {
	b := Board8x8{}
	b.CreateEmptyTable()

	tests := map[coord.ICoord]piece.IPiece{
		&coord.CartesianCoord{X: "A", Y: 1}:  &piece.Piece{Repr: "p"},
		&coord.CartesianCoord{X: "A", Y: 4}:  &piece.Piece{Repr: "_"},
		&coord.CartesianCoord{X: "C", Y: 7}:  &piece.Piece{Repr: "b"},
		&coord.CartesianCoord{X: "Z", Y: 54}: &piece.Piece{Repr: ""},
	}
	for coordi, goal := range tests {
		t.Run(coordi.String(), func(t *testing.T) {
			b.Place(coordi, goal)
			if got := b.GetPieceAt(coordi); got.GetRepr() != goal.GetRepr() {
				t.Errorf("Error expected: %s but got: %s", goal.GetRepr(), got.GetRepr())
			}
		})
	}
}

func TestBoard8x8_PlaceTeam(t *testing.T) {
	b := Board8x8{}
	b.CreateEmptyTable()

	pieces8x8 := map[string][]string{"Q": []string{"D"}, "K": []string{"E"}, "r": []string{"A", "H"}, "k": []string{"B", "G"}, "b": []string{"C", "F"}, "p": []string{"A", "B", "C", "D", "E", "F", "G", "H"}}

	type args struct {
		pieces map[string][]string
		team   string
	}
	tests := []struct {
		name string
		b    Board8x8
		args args
		want map[piece.IPiece]coord.ICoord
	}{
		{
			"Place team A",
			b,
			args{
				pieces8x8,
				"A",
			},
			map[piece.IPiece]coord.ICoord{
				&piece.Piece{Repr:"r"}: &coord.CartesianCoord{X: "A", Y: 0},
				&piece.Piece{Repr:"p"}: &coord.CartesianCoord{X: "A", Y: 1},
				&piece.Piece{Repr:"p"}: &coord.CartesianCoord{X: "B", Y: 1},
				&piece.Piece{Repr:"p"}: &coord.CartesianCoord{X: "C", Y: 1},
				&piece.Piece{Repr:"p"}: &coord.CartesianCoord{X: "A", Y: 1},
			},
		},
		{
			"Place team B",
			b,
			args{
				pieces8x8,
				"B",
			},
			map[piece.IPiece]coord.ICoord{
				&piece.Piece{Repr:"r"}: &coord.CartesianCoord{X: "A", Y: 7},
				&piece.Piece{Repr:"p"}: &coord.CartesianCoord{X: "A", Y: 6},
				&piece.Piece{Repr:"p"}: &coord.CartesianCoord{X: "B", Y: 6},
				&piece.Piece{Repr:"p"}: &coord.CartesianCoord{X: "C", Y: 6},
				&piece.Piece{Repr:"p"}: &coord.CartesianCoord{X: "A", Y: 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.PlaceTeam(tt.args.pieces, tt.args.team)
			for pieceOnBoard, coordi := range tt.want {
				if got := tt.b.GetPieceAt(coordi).GetRepr(); got != pieceOnBoard.GetRepr() {
					t.Errorf("Board8x8.PlaceTeam() = %v, want %v", got, pieceOnBoard.GetRepr())
				}
			}
		})
	}
}

func TestBoard8x8_Init(t *testing.T) {
	tests := []struct {
		name string
		b    Board8x8
	}{
		{
			"8x8",
			Board8x8{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Init()
			if tt.b.Table[0][0].GetRepr() != "r" {
				t.Errorf("Error the rook is not at [0, 0]")
			}
		})
	}
}

func TestBoard8x8_Move(t *testing.T) {
	board8 := Board8x8{}
	board8.Init()
	type args struct {
		FromCoordinate coord.ICoord
		ToCoordinate   coord.ICoord
	}
	tests := []struct {
		name string
		b    Board8x8
		args args
		want string
	}{
		{
			"A2 to A3",
			board8,
			args{
				&coord.CartesianCoord{X: "A", Y: 1},
				&coord.CartesianCoord{X: "A", Y: 2},
			},
			"p",
		},
		{
			"B4 to B5",
			board8,
			args{
				&coord.CartesianCoord{X: "B", Y: 3},
				&coord.CartesianCoord{X: "B", Y: 4},
			},
			"_",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Move(tt.args.FromCoordinate, tt.args.ToCoordinate)
			if got := tt.b.GetPieceAt(tt.args.ToCoordinate).GetRepr(); got != tt.want {
				t.Errorf("Board8x8.Move() = %v, want %v", got, tt.want)
			}
		})
	}
}
