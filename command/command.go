package command

// Command is a generic command between a pao server and client
type Command struct {
	Action, Argument string
}

// // PlayerCommand is a command from a player to the server
// type PlayerCommand struct {
// 	C Command
// 	P *player.Player
// }

// ChatCommand is a chat message broadcast from server to client
type ChatCommand struct {
	Action, Player, Color, Message string
	Auth                           bool
}

// BoardCommand is an update of the board state from server to client
type BoardCommand struct {
	Action     string
	Board      [][]string
	Dead       []string
	LastMove   []string
	LastDead   string
	YourTurn   bool
	WhoseTurn  string
	TurnColor  string
	NumPlayers int
}

// ColorCommand is a command indicating to a client which color is theirs
// upon a flip
type ColorCommand struct {
	Action, Color string
}

// GameOverCommand is a command from the server to clients indicating that
// the game has concluded and who has won
type GameOverCommand struct {
	Action, Message string
	YouWin          bool
}
