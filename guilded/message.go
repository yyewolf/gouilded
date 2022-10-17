package guilded

import (
	"encoding/json"
	"time"
)

type MessageType string

const (
	MessageTypeDefault MessageType = "default"
	MessageTypeSystem  MessageType = "system"
)

type Message struct {
	// ID is the ID of the message
	ID string `json:"id"`

	// Type is the type of the message
	Type MessageType `json:"type"`

	// ServerID is the ID of the server the message was sent in
	ServerID string `json:"serverId"`

	// ChannelID is the ID of the channel the message was sent in
	ChannelID string `json:"channelId"`

	// Content is the content of the message
	Content string `json:"content"`

	// Embeds is the list of embeds in the message
	Embeds []*MessageEmbed `json:"embeds"`

	// ReplyMessageIDs is the list of message IDs this message is replying to
	ReplyMessageIDs []string `json:"replyMessageIds"`

	// IsPrivate, if true, indicates that the message is only seen by those mentioned
	IsPrivate bool `json:"isPrivate"`

	// IsSilent, if true, indicates that the message did not ping
	IsSilent bool `json:"isSilent"`

	// Mentions is the list of users mentioned in the message
	Mentions []*Mention `json:"mentions"`

	// CreatedAt is the time the message was created at in ISO 8601 format
	CreatedAt time.Time `json:"createdAt"`

	// Author is the user who sent the message
	// If the message was sent by a webhook, this will be set to 'Ann6LewA'
	Author *User `json:"-"`

	// CreatedByWebhookID is the ID of the webhook that created the message
	CreatedByWebhook *Webhook `json:"-"`

	// UpdatedAt is the time the message was last updated at in ISO 8601 format
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m *Message) UnmarshalJSON(b []byte) error {
	type Alias Message
	aux := &struct {
		Author string `json:"createdBy"`
		*Alias
		CreatedByWebhookID string `json:"createdByWebhookId"`
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	m.Author = &User{UserID: aux.Author}
	m.CreatedByWebhook = &Webhook{WebhookID: aux.CreatedByWebhookID}

	return nil
}

func (m *Message) MarshalJSON() ([]byte, error) {
	type Alias Message
	aux := &struct {
		Author string `json:"createdBy"`
		*Alias
	}{
		Alias:  (*Alias)(m),
		Author: m.Author.UserID,
	}

	return json.Marshal(aux)
}
