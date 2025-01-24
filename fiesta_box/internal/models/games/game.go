package games

import (
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
}