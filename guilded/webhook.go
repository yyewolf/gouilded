package guilded

import (
	"encoding/json"
	"time"
)

type Webhook struct {
	// The ID of the webhook
	WebhookID string `json:"id"`

	// The name of the webhook (min length 1; max length 128)
	Name string `json:"name"`

	// The server
	Server *Server `json:"-"`

	// The channel
	Channel *Channel `json:"-"`

	// The ISO 8601 timestamp that the webhook was created at
	CreatedAt time.Time `json:"createdAt"`

	// The user who created this webhook
	CreatedBy *User `json:"-"`

	// The ISO 8601 timestamp that the webhook was deleted at
	DeletedAt time.Time `json:"deletedAt"`

	// The token of the webhook
	Token string `json:"token"`
}

func (w *Webhook) UnmarshalJSON(data []byte) error {
	type Alias Webhook
	aux := &struct {
		*Alias
		ServerID  string `json:"serverId"`
		ChannelID string `json:"channelId"`
		CreatedBy string `json:"createdBy"`
	}{
		Alias: (*Alias)(w),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	w.CreatedBy = &User{UserID: aux.CreatedBy}
	w.Channel = &Channel{ChannelID: aux.ChannelID}
	w.Server = &Server{ServerID: aux.ServerID}
	return nil
}
