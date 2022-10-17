package guilded

import "time"

type UserType string

const (
	UserTypeBot  UserType = "bot"
	UserTypeUser UserType = "user"
)

type User struct {
	// The ID of the user
	UserID string `json:"id"`

	// The type of user.
	// If this property is absent, it can assumed to be of type user
	Type UserType `json:"type"`

	// The name of the user
	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar"`

	// The banner image associated with the user
	Banner string `json:"banner"`

	// The ISO 8601 timestamp that the user was created at
	CreatedAt time.Time `json:"createdAt"`
}

type UserSummary struct {
	// The ID of the user
	UserID string `json:"id"`

	// The type of user.
	// If this property is absent, it can assumed to be of type user
	Type UserType `json:"type"`

	// The name of the user
	Name string `json:"name"`

	// The avatar image associated with the user
	Avatar string `json:"avatar"`
}
