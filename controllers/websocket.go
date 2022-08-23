package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]string)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))

		broadcastToAll(messageType, p)
	}
}

func HandleWS(c *gin.Context) {
	username := c.Query("username")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Println(err)
	}

	log.Println(username, "Successfully Connected...")

	clients[ws] = username

	log.Println(clients)

	go reader(ws)
}

func broadcastToAll(messageType int, p []byte) {
	for client := range clients {
		if err := client.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
