package guilded

type Mention struct {
	// Users is the list of users mentioned in the message.
	Users []*User `json:"users"`

	// Channels is the list of channels mentioned in the message.
	Channels []*Channel `json:"channels"`

	// TODO: Implement Roles
	// Roles is the list of roles mentioned in the message.
	// Roles []*Role `json:"roles"`

	// Everyone is true if the message mentions everyone.
	Everyone bool `json:"everyone"`

	// Here is true if the message mentions here.
	Here bool `json:"here"`
}
