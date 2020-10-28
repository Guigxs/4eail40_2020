package state

import (
	"fmt"

	board "../board"
	piece "../piece"
)

//TODO Implement game state related elements here

type IState interface {
	fmt.Stringer
	GetCurrentBoard() board.IBoard
	GetActionNumber() int
	Move(From, To []int) IState
}

type State8x8 struct {
	CurrentBoard  board.IBoard
	PreviousState *State8x8
	Player        string
	LastMove      [][]int
	ActionNumber  int
}

func (s *State8x8) GetCurrentBoard() board.IBoard {
	return s.CurrentBoard
}

func (s *State8x8) GetActionNumber() int {
	return s.ActionNumber
}

func (s *State8x8) String() string {
	return fmt.Sprintf("--------------------------------\nAction: %d\nPlayer: %s\nMove: %d\nGame:\n\n%s", s.ActionNumber, s.Player, s.LastMove, s.CurrentBoard.String())
}

func (s *State8x8) Move(From, To []int) IState { // TODO : Divide the function -> Not SOLID
	PreviousTable := s.GetCurrentBoard().GetTable()
	CurrentTable := make([][]piece.IPiece, len(PreviousTable))
	copy(CurrentTable, PreviousTable)

	var CurrentBoard board.IBoard = &board.Board8x8{CurrentTable}
	CurrentBoard.Move(From, To)
	var state IState = &State8x8{CurrentBoard: CurrentBoard, PreviousState: s, Player: "TOP", LastMove: [][]int{From, To}, ActionNumber: s.GetActionNumber() + 1}
	return state
}
