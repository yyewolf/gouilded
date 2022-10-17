package guilded

import (
	"encoding/json"
	"time"
)

type ChannelType string

const (
	ChannelTypeTeam ChannelType = "Team"
	ChannelTypeDM   ChannelType = "DM"
)

type ChannelContentType string

const (
	ChannelContentTypeAnnouncements ChannelContentType = "announcements"
	ChannelContentTypeChat          ChannelContentType = "chat"
	ChannelContentTypeCalendar      ChannelContentType = "calendar"
	ChannelContentTypeDocuments     ChannelContentType = "docs"
	ChannelContentTypeForums        ChannelContentType = "forums"
	ChannelContentTypeMedia         ChannelContentType = "media"
	ChannelContentTypeList          ChannelContentType = "list"
	ChannelContentTypeScheduling    ChannelContentType = "scheduling"
	ChannelContentTypeStream        ChannelContentType = "stream"
	ChannelContentTypeVoice         ChannelContentType = "voice"
)

type Channel struct {
	// ID is the unique identifier for the channel.
	ChannelID string `json:"id"`

	// Type is the type of channel.
	Type ChannelType `json:"type"`

	// Name is the name of the channel.
	Name string `json:"name"`

	// Topic is the topic of the channel.
	Topic string `json:"topic"`

	// CreatedAt is the time the channel was created.
	CreatedAt time.Time `json:"createdAt"`

	// CreatedBy is the user who created the channel.
	CreatedBy *User `json:"-"`

	// UpdatedAt is the time the channel was last updated.
	UpdatedAt time.Time `json:"updatedAt"`

	// ServerID is the ID of the server the channel is in.
	ServerID string `json:"serverId"`

	// ParentID is the ID of the parent channel.
	ParentID string `json:"parentId"`

	// CategoryID is the ID of the category the channel is in.
	CategoryID int `json:"categoryId"`

	// GroupID is the ID of the group the channel is in.
	GroupID string `json:"groupId"`

	// IsPublic is whether the channel is public.
	IsPublic bool `json:"isPublic"`

	// ArchivedBy is the user who archived the channel.
	ArchivedBy *User `json:"-"`

	// ArchivedAt is the time the channel was archived.
	ArchivedAt time.Time `json:"archivedAt"`
}

func (c *Channel) UnmarshalJSON(data []byte) error {
	type Alias Channel
	aux := &struct {
		*Alias
		ArchivedBy string `json:"archivedBy"`
		CreatedBy  string `json:"createdBy"`
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.CreatedBy = &User{UserID: aux.CreatedBy}
	c.ArchivedBy = &User{UserID: aux.ArchivedBy}
	return nil
}
