package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

	"fiesta_box/internal/handlers"
	"fiesta_box/internal/models/messages"
)

var upgrader = websocket.Upgrader{}

func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()

	// Apply CORS middleware
	r.Use(s.corsMiddleware)

	r.HandleFunc("/", s.healthHandler)

	r.HandleFunc("/games/health", s.gameServiceHealthHandler)

	// Register websocket message handlers
	handlers.RegisterHandler(messages.MessageTypeStartGame, handlers.StartGameHandler)
	handlers.RegisterHandler(messages.MessageTypeTransferMaster, handlers.TransferMasterHandler)
	handlers.RegisterHandler(messages.MessageTypeConfigurePromptCount, handlers.ConfigurePromptHandler)
	handlers.RegisterHandler(messages.MessageTypeUseSavedPrompt, handlers.UseSavedPromptHandler)
	handlers.RegisterHandler(messages.MessageTypeWritePrompt, handlers.WritePromptHandler)
	handlers.RegisterHandler(messages.MessageTypeReceivePrompt, handlers.ReceivePromptHandler)
	handlers.RegisterHandler(messages.MessageTypePerformPrompt, handlers.PerformPromptHandler)
	handlers.RegisterHandler(messages.MessageTypeDrinkForPrompt, handlers.DrinkForPromptHandler)
	handlers.RegisterHandler(messages.MessageTypeChangePlayerName, handlers.ChangePlayerNameHandler)
	handlers.RegisterHandler(messages.MessageTypeJoinGame, handlers.JoinGameHandler)
	handlers.RegisterHandler(messages.MessageTypeLeaveGame, handlers.LeaveGameHandler)
	handlers.RegisterHandler(messages.MessageTypeCreateGame, handlers.CreateGameHandler)

	r.HandleFunc("/websocket", s.websocketHandler)

	// r.HandleFunc("/gorilla", s.GorillaHandler)


	return r
}

// CORS middleware
func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// CORS Headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Wildcard allows all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Credentials not allowed with wildcard origins

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}

func (s *Server) gameServiceHealthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.game.ServiceHealth())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResp)
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer c.Close()

	// Handle websocket connection

	go func() {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Error on reading message from client:", err)
				break
			}
			log.Printf("Received from client the message: %s", message)
	
			// Determine message type
			var clientMsg messages.Message
			err = json.Unmarshal(message, &clientMsg)
	
			if err != nil {
				log.Println("Error on parsing JSON message from client:", err)
				break
			}
	
			handlerArgs := handlers.HandlerFuncArgs{
				Message: clientMsg,
				GameService: s.game,
				Client: c,
			}
	
			response, err := handlers.HandleMessage(handlerArgs)
			if err != nil {
				log.Println("Error on handling message from client:", err)
				continue
			}
	
			responseJson, err := json.Marshal(response)
			if err != nil {
				log.Println("Error on parsing JSON message for response:", err)
				continue
			}
	
			err = c.WriteMessage(mt, responseJson)
			if err != nil {
				log.Println("Error on writing response to client:", err)
				break
			}
		}
	}()

	// Keep the connection alive with periodic pings
	for {
		if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
			log.Printf("Ping failed: %v", err)
			break
		}
		time.Sleep(30 * time.Second) // Send a ping every 30 seconds
	}
	
	
}
