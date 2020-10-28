package state

import (
	"fmt"

	board "../board"
	coord "../coord"
	piece "../piece"
)

// IState is a state interface
type IState interface {
	fmt.Stringer
	GetCurrentBoard() board.IBoard
	GetActionNumber() int
	Move(From, To coord.ICoord) IState
}

// State8x8 is a state for 8x8 boards
type State8x8 struct {
	CurrentBoard  board.IBoard
	PreviousState *State8x8
	Player        string
	LastMove      []coord.ICoord
	ActionNumber  int
}

// GetCurrentBoard return the Current board of the state
func (s *State8x8) GetCurrentBoard() board.IBoard {
	return s.CurrentBoard
}

// GetActionNumber return the action number of the state.
// This is a move counter
func (s *State8x8) GetActionNumber() int {
	return s.ActionNumber
}

func (s *State8x8) String() string {
	return fmt.Sprintf("--------------------------------\nAction: %d\nPlayer: %s\nMove: %d\nGame:\n\n%s", s.ActionNumber, s.Player, s.LastMove, s.CurrentBoard.String())
}

// Move move a IPiece from the "From" coordinate to the "To" coordinate
func (s *State8x8) Move(From, To coord.ICoord) IState { // TODO : Divide the function -> Not SOLID
	PreviousTable := s.GetCurrentBoard().GetTable()
	CurrentTable := make([][]piece.IPiece, len(PreviousTable))
	copy(CurrentTable, PreviousTable)

	var CurrentBoard board.IBoard = &board.Board8x8{Table: CurrentTable}
	CurrentBoard.Move(From, To)
	var state IState = &State8x8{CurrentBoard: CurrentBoard, PreviousState: s, Player: "TOP", LastMove: []coord.ICoord{From, To}, ActionNumber: s.GetActionNumber() + 1}
	return state
}
