package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"Xilonen-1/sensor/domain/models"
)

// Configuración de WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketServer struct {
	Clients map[*websocket.Conn]bool
}

// Crear un nuevo servidor WebSocket
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		Clients: make(map[*websocket.Conn]bool),
	}
}

// Manejar nuevas conexiones WebSocket
func (ws *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("❌ Error al actualizar la conexión a WebSocket:", err)
		return
	}
	defer conn.Close()

	ws.Clients[conn] = true
	log.Println("✅ Nueva conexión WebSocket establecida.")

	// Mantener la conexión abierta hasta que el cliente se desconecte
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("🔌 Cliente WebSocket desconectado.")
			delete(ws.Clients, conn)
			conn.Close()
			break
		}
	}
}


// Enviar datos del sensor a todos los clientes conectados
func (ws *WebSocketServer) SendSensorData(sensor models.SensorMQ135) {
	data, err := json.Marshal(sensor)
	if err != nil {
		log.Println("❌ Error al serializar datos del sensor:", err)
		return
	}

	log.Println("📡 Enviando datos al WebSocket:", string(data)) 

	for client := range ws.Clients {
		if err := client.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Println("⚠️ Error al enviar mensaje a cliente WebSocket:", err)
			client.Close()
			delete(ws.Clients, client)
		}
	}
}

