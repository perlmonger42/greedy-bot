package game

import (
	"testing"
)

func Test_NewBoard_and_At(t *testing.T) {
	const _O, __ = FaceDown, None
	const rQ, bQ = RedCannon, BlackCannon
	const rP, bP = RedPawn, BlackPawn
	const rH, bH = RedHorse, BlackHorse
	const rC, bC = RedCart, BlackCart
	const rE, bE = RedElephant, BlackElephant
	const rG, bG = RedGuard, BlackGuard
	const rK, bK = RedKing, BlackKing
	expected := [4][8]Piece{
		{_O, __, _O, __, _O, bE, __, _O},
		{_O, _O, rP, _O, bH, rK, _O, rE},
		{rG, bG, bC, bQ, _O, rP, __, rQ},
		{rH, rP, _O, _O, rC, rP, bP, _O},
	}

	boardSymbols := [][]string{
		{"?", ".", "?", ".", "?", "E", ".", "?"},
		{"?", "?", "p", "?", "H", "k", "?", "e"},
		{"g", "G", "C", "Q", "?", "p", ".", "q"},
		{"h", "p", "?", "?", "c", "p", "P", "?"},
	}
	board := NewBoard(boardSymbols)

	for r := 0; r < 4; r++ {
		for c := 0; c < 8; c++ {
			if board.At(r, c) != expected[r][c] {
				t.Errorf("NewBoard failed: expected %s at %d,%d; got %s\n",
					expected[r][c].String(), r, c, board.At(r, c).String())
			}
		}
	}
}
