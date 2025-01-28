package services

import (
	"fmt"
	"log"
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

type GameServiceState struct {
	Games int `json:"games"`
	GameStates map[string]games.GameState `json:"gameStates"`

}

func NewGameService() *GameService {
	return &GameService{
		games: map[string]*games.Game{},
		mutex: sync.Mutex{},
	}
}

func (s *GameService) CreateGameClient(c *websocket.Conn, room string) *games.GameClient {
	client := games.GameClient{
		Room: room,
		Client: c,
		UserID: uuid.NewString(),
		Connected: true,
	}
	log.Printf("Created game client %s", client.UserID)
	return &client
}


func (s *GameService) NewGame(c *websocket.Conn, done chan *games.Game) (*games.Game, error) {
	// get access to games map
	log.Print("[NewGame] - Getting gameService lock")
	s.mutex.Lock()
	defer s.mutex.Unlock()
	defer log.Print("[NewGame] - Releasing gameService lock")

	room := uuid.NewString()

	// check if room exists, fail if it does
	if _, ok := s.games[room]; ok {
		var g *games.Game
		return g, fmt.Errorf("game room %s already exists", room)
	}

	// create game client for this websocket connection
	client := s.CreateGameClient(c, room)

	// add game client to the game room's client map
	clients := map[*websocket.Conn]*games.GameClient{
		c: client,
	}

	// create game room
	game := games.Game{
		Clients: clients,
		Broadcast: make(chan responses.SocketResponse),
		Status: games.NotStarted,
		Mutex: sync.Mutex{},
		Room: room,
	}
	// add game room to game service map
	s.games[room] = &game
	log.Printf("Created game room %s", game.Room)

	done <- &game

	// return the created game room
	return &game, nil
}

func (s *GameService) AddToGame(c *websocket.Conn, room string, done chan bool) (*games.Game, error) {
	// get access to games map
	log.Print("[AddToGame] - Getting gameService lock")
	s.mutex.Lock()
	defer s.mutex.Unlock()
	defer log.Print("[AddToGame] - Releasing gameService lock")

	// check if room exists, fail if it doesn't
	game, ok := s.games[room]
	if !ok {
		var g *games.Game
		err := fmt.Errorf("game room %s does not exist - failed to join game", room)
		log.Print(err.Error())
		done <- false
		return g, err
	} 
	
	// get access to game room
	log.Printf("[AddToGame] - Getting game %s lock", game.Room)
	game.Mutex.Lock()
	defer game.Mutex.Unlock()
	defer log.Printf("[AddToGame] - Releasing game %s lock", game.Room)

	client := s.CreateGameClient(c, room)
	game.Clients[c] = client

	message := fmt.Sprintf("client %s joined game %s", client.UserID, room)
	log.Print(message)

	// serverResponse := responses.SocketResponse{
	// 	Status: responses.Success,
	// 	Message: message,
	// 	}

	// game.Broadcast <- serverResponse
	done <- true

	return game, nil
	
}

func (s *GameService) RemoveFromGame(c *websocket.Conn, room string, done chan bool) (*games.Game, error) {
	// get access to games map
	log.Print("[RemoveFromGame] - Getting gameService lock")
	s.mutex.Lock()
	defer s.mutex.Unlock()
	defer log.Print("[RemoveFromGame] - Releasing gameService lock")

	// check if room exists, fail if it doesn't
	game, ok := s.games[room]
	if !ok {
		var g *games.Game
		done <- false
		err := fmt.Errorf("game room %s does not exist - failed to leave game", room)
		log.Print(err.Error())
		return g, err
	}

	// get access to game room
	log.Printf("[RemoveFromGame] - Getting game %s lock", game.Room)
	game.Mutex.Lock()
	defer game.Mutex.Unlock()
	defer log.Printf("[RemoveFromGame] - Releasing game %s lock", game.Room)

	client, ok := game.Clients[c]
	if !ok {
		var g *games.Game
		done <- false
		err := fmt.Errorf("game client does not exist in room %s - failed to leave game", room)
		log.Print(err.Error())
		return g, err
	}

	clientID := client.UserID
	
	// kick client from game room's client map
	delete(game.Clients, c)

	message := fmt.Sprintf("Client %s left game room %s", clientID, room)

	log.Print(message)

	// if len(game.Clients) == 0 {
	// 	// remove game room from game service map if no clients remain
	// 	defer delete(s.games, room)
	// 	log.Printf("Deleted game room %s. No players remaining.", room)
	// }

	// serverResponse := responses.SocketResponse{
	// 	Status: responses.Success,
	// 	Message: message,
	// 	}

	// game.Broadcast <- serverResponse
	done <- true

	return game, nil
}

func (s *GameService) ServiceHealth() GameServiceState {
	// get access to games map
	log.Print("[ServiceHealth] - Getting gameService lock")
	s.mutex.Lock()
	defer s.mutex.Unlock()
	defer log.Print("[ServiceHealth] - Releasing gameService lock")


	gameStates := make(map[string]games.GameState)

	for room, game := range s.games {
		gameStates[room] = games.GameState{
			Clients: len(game.Clients),
			Status: game.Status,
			Room: room,
		}
	}

	games := len(s.games)

	return GameServiceState{
		Games: games,
		GameStates: gameStates,
	}
}