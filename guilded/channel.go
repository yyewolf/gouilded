package guilded

import (
	"time"
)

type ChannelType string

const (
	ChannelTypeTeam ChannelType = "Team"
	ChannelTypeDM   ChannelType = "DM"
)

type ChannelContentType string

const (
	ChannelContentTypeAnnouncement ChannelContentType = "announcement"
	ChannelContentTypeChat         ChannelContentType = "chat"
	ChannelContentTypeDocument     ChannelContentType = "doc"
	ChannelContentTypeForum        ChannelContentType = "forum"
	ChannelContentTypeMedia        ChannelContentType = "media"
	ChannelContentTypeList         ChannelContentType = "list"
	ChannelContentTypeScheduling   ChannelContentType = "scheduling"
	ChannelContentTypeStreaming    ChannelContentType = "streaming"
	ChannelContentTypeVoice        ChannelContentType = "voice"
)

type Channel struct {
	// ID is the ID of the channel
	ID string `json:"id"`

	// Type is the type of the channel
	Type ChannelType `json:"type"`

	// CreatedAt is the time the channel was created at in ISO 8601 format
	CreatedAt time.Time `json:"createdAt"`

	// CreatedBy is the User who created the channel
	// CreatedBy *User `json:"createdBy"`

	// UpdatedAt is the time the channel was updated at in ISO 8601 format
	UpdatedAt time.Time `json:"updatedAt"`

	// Name is the name of the channel
	Name string `json:"name"`

	// Content Type is the type of content the channel is for
	ContentType ChannelContentType `json:"contentType"`

	// ArchivedAt is the time the channel was archived at in ISO 8601 format
	ArchivedAt time.Time `json:"archivedAt"`

	// ParentChannelID is the ID of the parent channel
	ParentChannelID string `json:"parentChannelId"`

	// AutoArchiveAt is the time the channel will auto-archive at in ISO 8601 format (only for threads)
	AutoArchiveAt time.Time `json:"autoArchiveAt"`

	// DeletedAt is the time the channel was deleted at in ISO 8601 format
	DeletedAt time.Time `json:"deletedAt"`

	// ArchivedBy is the User who archived the channel
	// ArchivedBy *User `json:"archivedBy"`

	// Description is the description of the channel
	Description string `json:"description"`

	// CreatedByWebhookID is the ID of the webhook that created the channel
	CreatedByWebhookID string `json:"createdByWebhookId"`

	// ArchivedByWebhookID is the ID of the webhook that archived the channel
	ArchivedByWebhookID string `json:"archivedByWebhookId"`

	// TeamID is the ID of the server the channel is in
	TeamID string `json:"teamId"`

	// ChannelCategoryID is the ID of the channel category the channel is in
	ChannelCategoryID string `json:"channelCategoryId"`

	// AddedAt is the time the channel was added at in ISO 8601 format
	AddedAt time.Time `json:"addedAt"`

	// ChannelID is the ID of the channel
	ChannelID string `json:"channelId"`

	// IsRoleSynced is true if this channel's roles are synced with the linked Discord guild's roles
	IsRoleSynced bool `json:"isRoleSynced"`

	// Roles is a list of roles that can access the channel
	// Roles []*Role `json:"roles"`

	// UserPermissions is a list of users who have overwrites in this channel
	// UserPermissions []*User `json:"userPermissions"`

	// TournamentRoles ??
	// TournamentRoles []*TournamentRole `json:"tournamentRoles"`

	// IsPublic is true if the channel is public
	IsPublic bool `json:"isPublic"`

	// Priority ??
	Priority int `json:"priority"`

	// GroupID is the ID of the group the channel is in (for team channels)
	GroupID string `json:"groupId"`

	// Settings ???
	// Settings *ChannelSettings `json:"settings"`

	// GroupType ????
	GroupType string `json:"groupType"`

	// RolesByID mapping of roleID to role with undescript permissions
	// RolesByID map[string]*Role `json:"rolesById"`

	// TournamentRolesByID mapping of tournamentRoleID to tournamentRole
	// TournamentRolesByID map[string]*TournamentRole `json:"tournamentRolesById"`

	// CreatedByInfo contains partial info about the user who created the channel
	// CreatedByInfo *User `json:"createdByInfo"`
}
