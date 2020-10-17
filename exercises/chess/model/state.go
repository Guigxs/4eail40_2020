package model

//TODO Implement game state related elements here

type State interface{
	ShowState() string
}

type State8x8 struct{
	CurrentBoard Board8x8
	PreviousState *State8x8
	Player string
	LastMove map[CoordinateSquare]Piece
}

func (s State8x8) ShowState() string{
	return "This is the new state"
}
