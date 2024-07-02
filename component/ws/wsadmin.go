package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var adminClients = make(map[*websocket.Conn]bool)
var adminBroadcast = make(chan Message)

func handleAdminConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	adminClients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(adminClients, ws)
			break
		}
		adminBroadcast <- msg
	}
}

func handleAdminMessages() {
	for {
		msg := <-adminBroadcast
		for client := range adminClients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(adminClients, client)
			}
		}
	}
}
