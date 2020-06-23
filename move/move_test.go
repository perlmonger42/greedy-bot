package move

import (
	"testing"

	"github.com/perlmonger42/greedy-bot/game"
)

func TestNoMovesPossible(t *testing.T) {
	gs := game.NewState("Black", [][]string{
		{"?", "H", "E", "G", "?", "?", "?", "?"},
		{"?", "?", "C", "?", "?", "?", "?", "p"},
		{"?", "?", "p", "?", "?", "c", "?", "P"},
		{"?", "?", "?", "?", "?", "?", "?", "?"},
	}, []string{})
	moves := LegalMoves(gs.Us, gs.Them, gs.Board)

	have_m1 := moves[len(moves)-2].String()
	want_m1 := "BlackCart at C2 takes RedPawn at C3"
	if have_m1 != want_m1 {
		t.Errorf("expected %s; got %s", want_m1, have_m1)
	}

	have_m2 := moves[len(moves)-1].String()
	want_m2 := "BlackPawn at H3 takes RedPawn at H2"
	if have_m2 != want_m2 {
		t.Errorf("expected %s; got %s", want_m2, have_m2)
	}
}
