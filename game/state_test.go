package game

import (
	"testing"
)

func TestInitialState(t *testing.T) {
	gs := NewState("Red", [][]string{
		{"?", "?", "?", "?", "?", "?", "?", "?"},
		{"?", "?", "?", "?", "?", "?", "?", "?"},
		{"?", "?", "?", "?", "?", "?", "?", "?"},
		{"?", "?", "?", "?", "?", "?", "?", "?"},
	}, []string{})
	for row := 0; row < 4; row++ {
		for col := 0; col < 8; col++ {
			if gs.Board[row][col] != FaceDown {
				t.Errorf("initial game's state should be all face-down; got %s",
					gs.Board.At(row, col).String())
			}
		}
	}
	if gs.Down[RedCannon] != 2 {
		t.Errorf("should be two red cannons face down; got %d", gs.Down[RedCannon])
	}

	if gs.Score != 0 {
		t.Errorf("initial gamestate score should be zero; got %d", gs.Score)
	}
}

func TestSingleTurnedUpPiece(t *testing.T) {
	gs := NewState("Red", [][]string{
		{"?", "?", "?", "?", "?", "?", "?", "?"},
		{"?", "?", "?", "?", "?", "?", "?", "?"},
		{"?", "?", "?", "?", "?", "?", "q", "?"},
		{"?", "?", "?", "?", "?", "?", "?", "?"},
	}, []string{})
	if gs.Board[2][6] != RedCannon {
		t.Errorf("board was constructed incorrectly: "+
			"want RedCannon but have %s\n",
			gs.Board.At(2, 6).String())
	}
	if gs.Down[RedCannon] != 1 {
		t.Errorf("should be one red cannon face down; got %d", gs.Down[RedCannon])
	}
	if gs.Score != PiecePoints[RedCannon] {
		t.Errorf("only red cannon is face up; score should be %d, but is %d\n",
			PiecePoints[RedCannon], gs.Score)
	}
	gs.Board[2][6] = FaceDown
	requireAllFaceDown(t, gs)
}

func TestOnlyKingRemains(t *testing.T) {
	gs := NewState("Red", [][]string{
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "K", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
	}, []string{
		"q", "q",
		"p", "p", "p", "p", "p",
		"h", "h", "c", "c", "e", "e", "g", "g",
		"k",
		"Q", "Q",
		"P", "P", "P", "P", "P",
		"H", "H", "C", "C", "E", "E", "G", "G",
		// "K",
	})
	if gs.Board[1][2] != BlackKing {
		t.Errorf("board was constructed incorrectly: "+
			"want BlackKing but have %s\n",
			gs.Board[1][2].String())
	}
	if gs.Score != -2*PiecePoints[RedKing] {
		t.Errorf("only black king remains; score should be %d, but is %d\n",
			-2*PiecePoints[RedKing], gs.Score)
	}
	gs.Board[1][2] = None
	requireAllEmpty(t, gs)
}

func TestEmptyBoard(t *testing.T) {
	gs := NewState("Red", [][]string{
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
	}, []string{
		"q", "q",
		"p", "p", "p", "p", "p",
		"h", "h", "c", "c", "e", "e", "g", "g",
		"k",
		"Q", "Q",
		"P", "P", "P", "P", "P",
		"H", "H", "C", "C", "E", "E", "G", "G",
		"K",
	})
	if gs.Down[RedCannon] != 0 {
		t.Errorf("should be no red cannons face down; got %d", gs.Down[RedCannon])
	}
	if gs.Score != 0 {
		t.Errorf("empty board should score zero; got %d\n", gs.Score)
	}
}

func TestPieceMix(t *testing.T) {
	// only black king face up
	// only one red horse and one black pawn face down
	// all other pieces dead
	gs := NewState("Red", [][]string{
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "K", ".", "?", ".", ".", "."},
		{".", ".", ".", ".", "?", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
	}, []string{
		"q", "q",
		"p", "p", "p", "p", "p",
		"h", "h", "c", "c", "e", "e", "g", // "g",
		"k",
		"Q", "Q",
		"P", "P", "P", "P", // "P",
		"H", "H", "C", "C", "E", "E", "G", "G",
		// "K",
	})
	if gs.Board[1][2] != BlackKing {
		t.Errorf("board was constructed incorrectly: "+
			"want BlackKing but have %s\n",
			gs.Board[1][2].String())
	}
	if gs.Board[1][4] != FaceDown {
		t.Errorf("board was constructed incorrectly: "+
			"want FaceDown but have %s\n",
			gs.Board[1][4].String())
	}
	if gs.Board[2][4] != FaceDown {
		t.Errorf("board was constructed incorrectly: "+
			"want FaceDown but have %s\n",
			gs.Board[2][4].String())
	}
	expected := 2*PiecePoints[BlackKing] +
		PiecePoints[BlackPawn] +
		PiecePoints[RedGuard]
	if gs.Score != expected {
		t.Errorf("score should be %d, but is %d\n", expected, gs.Score)
	}
	gs.Board[1][2] = None
	gs.Board[1][4] = None
	gs.Board[2][4] = None
	requireAllEmpty(t, gs)
}

func requireAllFaceDown(t *testing.T, gs State) {
	for row := 0; row < 4; row++ {
		for col := 0; col < 8; col++ {
			if p := gs.Board[row][col]; p != FaceDown {
				t.Errorf("board[%d][%d] should be FaceDown, but is %s\n",
					row, col, p.String())
			}
		}
	}
}

func requireAllEmpty(t *testing.T, gs State) {
	for row := 0; row < 4; row++ {
		for col := 0; col < 8; col++ {
			if p := gs.Board[row][col]; p != None {
				t.Errorf("board[%d][%d] should be None, but is %s\n",
					row, col, p.String())
			}
		}
	}
}
