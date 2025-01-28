package games

import (
	"sync"

	"github.com/gorilla/websocket"

	"fiesta_box/internal/models/responses"
)

type GameStatus string

const (
	Started GameStatus = "started"
	NotStarted GameStatus = "not_started"
	Completed GameStatus = "completed"
)

type GameClient struct {
	Room string `json:"room"`
	Client *websocket.Conn `json:"client"`
	UserID string `json:"userID"`
	Connected bool `json:"connected"`
}

type Game struct {
	Clients map[*websocket.Conn]*GameClient `json:"clients"`
	Broadcast chan responses.SocketResponse `json:"broadcast"`
	Status GameStatus `json:"started"`
	Mutex sync.Mutex `json:"mutex"`
	Room string `json:"room"`
}

type GameState struct {
	Clients int `json:"clients"`
	Status GameStatus `json:"status"`
	Room string `json:"room"`
}