package service

import (
	"WeiYangWork/global"
	"WeiYangWork/utils"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"sync"
)

var clients sync.Map

type Client struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

// BroadcastMessage 广播消息给所有连���的客户端
func BroadcastMessage(msg string) {
	clients.Range(func(key, value interface{}) bool {
		client := key.(*Client)
		client.mu.Lock()
		err := client.conn.WriteMessage(websocket.TextMessage, []byte(msg))
		client.mu.Unlock()
		if err != nil {
			log.Printf("Error sending message: %v", err)
			client.conn.Close()
			clients.Delete(client)
		}
		return true
	})
}

func WsHandler(conn *websocket.Conn, teamID uint) {
	client := &Client{conn: conn}
	clients.Store(client, true)
	defer func() {
		clients.Delete(client)
		conn.Close()
	}()
	go listenToTeamMessages(teamID)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		// 将消息缓存在 Redis 中
		if err = utils.CacheMessage(teamID, string(message)); err != nil {
			log.Printf("Error caching message: %v", err)
		}
		// 将消息发布到队伍频道
		if err = utils.PublishMessage(teamID, string(message)); err != nil {
			log.Printf("Error publishing message: %v", err)
		}
	}
	log.Println("WebSocket connection closed")
}

// listenToTeamMessages 监听队伍消息，并推送到 WebSocket 客户端
func listenToTeamMessages(teamID uint) {
	pubsub := global.Redis.Subscribe("team:" + strconv.Itoa(int(teamID)))
	defer pubsub.Close()

	ch := pubsub.Channel()
	for msg := range ch {
		clients.Range(func(key, value interface{}) bool {
			client := key.(*Client)
			client.mu.Lock()
			err := client.conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			client.mu.Unlock()
			if err != nil {
				log.Printf("Error sending message: %v", err)
				client.conn.Close()
				clients.Delete(client)
			}
			return true
		})
	}
}
