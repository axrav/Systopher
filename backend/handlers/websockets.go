package handlers

import (
	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/websocket/v2"
)

func ServerWS(c *websocket.Conn) {
	servers := c.Locals("servers").([]types.Server)
	serverChannel := make(chan []types.Server, 1)
	dataChannel := make(chan types.ServerData)
	serverChannel <- servers
	go helpers.ServerStats(serverChannel, dataChannel, c)
	for {
		data := <-dataChannel
		c.WriteJSON(data)
	}
}
