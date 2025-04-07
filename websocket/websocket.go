package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocketServer maneja todas las conexiones WebSocket
type WebSocketServer struct {
	clients  map[*websocket.Conn]bool
	mu       sync.Mutex
	upgrader websocket.Upgrader
}

// NewWebSocketServer crea una nueva instancia del servidor WebSocket
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients: make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Permitir conexiones desde cualquier origen
			},
		},
	}
}

// HandleConnection maneja nuevas conexiones WebSocket
func (ws *WebSocketServer) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("❌ Error al actualizar la conexión a WebSocket:", err)
		return
	}

	ws.mu.Lock()
	ws.clients[conn] = true
	ws.mu.Unlock()
	log.Println("✅ Nueva conexión WebSocket establecida.")

	// Manejar cierre de conexión
	go func() {
		defer func() {
			ws.mu.Lock()
			delete(ws.clients, conn)
			ws.mu.Unlock()
			conn.Close()
			log.Println("⚠️ Conexión WebSocket cerrada.")
		}()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}()
}

// BroadcastMessage envía datos a todos los clientes conectados
func (ws *WebSocketServer) BroadcastMessage(sensorType string, data interface{}) {
	// Convertir el mensaje a un mapa genérico
	msgMap, ok := data.(map[string]interface{})
	if !ok {
		log.Println("❌ Error: los datos enviados no son un mapa válido")
		return
	}

	// Agregar el tipo de sensor al mensaje
	msgMap["tipo"] = sensorType

	// Serializar a JSON
	message, err := json.Marshal(msgMap)
	if err != nil {
		log.Println("❌ Error al serializar mensaje:", err)
		return
	}

	// Enviar mensaje a cada cliente conectado
	ws.mu.Lock()
	defer ws.mu.Unlock()
	for client := range ws.clients {
		if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("❌ Error al enviar mensaje a cliente WebSocket:", err)
			client.Close()
			delete(ws.clients, client)
		}
	}
}
