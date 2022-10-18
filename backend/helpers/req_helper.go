package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/axrav/SysAnalytics/backend/db"
	"github.com/axrav/SysAnalytics/backend/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ServerStats(serverChannel chan []string, dataChannel chan types.ServerData, c *websocket.Conn) {
	servers := <-serverChannel
	var data types.ServerData
	var client http.Client
	for {
		for _, server := range servers {
			key, err := db.RedisClient.Get(db.Ctx, server).Result()
			if err != nil {
				fmt.Println(err)
				c.WriteMessage(websocket.TextMessage, []byte("Error: Token not found"))
				c.Close()
			}
			req, err := http.NewRequest("GET", server, nil)
			if err != nil {
				fmt.Println(err)
				c.WriteJSON(fiber.Map{"error": "Error : There was an error in communicating with the server(GET REQUEST)", "server": server, "errorType": "GET REQUEST"})
			}
			req.Header.Set("X-API-KEY", key)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				c.WriteJSON(fiber.Map{"error": "Error : There was an error in communicating with the server(Perform REQUEST)", "server": server, "errorType": "NoResponse"})
			} else {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					c.WriteJSON(fiber.Map{"error": "Error : There was an error in communicating with the server(READ REQUEST)", "server": server, "errorType": "read"})
				}
				err = json.Unmarshal(body, &data)
				data.Ip = server
				if err != nil {
					fmt.Println(err)
				}
				if data.Ping == "" {
					c.WriteJSON(fiber.Map{"error": "Error : There was an error in communicating with the server('TOKEN MISMATCH')", "server": server, "errorType": "TOKEN MISMATCH"})
				}
			}

			dataChannel <- data
		}
		time.Sleep(30 * time.Second)

	}
}
