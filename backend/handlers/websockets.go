package handlers

import (
	"context"
	"fmt"

	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/models"
	"github.com/gofiber/websocket/v2"
)

func ServerWS(c *websocket.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
	servers := c.Locals("servers").([]db.Server)
	serverChannel := make(chan []db.Server, 1)
	dataChannel := make(chan []models.ServerData)
	serverChannel <- servers
	go helpers.ServerStats(serverChannel, dataChannel, c, ctx)
	for {
		data := <-dataChannel
		err := c.WriteJSON(data)
		if err != nil {
			fmt.Println(err)
			cancel()
		}

	}
}
