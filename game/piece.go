package game

import "fmt"

// Prerequisites for building this file:
// - after getting sources:
//     `go get -u golang.org/x/tools/cmd/stringer`
// - Before compiling code that uses Piece.String():
//     `(cd game && go generate)`
//go:generate go run golang.org/x/tools/cmd/stringer -type=Piece
func init() {
	var _ string = RedCannon.String()
}

type bitset uint16      // a compact representation of "SET OF [0..15]"
type member bitset      // a singleton set (a bitset with only one bit set)
type Piece member       // a bitset with only one bit set
type SetOfPieces bitset // a bitset of Piece

const (
	None     Piece = 0
	FaceDown Piece = 1 << iota

	RedCannon
	RedPawn
	RedHorse
	RedCart
	RedElephant
	RedGuard
	RedKing

	BlackCannon
	BlackPawn
	BlackHorse
	BlackCart
	BlackElephant
	BlackGuard
	BlackKing
)

func NewPiece(paoDescriptor string) Piece {
	if piece, ok := paoStringToPiece[paoDescriptor]; ok {
		return piece
	} else {
		panic(fmt.Sprintf("unknown piece descriptor: %q", paoDescriptor))
	}
}

var paoStringToPiece map[string]Piece = map[string]Piece{
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
	"?": FaceDown,
	".": None,
	" ": None,
}

func (p Piece) AsSingletonSet() SetOfPieces {
	return SetOfPieces(p)
}

func (p Piece) In(set SetOfPieces) bool {
	return (p.AsSingletonSet() & set) != 0
}

func (a Piece) CanTakeIfAdjacent(d Piece) bool {
	vulnerable := canTakeIfAdjacent[a]
	return (vulnerable & d.AsSingletonSet()) != 0
}

func NewSetOfPieces(pieces ...Piece) (u SetOfPieces) {
	for _, p := range pieces {
		u = u | p.AsSingletonSet()
	}
	return
}

func (set SetOfPieces) Contains(p Piece) bool {
	return (p.AsSingletonSet() & set) != 0
}

var canTakeIfAdjacent map[Piece]SetOfPieces

func init() {
	canTakeIfAdjacent = make(map[Piece]SetOfPieces)
	canTakeIfAdjacent[FaceDown] = NewSetOfPieces()
	for i := 0; i < 2; i++ {
		a, d := Teams[i], Teams[1-i] // team a: "attacker"; team d: "defender"
		canTakeIfAdjacent[a.K] = NewSetOfPieces(d.Q, d.H, d.C, d.E, d.G, d.K)
		canTakeIfAdjacent[a.G] = NewSetOfPieces(d.Q, d.P, d.H, d.C, d.E, d.G)
		canTakeIfAdjacent[a.E] = NewSetOfPieces(d.Q, d.P, d.H, d.C, d.E)
		canTakeIfAdjacent[a.C] = NewSetOfPieces(d.Q, d.P, d.H, d.C)
		canTakeIfAdjacent[a.H] = NewSetOfPieces(d.Q, d.P, d.H)
		canTakeIfAdjacent[a.P] = NewSetOfPieces(d.P, d.K)
		canTakeIfAdjacent[a.Q] = NewSetOfPieces()
	}
}
