package guilded

import (
	"encoding/json"
	"time"
)

type ForumTopic struct {
	// The ID of the forum topic
	ForumID int `json:"id"`

	// The server
	Server *Server `json:"-"`

	// The channel
	Channel *Channel `json:"-"`

	// The title of the forum topic (min length 1; max length 500)
	Title string `json:"title"`

	// The ISO 8601 timestamp that the forum topic was created at
	CreatedAt time.Time `json:"createdAt"`

	// The user who created this forum topic
	// Note: If this event has createdByWebhookId present, this field will still be populated, but can be ignored.
	// In this case, the value of this field will always be Ann6LewA
	CreatedBy *User `json:"-"`

	// The webhook who created this forum topic, if it was created by a webhook
	CreatedByWebhook *Webhook `json:"-"`

	// The ISO 8601 timestamp that the forum topic was updated at, if relevant
	UpdatedAt time.Time `json:"updatedAt"`

	// The ISO 8601 timestamp that the forum topic was bumped at.
	// This timestamp is updated whenever there is any activity on the posts within the forum topic.
	BumpedAt time.Time `json:"bumpedAt"`

	// default false
	IsPinned bool `json:"isPinned"`

	// default false
	IsLocked bool `json:"isLocked"`

	// The content of the forum topic
	Content string `json:"content"`

	Mentions *Mention `json:"mentions"`
}

func (f *ForumTopic) UnmarshalJSON(data []byte) error {
	type Alias ForumTopic
	aux := &struct {
		*Alias
		ServerID           string `json:"serverId"`
		ChannelID          string `json:"channelId"`
		CreatedBy          string `json:"createdBy"`
		CreatedByWebhookID string `json:"createdByWhebhookId"`
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	f.CreatedByWebhook = &Webhook{WebhookID: aux.CreatedByWebhookID}
	f.CreatedBy = &User{UserID: aux.CreatedBy}
	f.Channel = &Channel{ChannelID: aux.ChannelID}
	f.Server = &Server{ServerID: aux.ServerID}
	return nil
}
