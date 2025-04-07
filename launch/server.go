package launch

import (
	"fmt"
	"log"

	"Xilonen-1/core"
	"Xilonen-1/websocket"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	core.LoadConfig()
	router := gin.Default()

	wsServer := websocket.NewWebSocketServer()

	router.GET("/ws", func(c *gin.Context) {
		wsServer.HandleConnection(c.Writer, c.Request)
	})


	fmt.Println("🚀 Servidor corriendo en http://localhost:8080")
	log.Fatal(router.Run(":8080"))
}
