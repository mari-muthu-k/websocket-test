package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func main() {
	r := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"} // Specify allowed origins
	// config.AllowMethods = []string{"*"} // Specify allowed HTTP methods
	// config.AllowHeaders = []string{"*"}
	// config.AllowWebSockets = true

	// r.Use(cors.New(config))
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("Client connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		log.Printf("Message from client: %s\n", msg)
		server.BroadcastToRoom("/", "chat", "message", msg)
	})

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	r.StaticFS("/public", http.Dir("./public"))

	if err := r.Run(":5000"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
