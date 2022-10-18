package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
				c.WriteJSON(fiber.Map{"server": server, "errorType": "server not found"})
			}
			req, err := http.NewRequest("GET", server, nil)
			if err != nil {
				fmt.Println(err)
				c.WriteJSON(fiber.Map{"server": server, "errorType": "GET REQUEST"})
			}
			req.Header.Set("X-API-KEY", key)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				c.WriteJSON(fiber.Map{"server": server, "errorType": "NoResponse"})
			} else {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					c.WriteJSON(fiber.Map{"server": server, "errorType": "read"})
				}
				err = json.Unmarshal(body, &data)
				data.Ip = server
				if err != nil {
					fmt.Println(err)
				}
				if data.Ping == "" {
					c.WriteJSON(fiber.Map{"server": server, "errorType": "TOKEN MISMATCH"})
				}
			}

			dataChannel <- data
		}
		time.Sleep(30 * time.Second)

	}
}

func TestRequest(ip string, port string, token string) (bool, error) {
	client := &http.Client{
		Timeout: time.Second * 8,
	}
	req, err := http.NewRequest("GET", "http://"+ip+":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	req.Header.Set("X-API-KEY", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		if strings.HasSuffix(err.Error(), "(Client.Timeout exceeded while awaiting headers)") {
			return false, fmt.Errorf("timeout, is the Port exposed?")

		} else if strings.HasSuffix(err.Error(), "connection refused") {
			return false, fmt.Errorf("is the server running?")

		} else {
			return false, err
		}
	}
	if resp.StatusCode == 200 {
		return true, nil
	} else if resp.StatusCode == 403 {
		return false, fmt.Errorf("the token is invalid")
	} else {
		return false, fmt.Errorf("the server is not responding")
	}
	// return false, fmt.Errorf("error sending request")
}
