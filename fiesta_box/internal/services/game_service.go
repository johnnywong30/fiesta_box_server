package services

import (
	"fmt"
	"sync"

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
	mutex sync.Mutex // mutex around games map
} 


func (s *GameService) NewGame(c *websocket.Conn, room string) (*games.Game, error) {
	// get access to games map
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// check if room exists, fail if it does
	if _, ok := s.games[room]; ok {
		var g *games.Game
		return g, fmt.Errorf("game room %s already exists", room)
	}

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
		Mutex: sync.Mutex{},
	}
	// add game room to game service map
	s.games[room] = &game
	// return the created game room
	return &game, nil
}

func (s *GameService) AddToGame(c *websocket.Conn, room string) (*games.Game, error) {
	// get access to games map
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// check if room exists, fail if it doesn't
	game, ok := s.games[room]
	if !ok {
		var g *games.Game
		return g, fmt.Errorf("game room %s does not exist - failed to join game", room)
	} 
	
	// get access to game room
	game.Mutex.Lock()
	defer game.Mutex.Unlock()

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

	return game, nil
	
}

func (s *GameService) RemoveFromGame(c *websocket.Conn, room string) (*games.Game, error) {
	// get access to games map
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// check if room exists, fail if it doesn't
	game, ok := s.games[room]
	if !ok {
		var g *games.Game
		return g, fmt.Errorf("game room %s does not exist - failed to leave game", room)
	}

	// get access to game room
	game.Mutex.Lock()
	defer game.Mutex.Unlock()

	client, ok := game.Clients[c]
	if !ok {
		var g *games.Game
		return g, fmt.Errorf("game client does not exist in room %s - failed to leave game", room)
	}

	clientID := client.UserID
	
	// kick client from game room's client map
	delete(game.Clients, c)

	serverResponse := responses.SocketResponse{
		Status: responses.Success,
		Message: fmt.Sprintf("client %s left game %s", clientID, room),
		}

	game.Broadcast <- serverResponse

	return game, nil
}