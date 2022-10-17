package gateway

import "encoding/json"

// Raw messages received by the API
type Message struct {
	// Op is the operation code of the message
	Op int `json:"op"`

	// D is the data of the message
	D    MessageData     `json:"d"`
	RawD json.RawMessage `json:"-"`

	// S is the sequence ID
	S string `json:"s"`

	// T is the event name
	T string `json:"t"`
}

type MessageData interface {
	messageData()
}

func UnmarshalEventData(data []byte, eventType EventType) (EventData, error) {
	var (
		eventData EventData
		err       error
	)
	switch eventType {
	case EventTypeChatMessageCreated:
		var d EventChatMessageCreated
		err = json.Unmarshal(data, &d)
		eventData = d
	}

	return eventData, err
}
