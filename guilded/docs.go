package guilded

import (
	"encoding/json"
	"time"
)

type Docs struct {
	// The ID of the doc (min 1)
	DocsID int `json:"id"`

	// The server
	Server *Server `json:"-"`

	// The channel
	Channel *Channel `json:"-"`

	// The title of the doc (min length 1)
	Title string `json:"title"`

	// The content of the doc
	Content string `json:"content"`

	Mentions *Mention `json:"mentions"`

	// The ISO 8601 timestamp that the doc was created at
	CreatedAt time.Time `json:"createdAt"`

	// The user who created this doc
	CreatedBy *User `json:"-"`

	// The ISO 8601 timestamp that the doc was updated at, if relevant
	UpdatedAt time.Time `json:"updatedAt"`

	// The user who updated this doc
	UpdatedBy *User `json:"-"`
}

func (d *Docs) UnmarshalJSON(data []byte) error {
	type Alias Docs
	aux := &struct {
		*Alias
		ServerID  string `json:"serverId"`
		ChannelID string `json:"channelId"`
		CreatedBy string `json:"createdBy"`
		UpdatedBy string `json:"updatedBy"`
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.Server = &Server{ServerID: aux.ServerID}
	d.Channel = &Channel{ChannelID: aux.ChannelID}
	d.CreatedBy = &User{UserID: aux.CreatedBy}
	d.UpdatedBy = &User{UserID: aux.UpdatedBy}
	return nil
}
