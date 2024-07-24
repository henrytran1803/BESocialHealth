package ws

import (
	"BESocialHealth/Internal/account/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type WebSocketManager struct {
	clients      map[string]*websocket.Conn
	lastActivity map[string]time.Time
	activeUsers  map[string]bool
	mu           sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:      make(map[string]*websocket.Conn),
		lastActivity: make(map[string]time.Time),
		activeUsers:  make(map[string]bool),
	}
}

func (wm *WebSocketManager) AddClient(userID string, conn *websocket.Conn) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	wm.clients[userID] = conn
	wm.lastActivity[userID] = time.Now()
	wm.activeUsers[userID] = true
}

func (wm *WebSocketManager) RemoveClient(userID string) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	delete(wm.clients, userID)
	delete(wm.lastActivity, userID)
	delete(wm.activeUsers, userID)
}

func (wm *WebSocketManager) BroadcastToAll(message string) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	for _, conn := range wm.clients {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Error broadcasting message:", err)
		}
	}
}

func (wm *WebSocketManager) SendToUser(userID, message string) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	conn, ok := wm.clients[userID]
	if ok {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Error sending message to user:", err)
		}
	}
}

func (wm *WebSocketManager) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // Allow all origins (for testing)
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	defer conn.Close()
	userID := r.URL.Query().Get("userID")

	wm.AddClient(userID, conn)
	defer wm.RemoveClient(userID)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client connection error:", err)
			return
		}
		wm.mu.Lock()
		wm.lastActivity[userID] = time.Now()
		wm.activeUsers[userID] = true
		wm.mu.Unlock()

		var message struct {
			TargetUserID string `json:"target_user_id"`
			Message      string `json:"message"`
		}
		err = json.Unmarshal(msg, &message)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
			continue
		}
		if message.TargetUserID != "" {
			wm.SendToUser(message.TargetUserID, message.Message)
		}
	}
}

func (wm *WebSocketManager) CheckUserActivity() {
	for {
		time.Sleep(30 * time.Second) // Check every 30 seconds
		wm.mu.Lock()
		for userID, lastActive := range wm.lastActivity {
			if time.Since(lastActive) > 1*time.Minute { // If inactive for more than 1 minute
				conn := wm.clients[userID]
				err := conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					fmt.Println("Error sending close message:", err)
				}
				conn.Close()
				delete(wm.clients, userID)
				delete(wm.lastActivity, userID)
				delete(wm.activeUsers, userID)
				fmt.Printf("User %s disconnected due to inactivity\n", userID)
			}
		}
		wm.mu.Unlock()
	}
}

func (wm *WebSocketManager) GetActiveUsers() []accountmodels.UserActive {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	var activeUsers []accountmodels.UserActive
	for userID := range wm.activeUsers {
		activeUsers = append(activeUsers, accountmodels.UserActive{Id_user: userID, LastLogin: wm.lastActivity[userID]})
	}
	return activeUsers
}
