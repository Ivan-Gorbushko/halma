package halma

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

type Board struct {
	Pos [SIZE * SIZE]int
	CellValues [SIZE * SIZE]int
	DelTab [4]int
	WhitePos [12]int
	WhiteStartPos [12]int
	WhiteEndPos [12]int
	BlackPos [12]int
	BlackStartPos [12]int
	BlackEndPos [12]int
}

func (b *Board) MakeMove(move Move, index int) {
	b.Pos[move.Source] = 0
	b.Pos[move.Dest] = move.Color
	checkersPos := b.GetCheckersPos(move.Color)
	checkersPos[index] = move.Dest
}

func (b *Board) UnMakeMove(move Move, index int) {
	b.Pos[move.Dest] = 0
	b.Pos[move.Source] = move.Color
	checkersPos := b.GetCheckersPos(move.Color)
	checkersPos[index] = move.Source
}

func (b Board) PrintBoard(finished bool)  {
	table := tablewriter.NewWriter(os.Stdout)
	var data [SIZE][]int
	var slice []int
	rowIndx := 0
	n := 0

	for n=0; n<(len(b.Pos)-1); n+=SIZE {
		slice = b.Pos[n:n+SIZE]
		data[rowIndx] = slice
		rowIndx++
	}

	index := 0
	for _, row := range data {
		strRow := make([]string, len(row))
		var colors []tablewriter.Colors
		for i, v := range row {
			switch v {
			case B:
				colors = append(colors, tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold})
			case F:
				colors = append(colors, tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold})
			case W:
				colors = append(colors, tablewriter.Colors{tablewriter.FgWhiteColor, tablewriter.Bold})
			default:
				colors = append(colors, tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Bold})
			}
			//strRow[i] = fmt.Sprintf("%d (%d)", index, b.GetCellValues(W)[index])

			val := -1
			for _, color := range []int{W, B} {
				for name, cell := range b.GetCheckersPos(color) {
					if index == cell {
						val = name
					}
				}
			}

			strRow[i] = fmt.Sprintf("%d (%d)", index, val)
			index++
		}
		table.Rich(strRow, colors)
	}

	var color tablewriter.Colors
	if finished {
		color = tablewriter.Colors{tablewriter.BgHiRedColor, tablewriter.Bold}
		table.SetCaption(true, "========= FINISHED THE BOARD POS !!! =========")
	} else {
		color = tablewriter.Colors{tablewriter.Bold}
	}

	table.SetRowSeparator("=")
	table.SetColumnSeparator("|")
	table.SetCenterSeparator("|")

	headerFooterColumns := make([]string, SIZE)
	for i:=0; i<SIZE; i++ {
		headerFooterColumns[i] = "========="
	}
	table.SetFooter(headerFooterColumns)
	table.SetHeader(headerFooterColumns)

	table.SetHeaderColor(color, color, color, color, color, color, color, color, color, color)
	table.SetFooterColor(color, color, color, color, color, color, color, color, color, color)

	table.SetBorder(true)
	table.Render()
}

func (b *Board) Create() Board {

	b.Pos = [SIZE * SIZE]int{
		F, F, F, F, F, F, F, F, F, F,
		F, B, B, B, B, 0, 0, 0, 0, F,
		F, B, B, B, B, 0, 0, 0, 0, F,
		F, B, B, B, B, 0, 0, 0, 0, F,
		F, 0, 0, 0, 0, 0, 0, 0, 0, F,
		F, 0, 0, 0, 0, 0, 0, 0, 0, F,
		F, 0, 0, 0, 0, W, W, W, W, F,
		F, 0, 0, 0, 0, W, W, W, W, F,
		F, 0, 0, 0, 0, W, W, W, W, F,
		F, F, F, F, F, F, F, F, F, F,
	}

	b.CellValues = [SIZE * SIZE]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 0,
		0, 9, 10, 11, 12, 13, 14, 15, 16, 0,
		0, 10, 17, 18, 19, 20, 21, 22, 23, 0,
		0, 11, 18, 24, 25, 26, 27, 28, 29, 0,
		0, 12, 19, 25, 30, 31, 32, 33, 34, 0,
		0, 13, 20, 26, 31, 401, 402, 403, 404, 0,
		0, 14, 21, 27, 32, 402, 405, 406, 407, 0,
		0, 15, 22, 28, 33, 403, 406, 408, 409, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}

	b.WhiteStartPos = [12]int{
		65, 66, 67, 68,
		75, 76, 77, 78,
		85, 86, 87, 88,
	}

	b.WhitePos = b.WhiteStartPos

	b.WhiteEndPos = [12]int{
		32, 12, 13, 14,
		21, 22, 23, 24,
		31, 32, 33, 34,
	}

	b.BlackStartPos = [12]int{
		11, 12, 13, 14,
		21, 22, 23, 24,
		31, 32, 33, 34,
	}

	b.BlackPos = b.BlackStartPos

	b.BlackEndPos = [12]int{
		65, 66, 67, 68,
		75, 76, 77, 78,
		85, 86, 87, 88,
	}

	b.DelTab = [4]int{-1, 1, -SIZE, SIZE}

	return *b
}

func (b *Board) GetCheckersPos(color int) *[12]int {
	if color == W {
		return &b.WhitePos
	} else {
		return &b.BlackPos
	}
}

func (b Board) GetEndPos(color int) [12]int {
	if color == W {
		return b.WhiteEndPos
	} else {
		return b.BlackEndPos
	}
}

func (b Board) GetCellValues(color int) [SIZE * SIZE]int {
	if color == W {
		for i, j := 0, len(b.CellValues)-1; i < j; i, j = i+1, j-1 {
			b.CellValues[i], b.CellValues[j] = b.CellValues[j], b.CellValues[i]
		}
		return b.CellValues
		//cellValues := b.CellValues[:]
		//var result [SIZE * SIZE]int
		//sort.Sort(sort.Reverse(sort.IntSlice(cellValues)))
		//copy(result[:], cellValues)
		//return result
	} else {
		return b.CellValues
	}
}