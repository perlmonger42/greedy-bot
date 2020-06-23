// A server that runs the "greedy" Pao bot.
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/perlmonger42/greedy-bot/pao"
)

type WebsocketService interface {
	Run(*websocket.Conn)
}

var PaoService WebsocketService = pao.NewService()

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	host, port := os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")
	if port == "" {
		port = "1960"
	}
	bind := fmt.Sprintf("%v:%v", host, port)
	fmt.Printf("Listening on %s\n", bind)
	http.HandleFunc("/", httpHandler)
	http.ListenAndServe(bind, nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v - Got a new customer!\n", time.Now())
	fmt.Printf("%v - Request: %v\n", time.Now(), r)

	if conn, err := upgrader.Upgrade(w, r, nil); err != nil {
		fmt.Printf("%v - Websocket build error: %v\n", time.Now(), err.Error())
	} else {
		fmt.Printf("%v - Spinning up a websocket goroutine\n", time.Now())
		go PaoService.Run(conn)
	}

	fmt.Printf("%v - Exiting HTTP Handler\n", time.Now())
}
