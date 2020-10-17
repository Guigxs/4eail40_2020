package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	model "../../model"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(commandStr string) (e error) {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	args := strings.Fields(commandStr)
	switch args[0] {
	case "exit":
		os.Exit(0)

	// add another case here for custom commands.

	case "new":
		fmt.Println(InitGame())
		
		// TODO Create a new game on a classic 8x8 board.
		// TODO Display the board on console.

		break

	case "move":

		// TODO Move a piece. (syntax: move <from> <to>)
		// TODO The command line should be in the form of move A2 A4.
		// TODO     => meaning move piece from position A2 to A4
		// TODO Display the board on console.

		break
	default:
		e = fmt.Errorf("unknown command %s", args[0])
	}
	return
}

func InitGame() model.State{
	var board = model.Board8x8{}
	board.Init()
	var state model.State = model.State8x8{CurrentBoard:board, PreviousState:nil, Player: "root", LastMove:nil, ActionNumber:-2} 

	// var pieces8x8 = map[string]int {"Queen":1, "King":1, "Rook":2, "Knight":2, "Bishop":2, "Pawn":8,}
	// for name, qtty := range pieces8x8{
	// board.Place()

	return state
	// }
}