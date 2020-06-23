// Implement a Pao bot that tries all possible next moves, evaluating each
// resulting board, and makes the move that produces the best materiel score.
package bot

import (
	"fmt"
	"math/rand"

	"github.com/perlmonger42/greedy-bot/game"
	"github.com/perlmonger42/greedy-bot/move"
)

type GreedyBot struct{}

func NewGreedyBot() GreedyBot {
	return GreedyBot{}
}

func (bot GreedyBot) Name() string {
	return "Greedy"
}

func (bot GreedyBot) ChooseMove(state *game.State) move.T {
	fmt.Printf("time to choose a move; state is %v\n", *state)
	maxer := NewMaximizer(state)
	move := maxer.BestMove()
	fmt.Printf("best move is %s\n", move.String())
	return move
}

type Maximizer struct {
	gs                  *game.State
	redBlackMultiplier  int
	bestDelta           int
	bestMove            move.T
	flipScoreCalculated bool
	flipScore           int
}

func NewMaximizer(gs *game.State) *Maximizer {
	// make scores positive for us, negative for them *Maximizer
	mult := 1
	if gs.Us != nil && gs.Us.K == game.BlackKing {
		mult = -1
	}

	return &Maximizer{
		gs:                 gs,
		redBlackMultiplier: mult,
	}
}

func (maxer *Maximizer) BestMove() move.T {
	moves := move.LegalMoves(maxer.gs.Us, maxer.gs.Them, maxer.gs.Board)
	//fmt.Printf("found %d possible moves\n", len(moves))
	bestMoves := []move.T{move.NewQuit()}
	bestDelta := -1000000
	for _, m := range moves {
		delta := maxer.scoreDelta(m)
		fmt.Printf("%d for %s\n", delta, m.String())
		if delta > bestDelta {
			//fmt.Printf("Found new best score: %d for %s\n", delta, m.String())
			bestDelta = delta
			bestMoves = []move.T{m}
		} else if delta == bestDelta {
			//fmt.Printf("Found another at best score: %d for %s\n", delta, m.String())
			bestMoves = append(bestMoves, m)
		} else {
			//fmt.Printf("Found worse move: %d for %s\n", delta, m.String())
		}
	}
	return bestMoves[rand.Intn(len(bestMoves))]
}

func (maxer *Maximizer) scoreDelta(m move.T) int {
	switch m.Action() {
	case move.Quit:
		return -1000000
	case move.Flip:
		return maxer.computeFlipScore()
	case move.Move:
		return 0
	case move.Take:
		return -maxer.redBlackMultiplier * game.PiecePoints[m.Killed()]
	}
	return 0
}

func (maxer *Maximizer) computeFlipScore() int {
	if !maxer.flipScoreCalculated {
		pointSum, pieceCount := 0, 0
		for piece, count := range maxer.gs.Down {
			pieceCount += count
			pointSum += count * game.PiecePoints[piece]
		}
		maxer.flipScore = maxer.redBlackMultiplier * pointSum / pieceCount
		maxer.flipScoreCalculated = true
	}
	return maxer.flipScore
}
