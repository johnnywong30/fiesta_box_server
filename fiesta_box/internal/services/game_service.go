package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"fiesta_box/internal/models/games"
	"fiesta_box/internal/models/responses"
)

type GameServiceInterface interface {
	NewGame(c *websocket.Conn, room string) games.Game
}

type GameService struct{
	games map[string]*games.Game
} 


func (s *GameService) NewGame(c *websocket.Conn, room string) (games.Game, error) {

	// check if room exists, fail if it does
	if _, ok := s.games[room]; ok {
		var g games.Game
		return g, fmt.Errorf("Game room %s already exists", room)
	}

	// create new room

	// create game client for this websocket connection
	client := games.GameClient{
		Room: room,
		Client: c,
		UserID: uuid.NewString(),
		Connected: true,
	}

	// add game client to the game room's client map
	clients := map[*websocket.Conn]*games.GameClient{
		c: &client,
	}

	// create game room
	game := games.Game{
		Clients: clients,
		Broadcast: make(chan responses.SocketResponse),
		Status: games.NotStarted,
	}
	// add game room to game service map
	s.games[room] = &game
	// return the created game room
	return game, nil
}

func (s *GameService) AddToGame(c *websocket.Conn, room string) (games.Game, error) {
	// check if room exists, fail if it doesn't
	game, ok := s.games[room]
	if !ok {
		var g games.Game
		return g, fmt.Errorf("game room %s does not exist - failed to join game", room)
	} 
	client := games.GameClient{
		Room: room,
		Client: c,
		UserID: uuid.NewString(),
		Connected: true,
		}
	game.Clients[c] = &client

	serverResponse := responses.SocketResponse{
		Status: responses.Success,
		Message: fmt.Sprintf("client %s joined game %s", client.UserID, room),
		}

	game.Broadcast <- serverResponse

	return *game, nil
	
}

func (s *GameService) RemoveFromGame(c *websocket.Conn, room string) (games.Game, error) {
	// check if room exists, fail if it doesn't
	game, ok := s.games[room]
	if !ok {
		var g games.Game
		return g, fmt.Errorf("game room %s does not exist - failed to leave game", room)
	}

	client, ok := game.Clients[c]
	if !ok {
		var g games.Game
		return g, fmt.Errorf("game client does not exist in room %s - failed to leave game", room)
	}

	clientID := client.UserID
	
	// remove client from game room's client map
	delete(game.Clients, c)

	serverResponse := responses.SocketResponse{
		Status: responses.Success,
		Message: fmt.Sprintf("client %s left game %s", clientID, room),
		}

	game.Broadcast <- serverResponse

	return *game, nil
}