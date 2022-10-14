package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/axrav/SysAnalytics/backend/types"
	"github.com/gofiber/websocket/v2"
)

func ServerWS(c *websocket.Conn) {
	servers := c.Locals("servers").([]string)
	fmt.Println("Servers: ", servers)
	serverChannel := make(chan []string, 1)
	dataChannel := make(chan types.ServerData)
	serverChannel <- servers
	go func() {
		servers := <-serverChannel
		var data types.ServerData
		for {
			for _, server := range servers {
				resp, err := http.Get(server)
				if err != nil {
					fmt.Println(err)
					c.WriteMessage(websocket.TextMessage, []byte("Error"))
					c.Close()
				}
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					c.WriteMessage(websocket.TextMessage, []byte("Error"))
					c.Close()
				}
				json.Unmarshal(body, &data)
				dataChannel <- data
			}
			time.Sleep(30 * time.Second)

		}
	}()
	for {
		data := <-dataChannel
		c.WriteJSON(data)
	}
}
