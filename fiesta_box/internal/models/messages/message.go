package messages

type MessageType string

const (
	MessageTypeStartGame MessageType  			= "start_game"
	MessageTypeTransferMaster MessageType 		= "transfer_master"
	MessageTypeConfigurePromptCount MessageType = "configure_prompt_count"
	MessageTypeUseSavedPrompt MessageType 		= "use_saved_prompt"
	MessageTypeWritePrompt MessageType 			= "write_prompt"
	MessageTypeReceivePrompt MessageType 		= "receive_prompt"
	MessageTypePerformPrompt MessageType 		= "perform_prompt"
	MessageTypeDrinkForPrompt MessageType 		= "drink_for_prompt"
	MessageTypeChangePlayerName MessageType 	= "change_player_name"
)

type MessageInterface interface {
	GetType() MessageType
	GetContent() string	
}

type Message struct {
	Type MessageType `json:"type"`
	Content string `json:"content"`
}