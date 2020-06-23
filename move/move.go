package move

import (
	"fmt"
	"strconv"

	"github.com/perlmonger42/greedy-bot/command"
	"github.com/perlmonger42/greedy-bot/game"
)

// Prerequisites for building this file:
// - after getting sources:
//     `go get -u golang.org/x/tools/cmd/stringer`
// - Before compiling code that uses Action.String():
//     `(cd move &&go generate)`
//go:generate go run golang.org/x/tools/cmd/stringer -type=Action
type Action int

const (
	Quit Action = iota
	Flip
	Take
	Move
)

type Location struct {
	row, col int
}

func (loc Location) String() string {
	return string(rune(loc.col+'A')) + strconv.Itoa(loc.row+1)
}

type T struct {
	action Action
	at     Location   // not used for Quit
	to     Location   // only used for Move and Take
	actor  game.Piece // only used for Move and Take
	killed game.Piece // only used for Take
}

func (move *T) String() string {
	switch move.action {
	case Quit:
		return "resign"
	case Flip:
		return fmt.Sprintf("flip %s", move.at)
	case Take:
		return fmt.Sprintf("%s at %s takes %s at %s",
			move.actor, move.at.String(), move.killed.String(), move.to.String())
	case Move:
		return fmt.Sprintf("%s at %s moves to %s",
			move.actor.String(), move.at.String(), move.to.String())
	}
	data := fmt.Sprintf("{action: %d, at: %s, to: %s, actor: %s, killed: %s}",
		move.action, move.at, move.to, move.actor.String(), move.killed.String())
	panic(fmt.Sprintf("unrecognized move: %s", data))
}

func NewQuit() T {
	return T{action: Quit}
}

func NewFlip(r, c int) T {
	return T{action: Flip, at: Location{r, c}}
}

func NewTake(actor game.Piece, r, c int, killed game.Piece, r2, c2 int) T {
	return T{
		action: Take,
		at:     Location{r, c},
		to:     Location{r2, c2},
		actor:  actor,
		killed: killed,
	}
}

func NewMove(actor game.Piece, r, c, r2, c2 int) T {
	return T{action: Move, at: Location{r, c}, to: Location{r2, c2}, actor: actor}
}

func (m T) Command() command.Command {
	switch m.action {
	case Flip:
		return command.Command{
			Action:   "move",
			Argument: "?" + m.at.String(),
		}
	case Take:
		fallthrough
	case Move:
		return command.Command{
			Action:   "move",
			Argument: m.at.String() + ">" + m.to.String(),
		}
	default:
		fallthrough
	case Quit:
		return command.Command{
			Action: "resign",
		}
	}
}

func (m T) Action() Action {
	return m.action
}

func (m T) Killed() game.Piece {
	return m.killed
}

type moveFinder struct {
	team  *game.Team // "us"
	them  *game.Team // the other team
	board game.Board
	flips []T
	moves []T
}

func LegalMoves(team, them *game.Team, board game.Board) []T {
	searcher := moveFinder{
		team:  team,
		them:  them,
		board: board,
		flips: []T{},
		moves: []T{},
	}
	return searcher.verifyMoves(searcher.findMoves())
}

func (f *moveFinder) findMoves() []T {
	for r, row := range f.board {
		for c, piece := range row {
			if piece == game.FaceDown {
				f.flips = append(f.flips, NewFlip(r, c))
			} else if f.team.Contains(piece) {
				f.tryUsing(r, c, piece)
			}
		}
	}
	return append(f.flips, f.moves...)
}

func (f *moveFinder) tryUsing(r, c int, piece game.Piece) {
	if r > 0 {
		f.tryDirection(piece, r, c, -1, 0)
	}
	if c > 0 {
		f.tryDirection(piece, r, c, 0, -1)
	}
	if r < 4-1 {
		f.tryDirection(piece, r, c, +1, 0)
	}
	if c < 8-1 {
		f.tryDirection(piece, r, c, 0, +1)
	}
}

func (f *moveFinder) tryDirection(myPiece game.Piece, r, c, dr, dc int) {
	r2, c2 := r+dr, c+dc
	defender := f.board.At(r2, c2)

	// move
	if defender == game.None {
		f.moves = append(f.moves, NewMove(myPiece, r, c, r2, c2))
	}

	// take neighbor
	if myPiece.CanTakeIfAdjacent(defender) {
		f.moves = append(f.moves, NewTake(myPiece, r, c, defender, r2, c2))
	}

	// cannon jump
	if myPiece == f.team.Q {
		r2, c2, defender = f.cannonTarget(r, c, dr, dc)
		if f == nil {
			panic("how can f be nil?")
		}
		if f.them == nil {
			panic("how can f.them be nil?")
		}
		if defender.In(f.them.Set) {
			f.moves = append(f.moves, NewTake(myPiece, r, c, defender, r2, c2))
		}
	}
}

// cannonTarget returns the coordinates and Piece of a valid target of
// a cannon at r,c in the direction dr,dc; if there is no valid target,
// returns None and the coordinates of the last square in that direction.
func (f *moveFinder) cannonTarget(r, c, dr, dc int) (r2, c2 int, p game.Piece) {
	save_r, save_c := r, c
	pieceCount := 0

	// move r,c to the second piece in the direction dr,dc
	if dr == 1 {
		if r < 2 {
			for r += 1; r < 4; r += 1 {
				if f.board.At(r, c) != game.None {
					pieceCount += 1
					if pieceCount == 2 {
						break
					}
				}
			}
		}
	} else if dr == -1 {
		if r > 1 {
			for r -= 1; r >= 0; r -= 1 {
				if f.board.At(r, c) != game.None {
					pieceCount += 1
					if pieceCount == 2 {
						break
					}
				}
			}
		}
	} else if dc == 1 {
		if c < 6 {
			for c += 1; c < 8; c += 1 {
				if f.board.At(r, c) != game.None {
					pieceCount += 1
					if pieceCount == 2 {
						break
					}
				}
			}
		}
	} else /* dc == -1 */ {
		if c > 1 {
			for c -= 1; c >= 0; c -= 1 {
				if f.board.At(r, c) != game.None {
					pieceCount += 1
					if pieceCount == 2 {
						break
					}
				}
			}
		}
	}

	if pieceCount == 2 {
		return r, c, f.board.At(r, c)
	}
	return save_r, save_c, game.None
}

func (f *moveFinder) verifyMoves(moves []T) []T {
	for _, m := range moves {
		p := f.board.At(m.at.row, m.at.col)
		switch m.action {
		case Quit:

		case Flip:
			if p != game.FaceDown {
				panic(fmt.Sprintf("generated turnup of non-facedown piece: %s", m.String()))
			}
		case Move:
			if !f.team.Contains(p) {
				panic(fmt.Sprintf("moving a piece that isn't mine: %s", m.String()))
			}
		case Take:
			if !f.team.Contains(p) {
				panic(fmt.Sprintf("taking with a piece that isn't mine: %s", m.String()))
			}
			k := f.board.At(m.to.row, m.to.col)
			if !f.them.Contains(k) {
				panic(fmt.Sprintf("taking a piece that isn't the enemy: %s", m.String()))
			}
		}
	}
	return moves
}
