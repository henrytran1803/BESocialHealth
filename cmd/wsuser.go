package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var userClients = make(map[*websocket.Conn]bool)
var userBroadcast = make(chan Message)

func handleUserConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	userClients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(userClients, ws)
			break
		}
		userBroadcast <- msg
	}
}

func handleUserMessages() {
	for {
		msg := <-userBroadcast
		for client := range userClients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(userClients, client)
			}
		}
	}
}
