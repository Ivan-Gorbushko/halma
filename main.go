package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/tmthrgd/go-memset"
	"os"
)

const B = 1            // Black checker flag
const W = 2            // White checker flag
const F = 3            // Outboard flag
const INFINITY = 30000 // Max value
const MAX_PLY = 3      // Max deep

type TMove struct {
	source int
	dest   int
}

type TSet struct {
	data [16]byte
}

type Targ struct {
	score    int
	ply      int
	max      int
	bestMove TMove
	set      TSet
}

var Pos [10 * 10]int
var StTab [10 * 10]int
var PieceTab [12]int
var EndTab [12]int
var DelTab [4]int

func main() {
	//Pos = [10 * 10]int{
	//	F, F, F, F, F, F, F, F, F, F,
	//	F, B, B, B, B, 0, 0, 0, 0, F,
	//	F, B, B, B, B, 0, 0, 0, 0, F,
	//	F, B, B, B, B, 0, 0, 0, 0, F,
	//	F, 0, 0, 0, 0, 0, 0, 0, 0, F,
	//	F, 0, 0, 0, 0, 0, 0, 0, 0, F,
	//	F, 0, 0, 0, 0, W, W, W, W, F,
	//	F, 0, 0, 0, 0, W, W, W, W, F,
	//	F, 0, 0, 0, 0, W, W, W, W, F,
	//	F, F, F, F, F, F, F, F, F, F,
	//}

	Pos = [10 * 10]int{
		F, F, F, F, F, F, F, F, F, F,
		F, 0, 0, 0, 0, 0, 0, 0, 0, F,
		F, 0, 0, 0, 0, 0, 0, 0, 0, F,
		F, 0, 0, 0, B, 0, 0, 0, 0, F,
		F, 0, 0, 0, W, 0, 0, 0, 0, F,
		F, 0, 0, 0, 0, 0, 0, 0, 0, F,
		F, 0, 0, 0, W, 0, W, W, W, F,
		F, 0, 0, 0, 0, W, 0, 0, W, F,
		F, 0, 0, 0, 0, W, W, W, W, F,
		F, F, F, F, F, F, F, F, F, F,
	}

	StTab = [10 * 10]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 0,
		0, 9, 10, 11, 12, 13, 14, 15, 16, 0,
		0, 10, 17, 18, 19, 20, 21, 22, 23, 0,
		0, 11, 18, 24, 925, 26, 27, 28, 29, 0,
		0, 12, 19, 25, 926, 31, 32, 33, 34, 0,
		0, 13, 20, 26, 31, 401, 402, 403, 404, 0,
		0, 14, 21, 27, 32, 402, 405, 406, 407, 0,
		0, 15, 22, 28, 33, 403, 406, 408, 409, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	//EndTab = [12]int{
	//	65, 66, 67, 68,
	//	75, 76, 77, 78,
	//	85, 86, 87, 88,
	//}

	//PieceTab = [12]int{
	//	11, 12, 13, 14,
	//	21, 22, 23, 24,
	//	31, 32, 33, 34,
	//}

	PieceTab = [12]int{
		34,
	}

	EndTab = [12]int{
		88,
	}

	DelTab = [4]int{-1, 1, -10, 10}

	var arg Targ
	var n int
	memset.Memset([]byte(fmt.Sprintf("%v", arg)), 0)
	printBoard()
	search(&arg)

	fmt.Println(arg)

	for n = 0; n < 12; n++ {
		if PieceTab[n] == arg.bestMove.source {
			PieceTab[n] = arg.bestMove.dest
			Pos[arg.bestMove.source] = 0
			Pos[arg.bestMove.dest] = B
			break
		}
	}
}

func clear(set *TSet) {
	for n:=0; n<16; n++ {
		set.data[n] = 0
	}
}

func include(set *TSet, n int) {
	set.data[(n)>>3] |= 1 << ((n-1) & 7)
}

func in(set TSet, n int) bool {
	println(set.data[(n)>>3] & 1 << ((n-1) & 7))
	return (set.data[(n)>>3] & ( 1 << ((n-1) & 7))) != 0
}

func checkEnd() bool {
	var n int
	for n = 0; n < 1; n++ {
		if Pos[EndTab[n]] != B {
			return false
		}
	}
	return true
}

func makeMove(move TMove, index int) {
	Pos[move.source] = 0
	Pos[move.dest] = B
	PieceTab[index] = move.dest
}

func unMakeMove(move TMove, index int) {
	Pos[move.dest] = 0
	Pos[move.source] = B
	PieceTab[index] = move.source
}

func simpleMove(arg *Targ, index int) {
	var n int
	var move = TMove{}
	move.source = PieceTab[index]
	for n = 0; n < 4; n++ {
		move.dest = move.source + DelTab[n]
		if Pos[move.dest] == 0 {
			makeMove(move, index)
			callSearch(arg, move)
			unMakeMove(move, index)
		}
	}
}

func Move(arg *Targ, index int) {

	var n int
	var tmp int
	var move TMove
	for n = 0; n < 4; n++ {
		move.source = PieceTab[index]
		move.dest = move.source + DelTab[n]
		tmp = Pos[move.dest]
		if tmp == 0 || tmp == F {
			continue
		}
		move.dest += DelTab[n]
		tmp = Pos[move.dest]
		if tmp != 0 {
			continue
		}
		if in(arg.set, move.dest) {
			continue
		}
		include(&arg.set, move.dest)
		makeMove(move, index)
		printBoard()
		callSearch(arg, move)
		Move(arg, index)
		unMakeMove(move, index)
	}
}

func callSearch(arg *Targ, move TMove) {
	var nextArg Targ
	nextArg.ply = arg.ply + 1
	nextArg.score = arg.score + (StTab[move.dest] - StTab[move.source])
	search(&nextArg)
	if (arg.ply == 0) {
		//fmt.Println(nextArg)
	}
	if nextArg.max > arg.max {
		arg.max = nextArg.max
		arg.bestMove.source = move.source
		arg.bestMove.dest = move.dest
		//fmt.Println("test")
	}
}

func search(arg *Targ) {
	var n int
	if arg.ply >= MAX_PLY {
		arg.max = arg.score
		return
	}

	if checkEnd() {
		arg.max = INFINITY - arg.ply
		return
	}

	arg.max = -INFINITY
	for n = 0; n < 1; n++ {
		simpleMove(arg, n)
	}

	for n = 0; n < 1; n++ {
		clear(&arg.set)
		include(&arg.set, PieceTab[n])
		Move(arg, n)
	}
}

func printBoard()  {
	table := tablewriter.NewWriter(os.Stdout)
	var data [10][]int
	var slice []int
	rowIndx := 0
	n := 0

	for n=0; n<(len(Pos)-1); n+=10 {
		slice = Pos[n:n+10]
		data[rowIndx] = slice
		rowIndx++
	}

	index := 0
	for _, row := range data {
		strRow := make([]string, len(row))
		colors := []tablewriter.Colors{}
		for i, v := range row {
			switch v {
				case B:
					colors = append(colors, tablewriter.Colors{tablewriter.FgBlueColor, tablewriter.Bold})
				case F:
					colors = append(colors, tablewriter.Colors{tablewriter.FgWhiteColor, tablewriter.Bold})
				case W:
					colors = append(colors, tablewriter.Colors{tablewriter.FgHiCyanColor, tablewriter.Bold})
				default:
					colors = append(colors, tablewriter.Colors{tablewriter.FgHiYellowColor, tablewriter.Bold})
			}
			// strconv.Itoa(index)
			strRow[i] = fmt.Sprintf("%d (%d)", index, StTab[index])
			index++
		}
		table.Rich(strRow, colors)
	}

	table.SetBorder(true)
	table.Render()
}
