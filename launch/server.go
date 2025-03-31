package launch

import (
	"fmt"
	"log"
	"net/http"

	"Xilonen-1/core"
	"Xilonen-1/sensor/infraestructure/websocket"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	core.LoadConfig()
	router := gin.Default()

	wsServer := websocket.NewWebSocketServer()
	go wsServer.HandleMessages()

	router.GET("/ws", func(c *gin.Context) {
		wsServer.HandleConnections(c.Writer, c.Request)
	})

	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
