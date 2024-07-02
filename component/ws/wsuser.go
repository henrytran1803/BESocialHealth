package ws

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var userClients = make(map[*websocket.Conn]bool)
var userBroadcast = make(chan Message)

type Message struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Content string `json:"content"`
}

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

//func SendMessageToClient(userID int, message messagemodels.Message) error {
//	// Lấy danh sách kết nối websocket của người dùng từ userClients (nếu sử dụng cấu trúc map trong phần trước)
//	if conns, ok := userClients[userID]; ok {
//		for _, conn := range conns {
//			err := conn.WriteJSON(message)
//			if err != nil {
//				return err
//			}
//		}
//	}
//	return nil
//}
