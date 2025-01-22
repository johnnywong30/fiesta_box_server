package handlers

import (
	"fiesta_box/internal/models/messages"
	"fiesta_box/internal/models/responses"
)


type HandlerFunc func(messages.Message) (responses.SocketResponse, error)

var HandlerRegistry = make(map[messages.MessageType]HandlerFunc)

func RegisterHandler(messageType messages.MessageType, handler HandlerFunc) {
	HandlerRegistry[messageType] = handler
}

func HandleMessage(message messages.Message) (responses.SocketResponse, error) {
	handler, ok := HandlerRegistry[message.Type]
	if !ok {
		return responses.SocketResponse{
			Status: responses.UnknownMessageType,
			Message: "Unknown message type",
		}, nil
	}
	return handler(message)
}

func StartGameHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id

	// TODO: Start game logic
	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Started game!",
	}
	return response, nil
}

func TransferMasterHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id

	// TODO: Transfer master logic
	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Transferred master!",
	}
	return response, nil
}

func ConfigurePromptHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt count

	// TODO: Configure prompt logic
	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Configured prompt count to be N.",
	}
	return response, nil
}

func UseSavedPromptHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt id

	// TODO: Use saved prompt logic

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Used saved prompt X.",
	}
	return response, nil
}

func WritePromptHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt string

	// TODO: Write prompt logic; each prompt should have a uid

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Wrote prompt X.",
	}
	return response, nil
}

func ReceivePromptHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id

	// TODO: Receive prompt logic

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Received prompt X.",
	}
	return response, nil
}

func PerformPromptHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt id

	// TODO: Perform prompt logic

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Performed prompt X.",
	}
	return response, nil
}

func DrinkForPromptHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, prompt id

	// TODO: Drink for prompt logic

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Drank for prompt X.",
	}
	return response, nil
}

func ChangePlayerNameHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id, player name

	// TODO: Change player name logic	

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Changed player name to X.",
	}
	return response, nil
}

func JoinGameHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string

	// TODO: Join game logic	

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Joined game X.",
	}
	return response, nil
}

func LeaveGameHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string to get game id

	// TODO: Leave game logic	

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Left game X.",
	}
	return response, nil
}

func CreateGameHandler(msg messages.Message) (responses.SocketResponse, error) {
	// TODO: unmarshal content JSON string

	// TODO: Create game logic	

	response := responses.SocketResponse{
		Status: responses.Success,
		Message: "Created game X.",
	}
	return response, nil
}