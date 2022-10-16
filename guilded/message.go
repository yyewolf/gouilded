package guilded

import (
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

	// ChannelID is the ID of the channel the message was sent in
	ChannelID string `json:"channelId"`

	// Author is the User who created the message
	// Author *User `json:"createdBy"`

	// Content is the content of the message
	// Content *MessageContent `json:"content"`

	// CreatedAt is the time the message was created at in ISO 8601 format
	CreatedAt time.Time `json:"createdAt"`

	// EditedAt is the time the message was edited at in ISO 8601 format
	EditedAt time.Time `json:"editedAt"`

	// DeletedAt is the time the message was deleted at in ISO 8601 format
	DeletedAt time.Time `json:"deletedAt"`

	// Reactions is a list of reactions on the message
	// Reactions []*Reaction `json:"reactions"`

	// IsPinned is true if the message is pinned
	IsPinned bool `json:"isPinned"`

	// PinnedBy is the User who pinned the message
	// PinnedBy *User `json:"pinnedBy"`

	// WebhookID is the ID of the webhook that sent the message
	WebhookID string `json:"webhookId"`

	// Type is the type of the message
	Type MessageType `json:"type"`
}
