package guilded

import (
	"encoding/json"
	"time"
)

type ServerMember struct {
	User *User `json:"user"`

	// must have unique items true
	RoleIds []int `json:"roleIds"`

	Nickname string `json:"nickname"`

	// The ISO 8601 timestamp that the member was created at
	JoinedAt time.Time `json:"joinedAt"`

	// default false
	IsOwner bool `json:"isOwner"`
}

type ServerMemberBan struct {
	User *User `json:"user"`

	// The reason for the ban as submitted by the banner
	Reason string `json:"reason"`

	// The user who created this server member ban
	CreatedBy *User `json:"-"`

	// The ISO 8601 timestamp that the server member ban was created at
	CreatedAt time.Time `json:"createdAt"`
}

func (smb *ServerMemberBan) UnmarshalJSON(data []byte) error {
	type Alias ServerMemberBan
	aux := &struct {
		*Alias
		CreatedBy string `json:"createdBy"`
	}{
		Alias: (*Alias)(smb),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	smb.CreatedBy = &User{UserID: aux.CreatedBy}
	return nil
}
