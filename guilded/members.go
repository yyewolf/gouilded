package guilded

import "time"

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

type ServerMemberSummary struct {
	User *UserSummary `json:"user"`

	RoleIds []int `json:"roleIds"`
}
