// Interface a the Pao server
package pao

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/gorilla/websocket"
	"github.com/perlmonger42/greedy-bot/bot"
	"github.com/perlmonger42/greedy-bot/command"
	"github.com/perlmonger42/greedy-bot/game"
	"github.com/perlmonger42/greedy-bot/move"
)

type Service struct {
	conn     *websocket.Conn
	bot      Bot
	botColor string
}

type Bot interface {
	Name() string
	ChooseMove(*game.State) move.T
}

func NewService() *Service {
	return &Service{bot: bot.NewGreedyBot()}
}

func (svc *Service) Run(conn *websocket.Conn) {
	svc.conn = conn
	svc.PlayGame()
}

func (svc *Service) PlayGame() {
	defer func() {
		msg := fmt.Sprintf("%v - Terminating '%s' bot",
			time.Now(), svc.bot.Name())
		if r := recover(); r != nil {
			fmt.Println(msg+" because of %s", r)
			debug.PrintStack()
		}
		fmt.Println(msg)
		svc.closeConnection()
	}()

	for {
		action, text := svc.GetPaoCommand()
		fmt.Printf("%v - Command: %v\n", time.Now(), action)
		switch action {
		case "gameover":
			return
		case "board":
			if ok := svc.RunBoardCommand(text); !ok {
				return
			}
		case "color":
			fmt.Printf("Bot color is now %s\n", text)
			svc.RunColorCommand(text)
		default:
			fmt.Printf("%v - Ignoring: %v\n", time.Now(), string(text))
		}
	}
}

func (svc *Service) closeConnection() {
	for {
		fmt.Printf("%v - Closing conn\n", time.Now())
		if _, _, err := svc.conn.NextReader(); err != nil {
			svc.conn.Close()
			fmt.Printf("%v - Closed conn\n", time.Now())
			break
		}
	}
}

func (svc *Service) GetPaoCommand() (action string, command_text []byte) {
	var paoCommand command.Command

	if _, bytes, err := svc.conn.ReadMessage(); err != nil {
		panic(fmt.Sprintf("websocket read error (%v)", err.Error()))
	} else if err = json.Unmarshal(bytes, &paoCommand); err != nil {
		panic(fmt.Sprintf("command decode error: %v (input: %v)", err, bytes))
	} else {
		return paoCommand.Action, bytes
	}
}

func (svc *Service) RunBoardCommand(text []byte) bool {
	var bc command.BoardCommand
	if err := json.Unmarshal(text, &bc); err != nil {
		panic(fmt.Sprintf("board decode error: %v (input: %v)", err, text))
	}
	state := game.NewState(svc.botColor, bc.Board, bc.Dead)
	mv := svc.bot.ChooseMove(&state)
	fmt.Printf("Sending move: %s\n", mv.String())
	svc.SendCommand(mv.Command())
	return mv.Action() != move.Quit
}

func (svc *Service) RunColorCommand(text []byte) {
	var bc command.ColorCommand
	if err := json.Unmarshal(text, &bc); err != nil {
		panic(fmt.Sprintf("color decode error: %v (input: %v)", err, text))
	}
	svc.botColor = bc.Color
}

func (svc *Service) SendCommand(c command.Command) {
	if err := svc.conn.WriteJSON(c); err != nil {
		panic(fmt.Sprintf("websocket write error: %s (output: %v)", err, c))
	}
}
