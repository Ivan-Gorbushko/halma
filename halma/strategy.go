package halma

type Strategy struct {
	MaxPly	 int
	Score    int
	Ply      int
	Max      int
	BestMove Move
	Set      Set
	Color    int
}

func (s *Strategy) checkEnd(board *Board) bool {
	var n int
	for n = 0; n < len(board.GetEndPos(s.Color)); n++ {
		if board.Pos[board.GetEndPos(s.Color)[n]] != s.Color {
			return false
		}
	}
	return true
}

func (s *Strategy) simpleMove(index int, board *Board) {
	var n int
	var move Move
	move.Source = board.GetCheckersPos(s.Color)[index]
	move.Indx = index
	move.Color = s.Color
	for n = 0; n < 4; n++ {
		move.Dest = move.Source + board.DelTab[n]
		if board.Pos[move.Dest] == 0 {
			board.MakeMove(move, index)
			s.callSearch(move, board)
			board.UnMakeMove(move, index)
		}
	}
}

func (s *Strategy) move(index int, board *Board) {
	var n int
	var tmp int
	var move Move
	move.Indx = index
	move.Color = s.Color
	for n = 0; n < 4; n++ {
		move.Source = board.GetCheckersPos(s.Color)[index]
		move.Dest = move.Source + board.DelTab[n]
		tmp = board.Pos[move.Dest]
		if tmp == 0 || tmp == F {
			continue
		}
		move.Dest += board.DelTab[n]
		tmp = board.Pos[move.Dest]
		if tmp != 0 {
			continue
		}
		if s.Set.in(move.Dest) {
			continue
		}
		s.Set.include(move.Dest)
		board.MakeMove(move, index)
		s.callSearch(move, board)
		s.move(index, board)
		board.UnMakeMove(move, index)
	}
}

func (s *Strategy) callSearch(move Move, board *Board) {
	var nextStrategy Strategy
	nextStrategy.Color = s.Color
	nextStrategy.Ply = s.Ply + 1
	nextStrategy.Score = s.Score + (board.GetCellValues(s.Color)[move.Dest] - board.GetCellValues(s.Color)[move.Source])
	nextStrategy.Search(board)

	if nextStrategy.Max > s.Max {
		//board.PrintBoard(false)
		s.Max = nextStrategy.Max
		s.BestMove.Indx = move.Indx
		s.BestMove.Color = move.Color
		s.BestMove.Source = move.Source
		s.BestMove.Dest = move.Dest
	}
}

func (s *Strategy) Search(board *Board) {
	var n int
	if s.Ply >= s.MaxPly {
		s.Max = s.Score
		return
	}

	if s.checkEnd(board) {
		s.Max = INFINITY - s.Ply
		return
	}

	s.Max = -INFINITY
	for n = 0; n < len(board.GetEndPos(s.Color)); n++ {
		s.simpleMove(n, board)
	}

	for n = 0; n < len(board.GetEndPos(s.Color)); n++ {
		s.Set.clear()
		s.Set.include(board.GetCheckersPos(s.Color)[n])
		s.move(n, board)
	}
}
