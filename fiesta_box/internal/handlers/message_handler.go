package handlers

import (
	"fmt"

	"github.com/gorilla/websocket"

	"fiesta_box/internal/models/games"
	"fiesta_box/internal/models/messages"
	"fiesta_box/internal/models/responses"
	"fiesta_box/internal/services"
)


type HandlerFuncArgs struct {
	Message messages.Message
	GameService *services.GameService
	Client *websocket.Conn
}


type HandlerFunc func(HandlerFuncArgs) (responses.SocketResponse, error)

var HandlerRegistry = make(map[messages.MessageType]HandlerFunc)

func RegisterHandler(messageType messages.MessageType, handler HandlerFunc) {
	HandlerRegistry[messageType] = handler
}

func HandleMessage(args HandlerFuncArgs) (responses.SocketResponse, error) {
	handler, ok := HandlerRegistry[args.Message.Type]
	if !ok {
		return responses.SocketResponse{
			Status: responses.UnknownMessageType,
			Message: "Unknown message type",
		}, nil
	}
	return handler(args)
}

func StartGameHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id

	// TODO: Start game logic
	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Started game!",
	}
	return response, nil
}

func TransferMasterHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id

	// TODO: Transfer master logic
	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Transferred master!",
	}
	return response, nil
}

func ConfigurePromptHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt count

	// TODO: Configure prompt logic
	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Configured prompt count to be N.",
	}
	return response, nil
}

func UseSavedPromptHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt id

	// TODO: Use saved prompt logic

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Used saved prompt X.",
	}
	return response, nil
}

func WritePromptHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt string

	// TODO: Write prompt logic; each prompt should have a uid

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Wrote prompt X.",
	}
	return response, nil
}

func ReceivePromptHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id

	// TODO: Receive prompt logic

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Received prompt X.",
	}
	return response, nil
}

func PerformPromptHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt id

	// TODO: Perform prompt logic

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Performed prompt X.",
	}
	return response, nil
}

func DrinkForPromptHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt id

	// TODO: Drink for prompt logic

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Drank for prompt X.",
	}
	return response, nil
}

func ChangePlayerNameHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, player name

	// TODO: Change player name logic	

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Changed player name to X.",
	}
	return response, nil
}

func JoinGameHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string

	// TODO: Join game logic	

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Joined game X.",
	}
	return response, nil
}

func LeaveGameHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	
	// TODO: Leave game logic	

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Left game X.",
	}
	return response, nil
}

func CreateGameHandler(args HandlerFuncArgs) (responses.SocketResponse, error) {
	done := make(chan *games.Game)

	go args.GameService.NewGame(args.Client, done)

	createdGame := <- done

	content := map[string]interface{}{
        "gameID": createdGame.Room,
    }

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: fmt.Sprintf("Created game %s", createdGame.Room),
		Content: content,
	}
	return response, nil
}