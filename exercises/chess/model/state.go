package model

import (
	"fmt"
)

//TODO Implement game state related elements here

type State interface {
	fmt.Stringer
	GetCurrentBoard() Board
	GetActionNumber() int
}

type State8x8 struct {
	CurrentBoard  Board
	PreviousState *State
	Player        string
	LastMove      [][]int
	ActionNumber  int
}

func (s *State8x8) GetCurrentBoard() Board {
	return s.CurrentBoard
}

func (s *State8x8) GetActionNumber() int {
	return s.ActionNumber
}

func (s *State8x8) String() string {
	return fmt.Sprintf("--------------------------------\nAction: %d\nPlayer: %s\nMove: %d\nGame:\n\n%s", s.ActionNumber, s.Player, s.LastMove, s.CurrentBoard.String())
}
