package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

func main()  {
	origin := "http://localhost/"
	url := "ws://localhost:9001/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!")); err != nil {
		log.Fatal(err)
	}
	var recv = make([]byte, 512)
	var n int
	if n, err = ws.Read(recv); err != nil {
		log.Fatal(err)
	}
	msg := fmt.Sprintf("%s", recv[:n])
	fmt.Printf("Received: %s.\n", msg)
	if msg == "a" {
		if _, err := ws.Write([]byte("you have pressed 'a.' ")); err != nil {log.Fatal(err)}
	}
}