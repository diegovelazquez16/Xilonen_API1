package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Estructura para manejar WebSocket
type WebSocketServer struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan []byte
	Upgrader  websocket.Upgrader
}

// Inicializar WebSocket
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		Clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan []byte),
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

// Manejar nuevas conexiones WebSocket
func (ws *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("❌ Error al conectar WebSocket:", err)
		return
	}
	defer conn.Close()

	ws.Clients[conn] = true

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			delete(ws.Clients, conn)
			break
		}
	}
}

// Manejar mensajes entrantes y enviarlos a los clientes
func (ws *WebSocketServer) HandleMessages() {
	for {
		msg := <-ws.Broadcast
		for client := range ws.Clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("❌ Error enviando mensaje:", err)
				client.Close()
				delete(ws.Clients, client)
			}
		}
	}
}
