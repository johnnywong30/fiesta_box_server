package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (change this for production)
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	// Echo messages back to the client
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received: %s\n", message)

		err = conn.WriteMessage(messageType, message)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	port := "8080"
	fmt.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
