package game

/* The Team type allows code to refer to a piece of a certain rank for a given
 team without having to select between two different Pieces (e.g., RedCannon
 vs. BlackCannon).  For example, instead of writing:

    func do_something_with_king(board Board, color string) {
			if color == "Red" {
				do_something_with_this_piece(board, RedKing)
			} else if color == "Black" {
				do_something_with_this_piece(board, BlackKing)
			} else {
				panic "unknown color"
			}
		}
	}

you can write code like:

    func do_something_with_king(board Board, Team *t) {
		do_something_with_this_piece(board, t.King)
	}

The `QPHCEGK` field allows code to iterate over the pieces of a given color in
order of ascending rank, which is the order indicated by the symbols in the
field name.  The field _could_ have been called `List`, but this name serves as
a reminder of the piece ordering wherever the field is used.

"Ascending" is a bit of a misnomer, since there's no completely-consistent
place to put the cannons. So they're arbitrarily listed first.
*/
type Team struct {
	Q          Piece
	P          Piece
	H, C, E, G Piece
	K          Piece
	QPHCEGK    [7]Piece
	Set        SetOfPieces
}

var RedTeam Team = Team{
	Q: RedCannon,
	P: RedPawn,
	H: RedHorse,
	C: RedCart,
	E: RedElephant,
	G: RedGuard,
	K: RedKing,
	QPHCEGK: [7]Piece{
		RedCannon,
		RedPawn,
		RedHorse, RedCart, RedElephant, RedGuard,
		RedKing,
	},
	Set: SetOfPieces(RedCannon |
		RedPawn |
		RedHorse | RedCart | RedElephant | RedGuard |
		RedKing),
}

var BlackTeam Team = Team{
	Q: BlackCannon,
	P: BlackPawn,
	H: BlackHorse,
	C: BlackCart,
	E: BlackElephant,
	G: BlackGuard,
	K: BlackKing,
	QPHCEGK: [7]Piece{
		BlackCannon,
		BlackPawn,
		BlackHorse, BlackCart, BlackElephant, BlackGuard,
		BlackKing,
	},
	Set: SetOfPieces(BlackCannon |
		BlackPawn |
		BlackHorse | BlackCart | BlackElephant | BlackGuard |
		BlackKing),
}

var Teams [2]*Team = [2]*Team{&RedTeam, &BlackTeam}

func (team *Team) Contains(p Piece) bool {
	return team.Set.Contains(p)
}
