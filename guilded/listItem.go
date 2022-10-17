package guilded

import (
	"encoding/json"
	"time"
)

type ListItem struct {
	// The ID of the list item
	ListItemID string `json:"id"`

	// The server
	Server *Server `json:"-"`

	// The channel
	Channel *Channel `json:"-"`

	// The message of the list item
	Message string `json:"message"`

	Mentions *Mention `json:"mentions"`

	// The ISO 8601 timestamp that the list item was created at
	CreatedAt time.Time `json:"createdAt"`

	// The ID of the user who created this list item
	// Note: If this event has createdByWebhookId present, this field will still be populated, but can be ignored.
	// In this case, the value of this field will always be Ann6LewA
	CreatedBy *User `json:"-"`

	// The webhook who created this list item, if it was created by a webhook
	CreatedByWebhook *Webhook `json:"-"`

	// The ISO 8601 timestamp that the list item was updated at, if relevant
	UpdatedAt time.Time `json:"updatedAt"`

	// The user who updated this list item
	UpdatedBy *User `json:"-"`

	// The parent list item if this list item is nested
	ParentListItem *ListItem `json:"-"`

	// The ISO 8601 timestamp that the list item was completed at
	CompletedAt time.Time `json:"completedAt"`

	// The user who completed this list item
	CompletedBy *User `json:"-"`

	Note *Note `json:"note"`
}

type Note struct {
	// The ISO 8601 timestamp that the note was created at.
	// If this field is populated, then there's a note associated with the list item
	CreatedAt time.Time `json:"createdAt"`

	// The user who created this note
	CreatedBy *User `json:"-"`

	// The ISO 8601 timestamp that the list item was updated at, if relevant
	UpdatedAt time.Time `json:"updatedAt"`

	// The user who updated this list item
	UpdatedBy *User `json:"-"`

	Mentions *Mention `json:"mentions"`

	Content string `json:"content"`
}

func (l *ListItem) UnmarshalJSON(data []byte) error {
	type Alias ListItem
	aux := &struct {
		*Alias
		ServerID           string `json:"serverId"`
		ChannelID          string `json:"channelId"`
		CreatedBy          string `json:"createdBy"`
		CreatedByWebhookID string `json:"createdByWhebhookId"`
		UpdatedBy          string `json:"updatedBy"`
		ParentListItemId   string `json:"parentListItemId"`
		CompletedBy        string `json:"completedBy"`
	}{
		Alias: (*Alias)(l),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	l.Server = &Server{ServerID: aux.ServerID}
	l.Channel = &Channel{ChannelID: aux.ChannelID}
	l.CreatedBy = &User{UserID: aux.CreatedBy}
	l.CreatedByWebhook = &Webhook{WebhookID: aux.CreatedByWebhookID}
	l.UpdatedBy = &User{UserID: aux.UpdatedBy}
	l.ParentListItem = &ListItem{ListItemID: aux.ParentListItemId}
	l.CompletedBy = &User{UserID: aux.CreatedBy}
	return nil
}

func (n *Note) UnmarshalJSON(data []byte) error {
	type Alias Note
	aux := &struct {
		*Alias
		CreatedBy string `json:"createdBy"`
		UpdatedBy string `json:"updatedBy"`
	}{
		Alias: (*Alias)(n),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	n.CreatedBy = &User{UserID: aux.CreatedBy}
	n.UpdatedBy = &User{UserID: aux.UpdatedBy}
	return nil
}
