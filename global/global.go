package global

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net/http"
	"sync"
)

var (
	Db *gorm.DB
	UP = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		HandshakeTimeout: 0,
		WriteBufferPool:  nil,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
	}
	Redis *redis.Client
	mu    *sync.Mutex
)
