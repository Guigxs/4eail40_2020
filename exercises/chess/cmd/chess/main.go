package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	board "../../model/board"
	coord "../../model/coord"
	state "../../model/state"
)

// CurrentState is the game state of the program.
// Write "new" to reset
var CurrentState state.IState

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
		CurrentState = InitGame()
		fmt.Println(CurrentState)

		break

	case "move":
		firstNum := string(args[1][1])
		firstNumInt, _ := strconv.Atoi(firstNum)
		SecondNum := string(args[2][1])
		SecondNumInt, _ := strconv.Atoi(SecondNum)

		var from coord.ICoord = &coord.CartesianCoord{X: string(args[1][0]), Y: firstNumInt - 1}
		var to coord.ICoord = &coord.CartesianCoord{X: string(args[2][0]), Y: SecondNumInt - 1}

		CurrentState = CurrentState.Move(from, to)
		fmt.Println(CurrentState)

		break
	default:
		e = fmt.Errorf("unknown command %s", args[0])
	}
	return
}

// InitGame initialize the game for a 8x8 board with a 8x8 state
func InitGame() state.IState {
	var initBoard = board.Board8x8{}
	initBoard.Init()

	var state state.IState = &state.State8x8{CurrentBoard: &initBoard, PreviousState: nil, Player: "root", LastMove: nil, ActionNumber: 0}
	return state
}
