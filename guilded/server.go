package guilded

import (
	"encoding/json"
	"time"
)

type ServerType string

const (
	ServerTypeTeam         ServerType = "team"
	ServerTypeOrganization ServerType = "organization"
	ServerTypeCommunity    ServerType = "community"
	ServerTypeClan         ServerType = "clan"
	ServerTypeGuild        ServerType = "guild"
	ServerTypeFriends      ServerType = "friends"
	ServerTypeStreaming    ServerType = "streaming"
	ServerTypeOther        ServerType = "other"
)

type Server struct {
	// The id of the server
	ID string `json:"id"`

	// The ID of the user who created this server
	CreatedBy *User `json:"-"`

	// The type of server designated from the server's settings page
	Type ServerType `json:"type"`

	// The name given to the server
	Name string `json:"name"`

	// The URL that the server can be accessible from.
	// For example, a value of "Guilded-Official" means
	// the server can be accessible from https://www.guilded.gg/Guilded-Official
	Url string `json:"url"`

	// The description associated with the server
	About string `json:"about"`

	// The avatar image associated with the server
	Avatar string `json:"avatar"`

	// The banner image associated with the server
	Banner string `json:"banner"`

	// The timezone associated with the server
	Timezone string `json:"timezone"`

	// The verified status of the server
	IsVerified bool `json:"isVerified"`

	// The channel ID of the default channel of the server. Useful for welcome messages
	DefaultChannelId string `json:"defaultChannelId"`

	// The ISO 8601 timestamp that the server was created at
	CreatedAt time.Time `json:"createdAt"`
}

func (s *Server) UnmarshalJSON(data []byte) error {
	type Alias Server
	aux := &struct {
		*Alias
		OwnerID string `json:"ownerId"`
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	s.CreatedBy = &User{UserID: aux.OwnerID}
	return nil
}
