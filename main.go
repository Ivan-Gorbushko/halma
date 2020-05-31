package main

import (
	"./halma"
	"bufio"
	"fmt"
	"os"
)

const B = 1            // Black checker flag
const W = 2            // White checker flag
const INFINITY = 30000 // Max value
const TURNS_LIMIT = 400

var colorNames = [3]string{"NUN", "Black", "White"}

func main() {
	var board halma.Board
	var bestMove halma.Move
	board = board.Create()

	board.PrintBoard(false)

	MaxPly := 1
	currentColor := W
	for n := 0; n < TURNS_LIMIT ; n++ {

		fmt.Printf("Press any key to continue or esc to exit...")
		// only read single characters, the rest will be ignored!!
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		input, _ := consoleReader.ReadByte()
		// ESC = 27 and Ctrl-C = 3
		if input == 27 || input == 3 {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		strategy := halma.Strategy{}
		//memset.Memset([]byte(fmt.Sprintf("%v", strategy)), 0)
		strategy.MaxPly = MaxPly
		strategy.Color = currentColor
		strategy.Search(&board)
		if strategy.Max == INFINITY {
			board.PrintBoard(false)
			fmt.Printf("%s WON!!!", colorNames[strategy.Color])
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		bestMove = strategy.BestMove
		bestMove.Source = board.GetCheckersPos(bestMove.Color)[bestMove.Indx]

		board.MakeMove(bestMove, bestMove.Indx)

		if currentColor == W {
			currentColor = B
			MaxPly = 1
		} else {
			currentColor = W
			MaxPly = 1
		}

		fmt.Println(strategy)
		fmt.Println(bestMove)
		fmt.Printf("The current half-turn # %d (The %s move)\n", n, colorNames[strategy.Color])
		board.PrintBoard(true)
	}
	board.PrintBoard(false)
}
