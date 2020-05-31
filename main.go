package main

import (
	"./halma"
	"bufio"
	"fmt"
	"os"
)

const B = 1            // Black checker flag
const W = 2            // White checker flag
const F = 3
const INFINITY = 30000 // Max value
const TURNS_LIMIT = 300

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

		//if bestMove.Indx == 6 && bestMove.Color == B {
		//	println("error")
		//}

		//if bestMove.Source == 88 && bestMove.Dest == 68 {
			//println("error")
		//}

		//board.PrintBoard(true)
		board.MakeMove(bestMove, bestMove.Indx)
		//board.PrintBoard(true)

		//for n := 0; n < 12; n++ {
		//	if board.GetCheckersPos(currentColor)[n] == bestMove.Source {
		//
		//		//if bestMove.Source == 68 && bestMove.Dest == 66 {
		//		if bestMove.Source == 88 && bestMove.Dest == 68 {
		//			println("error")
		//		}
		//		board.PrintBoard(true)
		//		board.MakeMove(bestMove, n)
		//		board.PrintBoard(true)
		//		break
		//	}
		//}

		if currentColor == W {
			currentColor = B
			MaxPly = 1
		} else {
			currentColor = W
			MaxPly = 1
		}

		fmt.Println(strategy)
		fmt.Printf("The current half-turn # %d\n", n)
		board.PrintBoard(true)
	}
}

//
//const SIZE = 10  	   // Size of board 10*10
//const B = 1            // Black checker flag
//const W = 2            // White checker flag
//const F = 3            // Outboard flag
//const INFINITY = 30000 // Max value
//const MAX_PLY = 3      // Max deep
//
//type TMove struct {
//	source int
//	dest   int
//}
//
//type TSet struct {
//	data [16]byte
//}
//
//type Targ struct {
//	score    int
//	ply      int
//	max      int
//	bestMove TMove
//	set      TSet
//}



	//printBoard(false)
	//
	//var arg Targ
	//var n int
	//memset.Memset([]byte(fmt.Sprintf("%v", arg)), 0)
	//search(&arg)
	//
	//for n = 0; n < 12; n++ {
	//	if PieceTab[n] == arg.bestMove.source {
	//		makeMove(arg.bestMove, n)
	//		break
	//	}
	//}

	//

	//
	//DelTab = [4]int{-1, 1, -SIZE, SIZE}
	//
	//printBoard(false)
	//
	//var arg Targ
	//var n int
	//memset.Memset([]byte(fmt.Sprintf("%v", arg)), 0)
	//search(&arg)
	//
	//for n = 0; n < 12; n++ {
	//	if PieceTab[n] == arg.bestMove.source {
	//		makeMove(arg.bestMove, n)
	//		break
	//	}
	//}
	//
	//printBoard(false)



//
//func clear(set *TSet) {
//	for n:=0; n<16; n++ {
//		set.data[n] = 0
//	}
//}
//
//func include(set *TSet, n int) {
//	set.data[(n)>>3] |= 1 << ((n-1) & 7)
//}
//
//func in(set TSet, n int) bool {
//	return (set.data[(n)>>3] & ( 1 << ((n-1) & 7))) != 0
//}
//
//func checkEnd() bool {
//	var n int
//	for n = 0; n < len(EndTab); n++ {
//		if Pos[EndTab[n]] != B {
//			return false
//		}
//	}
//	return true
//}
//
//func makeMove(move TMove, index int) {
//	Pos[move.source] = 0
//	Pos[move.dest] = B
//	PieceTab[index] = move.dest
//}
//
//func unMakeMove(move TMove, index int) {
//	Pos[move.dest] = 0
//	Pos[move.source] = B
//	PieceTab[index] = move.source
//}
//
//func simpleMove(arg *Targ, index int) {
//	var n int
//	var move = TMove{}
//	move.source = PieceTab[index]
//	for n = 0; n < 4; n++ {
//		move.dest = move.source + DelTab[n]
//		if Pos[move.dest] == 0 {
//			makeMove(move, index)
//			callSearch(arg, move)
//			unMakeMove(move, index)
//		}
//	}
//}
//
//func Move(arg *Targ, index int) {
//	var n int
//	var tmp int
//	var move TMove
//	for n = 0; n < 4; n++ {
//		move.source = PieceTab[index]
//		move.dest = move.source + DelTab[n]
//		tmp = Pos[move.dest]
//		if tmp == 0 || tmp == F {
//			continue
//		}
//		move.dest += DelTab[n]
//		tmp = Pos[move.dest]
//		if tmp != 0 {
//			continue
//		}
//		if in(arg.set, move.dest) {
//			continue
//		}
//		include(&arg.set, move.dest)
//		makeMove(move, index)
//		callSearch(arg, move)
//		Move(arg, index)
//		unMakeMove(move, index)
//	}
//}
//
//func callSearch(arg *Targ, move TMove) {
//	var nextArg Targ
//	nextArg.ply = arg.ply + 1
//	nextArg.score = arg.score + (StTab[move.dest] - StTab[move.source])
//	search(&nextArg)
//	//if (arg.ply == 0) {
//		//fmt.Println(nextArg)
//	//}
//	if nextArg.max > arg.max {
//		arg.max = nextArg.max
//		arg.bestMove.source = move.source
//		arg.bestMove.dest = move.dest
//
//		if nextArg.ply >= MAX_PLY {
//			//printBoard(true)
//			//fmt.Println(nextArg)
//		} else {
//			//printBoard(false)
//		}
//	}
//}
//
//func search(arg *Targ) {
//	var n int
//	if arg.ply >= MAX_PLY {
//		arg.max = arg.score
//		return
//	}
//
//	if checkEnd() {
//		arg.max = INFINITY - arg.ply
//		return
//	}
//
//	arg.max = -INFINITY
//	for n = 0; n < len(EndTab); n++ {
//		simpleMove(arg, n)
//	}
//
//	for n = 0; n < len(EndTab); n++ {
//		clear(&arg.set)
//		include(&arg.set, PieceTab[n])
//		Move(arg, n)
//	}
//}
//
//func printBoard(finished bool)  {
//	table := tablewriter.NewWriter(os.Stdout)
//	var data [SIZE][]int
//	var slice []int
//	rowIndx := 0
//	n := 0
//
//	for n=0; n<(len(Pos)-1); n+=SIZE {
//		slice = Pos[n:n+SIZE]
//		data[rowIndx] = slice
//		rowIndx++
//	}
//
//	index := 0
//	for _, row := range data {
//		strRow := make([]string, len(row))
//		var colors []tablewriter.Colors
//		for i, v := range row {
//			switch v {
//				case B:
//					colors = append(colors, tablewriter.Colors{tablewriter.FgBlueColor, tablewriter.Bold})
//				case F:
//					colors = append(colors, tablewriter.Colors{tablewriter.FgWhiteColor, tablewriter.Bold})
//				case W:
//					colors = append(colors, tablewriter.Colors{tablewriter.FgHiCyanColor, tablewriter.Bold})
//				default:
//					colors = append(colors, tablewriter.Colors{tablewriter.FgHiYellowColor, tablewriter.Bold})
//			}
//			strRow[i] = fmt.Sprintf("%d (%d)", index, StTab[index])
//			index++
//		}
//		table.Rich(strRow, colors)
//	}
//
//	var color tablewriter.Colors
//	if finished {
//		color = tablewriter.Colors{tablewriter.BgHiRedColor, tablewriter.Bold}
//		table.SetCaption(true, "========= FINISHED THE BOARD POS !!! =========")
//	} else {
//		color = tablewriter.Colors{tablewriter.Bold}
//	}
//
//	table.SetRowSeparator("=")
//	table.SetColumnSeparator("|")
//	table.SetCenterSeparator("|")
//
//	headerFooterColumns := make([]string, SIZE)
//	for i:=0; i<SIZE; i++ {
//		headerFooterColumns[i] = "========="
//	}
//	table.SetFooter(headerFooterColumns)
//	table.SetHeader(headerFooterColumns)
//
//	table.SetHeaderColor(color, color, color, color, color, color, color, color, color, color)
//	table.SetFooterColor(color, color, color, color, color, color, color, color, color, color)
//
//	table.SetBorder(true)
//	table.Render()
//}
