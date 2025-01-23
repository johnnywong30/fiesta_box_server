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
}

type Game struct {
	Clients map[*GameClient]bool `json:"clients"`
	Broadcast chan responses.SocketResponse `json:"broadcast"`
	Started GameStatus `json:"started"`
}