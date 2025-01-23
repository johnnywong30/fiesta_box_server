package services

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"fiesta_box/internal/models/games"
	"fiesta_box/internal/models/responses"
)


func NewGame(c *websocket.Conn, room string) games.Game {
	client := games.GameClient{
		Room: room,
		Client: c,
		UserID: uuid.NewString(),
	}

	clients := map[*games.GameClient]bool{
		&client: true,
	}

	return games.Game{
		Clients: clients,
		Broadcast: make(chan responses.SocketResponse),
		Started: games.NotStarted,
	}
}