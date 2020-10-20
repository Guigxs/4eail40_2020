package model

import(
	"fmt"
)

//TODO Implement game pieces here

type Color struct{
	Code string
}
	// var Reset "\033[0m"
	// var Red     = "\033[31m"
	// var Green   = "\033[32m"
	// var Yellow  = "\033[33m"
	// var Blue    = "\033[34m"
	// var Purple  = "\033[35m"
	// var Cyan    = "\033[36m"
	// var Gray    = "\033[37m"
	// var White   = "\033[97m"

var Reset Color = Color{Code:"\033[0m"}
var Red Color = Color{Code:"\033[31m"}
var Green Color = Color{Code:"\033[32m"}
var Yellow Color = Color{Code:"\033[33m"}
var Blue Color = Color{Code:"\033[34m"}
var Purple Color = Color{Code:"\033[35m"}
var Cyan Color = Color{Code:"\033[36m"}
var Gray Color = Color{Code:"\033[37m"}
var White Color = Color{Code:"\033[38m"}


type IPiece interface{
	fmt.Stringer
	GetInitPlace() []int
}

type Piece struct {
	InitPlace []int 
	Repr string // To represent on map
	Col Color
	Direcitons []string // line, diag, L
	Quantity int // Nomber of case the piece can move
}

func (e Piece) GetInitPlace() []int{
	return e.InitPlace
}

func (e Piece) String() string{
	return e.Col.Code + e.Repr + Reset.Code
}
