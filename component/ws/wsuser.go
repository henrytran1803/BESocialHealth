package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

// WebSocketManager quản lý các kết nối WebSocket và các thông báo
type WebSocketManager struct {
	clients map[string]*websocket.Conn
	admin   *websocket.Conn
	mu      sync.Mutex
}

// NewWebSocketManager tạo một WebSocketManager mới
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients: make(map[string]*websocket.Conn),
	}
}

// AddClient thêm một client mới vào manager
func (wm *WebSocketManager) AddClient(userID string, conn *websocket.Conn) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	wm.clients[userID] = conn
}

// RemoveClient xóa một client khỏi manager
func (wm *WebSocketManager) RemoveClient(userID string) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	delete(wm.clients, userID)
}

// BroadcastToAll gửi thông báo đến tất cả các client
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

// SendToUser gửi thông báo đến một user cụ thể
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

// WebSocketHandler xử lý các kết nối WebSocket
func (wm *WebSocketManager) WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
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
		// Xử lý thông báo từ client đến user khác
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
