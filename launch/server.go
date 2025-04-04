package launch

import (
	"fmt"
	"log"

	"Xilonen-1/core"
	"Xilonen-1/sensor/infraestructure/websocket"
	"Xilonen-1/simulator" // Importamos el simulador

	"github.com/gin-gonic/gin"
)

// StartServer inicia el servidor con WebSocket y rutas HTTP
func StartServer() {
	core.LoadConfig()
	router := gin.Default()

	// Crear WebSocket Server
	wsServer := websocket.NewWebSocketServer()

	// Ruta WebSocket
	router.GET("/ws", func(c *gin.Context) {
		wsServer.HandleConnection(c.Writer, c.Request)
	})

	// Iniciar simulador de datos
	go simulator.SimulateSensorData(wsServer)

	// Iniciar servidor
	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
