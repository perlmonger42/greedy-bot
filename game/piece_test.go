package game

import (
	"testing"
)

func TestNewPiece(t *testing.T) {
	expected := map[string]Piece{
		" ": None,
		".": None,
		"?": FaceDown,
		"q": RedCannon,
		"p": RedPawn,
		"h": RedHorse,
		"c": RedCart,
		"e": RedElephant,
		"g": RedGuard,
		"k": RedKing,
		"Q": BlackCannon,
		"P": BlackPawn,
		"H": BlackHorse,
		"C": BlackCart,
		"E": BlackElephant,
		"G": BlackGuard,
		"K": BlackKing,
	}
	for name, want := range expected {
		if have := NewPiece(name); have != want {
			t.Errorf("%s should make a %s, but made a %s instead\n",
				name, want, have)
		}
	}
}

func TestNewPieceFailure(t *testing.T) {
	defer func() {
		expected := "unknown piece descriptor: \"x\""
		if r := recover(); r == nil {
			t.Errorf("NewPiece(\"x\") should panic")
		} else if r != expected {
			t.Errorf("NewPiece(\"x\") should panic %q; got %q", expected, r)
		}
	}()
	_ = NewPiece("x")
}

func TestCanTakeIfAdjacent(t *testing.T) {
	// expected[i][j] tells whether
	// team1.QPHCEGK[i] can take team2.QPHCEGK[j]
	// if the two are adjacent
	const T = true
	const F = false
	expected := [7][7]bool{
		//       Q  P  H  C  E  G  K
		/* Q */ {F, F, F, F, F, F, F}, // cannon can't take _ANY_ adjacent
		/* P */ {F, T, F, F, F, F, T},
		/* H */ {T, T, T, F, F, F, F},
		/* C */ {T, T, T, T, F, F, F},
		/* E */ {T, T, T, T, T, F, F},
		/* G */ {T, T, T, T, T, T, F},
		/* K */ {T, F, T, T, T, T, T},
	}
	team1, team2 := &RedTeam, &BlackTeam
	// twice: 1st run, check red attacking black; 2nd run, black attacking red
	for twice := 0; twice < 2; twice++ {
		for i, attackerType := range team1.QPHCEGK {
			attackerPiece := attackerType
			for j, defenderType := range team2.QPHCEGK {
				defenderPiece := defenderType
				takes := attackerPiece.CanTakeIfAdjacent(defenderPiece)
				if takes != expected[i][j] {
					t.Errorf("expected %v for 'can %s take %s?'\n",
						expected[i][j],
						attackerType.String(), defenderType.String())
				}
			}
		}
		team1, team2 = team2, team1 // swap sides for the next test
	}
}

func TestNonPiecesCanTakeIfAdjacent(t *testing.T) {
	for _, nonpiece := range []Piece{None, FaceDown} {
		for _, piece := range RedTeam.QPHCEGK {
			if nonpiece.CanTakeIfAdjacent(piece) {
				t.Errorf("%s shouldn't be able to take %s\n",
					nonpiece.String(), piece.String())
			}
			if piece.CanTakeIfAdjacent(nonpiece) {
				t.Errorf("%s shouldn't be able to take %s\n",
					piece.String(), nonpiece.String())
			}
		}
	}
}
