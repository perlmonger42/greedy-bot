package game

// Board represents the content of all the squares of a Ban Chi game
type Board [4][8]Piece

// NewBoard builds a Board object from a 2d array of Pao-style
// board-square state descriptions ("Q", "p", "G", ".", "?", etc).
func NewBoard(boardAsStrings [][]string) Board {
	board := [4][8]Piece{}
	for r, row := range boardAsStrings {
		for c, str := range row {
			board[r][c] = NewPiece(str)
		}
	}
	return board
}

func (board *Board) At(row, col int) Piece {
	return board[row][col]
}

func (board *Board) Each(visit func(Piece)) {
	for _, row := range board {
		for _, str := range row {
			visit(str)
		}
	}
}
