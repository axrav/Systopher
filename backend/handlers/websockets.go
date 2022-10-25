package handlers

import (
	"context"
	"fmt"

	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/websocket/v2"
)

func ServerWS(c *websocket.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
	servers := c.Locals("servers").([]types.Server)
	serverChannel := make(chan []types.Server, 1)
	dataChannel := make(chan types.ServerData)
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
