package models

import (
	"encoding/json"
	"fmt"
)

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

type Message interface {
	GetType() MessageType	
}

type ActionMessage interface {
	Message // embed Message interface
	GetContent() struct 
}

type TransferMasterMessageContent struct {
	UserId string `json:"userId"`
}

type TransferMasterMessage struct {
	Type MessageType `json:"type"`
	Content string `json:"content"`
}

func (m TransferMasterMessage) GetType() MessageType {
	return MessageTypeTransferMaster
}

func (m TransferMasterMessage) GetContent() TransferMasterMessageContent {
	var content TransferMasterMessageContent
	err := json.Unmarshal([]byte(m.Content), &content)
	if err != nil {
		fmt.Errorf(err) // handle error
	}
	return content
}