package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/axrav/Systopher/backend/db"
	"github.com/axrav/Systopher/backend/errors"
	"github.com/axrav/Systopher/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func ServerStats(serverChannel chan []models.Server, dataChannel chan []models.ServerData, c *websocket.Conn, ctx context.Context) {
	var client http.Client
	servers := <-serverChannel
	for {
		select {
		case <-ctx.Done():
			return
		default:
			var stats []models.ServerData
			var data models.ServerData
			for _, server_data := range servers {

				server := "http://" + server_data.Ip + ":" + server_data.Port
				key, err := db.RedisClient.Get(db.Ctx, server).Result()
				if err != nil {
					fmt.Println(err)
					c.WriteJSON(fiber.Map{"server": server, "error": errors.NotFound.Error()})
				}
				req, err := http.NewRequest("GET", server+"/stats", nil)
				if err != nil {
					fmt.Println(err)
					c.WriteJSON(fiber.Map{"server": server, "error": "GET REQUEST"})
				}
				req.Header.Set("X-API-KEY", key)
				resp, err := client.Do(req)
				if err != nil {
					fmt.Println(err)
					c.WriteJSON(fiber.Map{"server": server, "error": errors.NoResponse.Error()})
				} else {
					body, err := io.ReadAll(resp.Body)
					if err != nil {
						fmt.Println(err)
						c.WriteJSON(fiber.Map{"server": server, "error": "read"})
					}
					err = json.Unmarshal(body, &data)
					data.Ip = server
					if err != nil {
						fmt.Println(err)
					}
					if data.Ping == "" {
						c.WriteJSON(fiber.Map{"server": server, "error": "TOKEN MISMATCH"})
					} else {
						stats = append(stats, data)
						data = models.ServerData{}
					}
				}

			}
			dataChannel <- stats
			stats = nil
		}
		time.Sleep(15 * time.Second)
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
	fmt.Println(token)
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
	} else if resp.StatusCode == 401 {
		return false, fmt.Errorf("the token is invalid")
	} else {
		return false, errors.NoResponse.Error()
	}
	// return false, fmt.Errorf("error sending request")
}
