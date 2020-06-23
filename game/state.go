package game

import "strings"

// State represents the current state of a game of Ban Chi.
type State struct {
	Board Board         // state of the board's squares
	Dead  []Piece       // pieces that have been captured
	Down  map[Piece]int // pieces that have not been turned face-up yet
	Score int           // the heuristic score for this game state
	Us    *Team         // the team whose turn it is to play
	Them  *Team         // the other team
}

// NewState builds a game stage from arrays of strings representing
// the content of the board and boneyard. The strings are Pao-style
// content descriptors ("Q", "k", ".", "?", etc).
func NewState(toMove string, boardStrs [][]string, deadStrs []string) State {
	board := NewBoard(boardStrs)
	dead := deadList(deadStrs)
	down := downList(board, dead)
	score := computeScore(board, dead)
	var us, them *Team
	if toMove == "" { // first move of the game; teams not yet determined
		us, them = nil, nil
	} else if strings.ToLower(toMove[0:1]) == "b" {
		us, them = &BlackTeam, &RedTeam
	} else {
		us, them = &RedTeam, &BlackTeam
	}
	return State{Board: board, Dead: dead, Down: down, Score: score, Us: us, Them: them}
}

func deadList(deadAsStrings []string) []Piece {
	dead := []Piece{}
	for _, str := range deadAsStrings {
		dead = append(dead, NewPiece(str))
	}
	return dead
}

func downList(board Board, dead []Piece) map[Piece]int {
	// Assume all pieces are face down...
	pieceCounts := map[Piece]int{
		RedCannon:   2,
		RedPawn:     5,
		RedHorse:    2,
		RedCart:     2,
		RedElephant: 2,
		RedGuard:    2,
		RedKing:     1,

		BlackCannon:   2,
		BlackPawn:     5,
		BlackHorse:    2,
		BlackCart:     2,
		BlackElephant: 2,
		BlackGuard:    2,
		BlackKing:     1,
	}
	// ...but pieces face up on the board are NOT facedown...
	board.Each(func(p Piece) {
		pieceCounts[p] -= 1
	})
	// ...and dead pieces are also NOT facedown.
	for _, p := range dead {
		pieceCounts[p] -= 1
	}
	for k, v := range pieceCounts {
		if v < 1 { // deletes None and FaceDown as well as zeros
			delete(pieceCounts, k)
		}
	}
	return pieceCounts
}

// computeScore returns a rough estimation of the strength of a game state.
// Positive values mean Red seems to be winning; negative values favor Black.
func computeScore(board Board, dead []Piece) int {
	// This is stupid-simple valuation: sum up the values of the materiel.
	//
	// Each Piece has an associated point value, which is positive for red
	// pieces and negative for black pieces.
	// A face-up piece on the board adds twice its points to the score.
	// A face-down piece adds its points to the score.
	// A dead piece adds zero points to the score.

	// 1. Start off assuming all pieces are present and face-down. We'll
	// correct this assumption on a piece-by-piece basis in the later steps.
	// We get score = (sum of p's Points, for all p in the 32 initial pieces),
	// which is zero, since the red and black pieces balance each other out.
	var score int = 0

	// 2. We next adjust the score to account for face-up pieces. Each face-up
	// piece is worth twice its points, but we started off assuming it was
	// face-down, and so have already added its point value once. Now that we
	// know it's face-up, we only need to add its points once more.
	for _, row := range board {
		for _, piece := range row {
			score += PiecePoints[piece]
		}
	}

	// 3. We next adjust the score to account for dead pieces. Each dead piece
	// is worth zero points, but we originally assumed it was face-down and
	// added its points to the score. To correct this, we subtract its points
	// from the score.
	for _, t := range dead {
		score -= PiecePoints[t]
	}

	return score
}

var PiecePoints map[Piece]int = map[Piece]int{
	None:     0,
	FaceDown: 0,

	RedKing:     700,
	RedGuard:    6,
	RedCannon:   5,
	RedElephant: 4,
	RedCart:     3,
	RedHorse:    2,
	RedPawn:     1,

	BlackKing:     -700,
	BlackGuard:    -6,
	BlackCannon:   -5,
	BlackElephant: -4,
	BlackCart:     -3,
	BlackHorse:    -2,
	BlackPawn:     -1,
}
