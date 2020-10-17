package model

import (
	"fmt"
)
//TODO Implement game state related elements here

type State interface{
	fmt.Stringer
	ShowState() string
}

type State8x8 struct{
	CurrentBoard Board8x8
	PreviousState *State8x8
	Player string
	LastMove map[CoordinateSquare]Piece
	ActionNumber int
}

func (s State8x8) ShowState() string{
	return "This is the new state"
}

func (s State8x8) String() string{
	return fmt.Sprintf("--------------------------------\nAction: %d\nPlayer: %s\nMove: %s\nGame:\n\n%s", s.ActionNumber, s.Player, s.LastMove, s.CurrentBoard.String())
}
