package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	model "../../model"
)

var CurrentState model.State

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

		from := []int{firstNumInt - 1, GetAlphabetToint(args[1][0])}
		to := []int{SecondNumInt - 1, GetAlphabetToint(args[2][0])}

		CurrentState = Move(CurrentState, from, to)
		fmt.Println(CurrentState)

		break
	default:
		e = fmt.Errorf("unknown command %s", args[0])
	}
	return
}

func InitGame() model.State {
	var board = model.Board8x8{}
	board.Init()
	fmt.Println(board)

	var state model.State = &model.State8x8{CurrentBoard: &board, PreviousState: nil, Player: "root", LastMove: nil, ActionNumber: 0}
	return state
}

func GetAlphabetToint(letter byte) int {
	alphabetMaj := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetMin := "abcdefghijklmnopqrstuvwxyz"

	for i := range alphabetMaj {
		if letter == alphabetMaj[i] || letter == alphabetMin[i] {
			return i
		}
	}
	return -1
}

func Move(CurrentState model.State, From, To []int) model.State { // TODO : Divide the function -> Not SOLID
	PreviousTable := CurrentState.GetCurrentBoard().GetTable()
	CurrentTable := make([][]model.IPiece, len(PreviousTable))
	copy(CurrentTable, PreviousTable)
	var CurrentBoard model.Board = &model.Board8x8{CurrentTable}
	CurrentBoard.Move(From, To)
	var state model.State = &model.State8x8{CurrentBoard: CurrentBoard, PreviousState: &CurrentState, Player: "TOP", LastMove: [][]int{From, To}, ActionNumber: CurrentState.GetActionNumber() + 1}
	return state
}
