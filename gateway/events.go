package gateway

import "github.com/yyewolf/gouilded/guilded"

type EventData interface {
	MessageData
	eventData()
}

// EventChatMessageCreated is the event sent by Guilded when a message is created.
type EventChatMessageCreated struct {
	// ServerID is the Server in which the event occured
	ServerID string `json:"serverId"`

	// Message is the message that was created
	Message *guilded.Message `json:"message"`
}

func (EventChatMessageCreated) messageData() {}
func (EventChatMessageCreated) eventData()   {}
