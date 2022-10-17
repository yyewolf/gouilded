package guilded

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type channelTest struct {
	Channel *Channel `json:"message,omitempty"`
}

func TestChannel_UnmarshalJSON(t *testing.T) {
	v := channelTest{
		Channel: &Channel{},
	}

	data := []byte(`{
		"priority": 4,
		"id": "d2242862-401c-42f8-9ce2-0c75977475a6",
		"type": "Team",
		"name": "general",
		"description": "wacky hi-jinks ensue!",
		"settings": null,
		"roles": null,
		"rolesById": {},
		"tournamentRolesById": {},
		"teamId": "4R5q39VR",
		"channelCategoryId": null,
		"addedAt": null,
		"channelId": "d2242862-401c-42f8-9ce2-0c75977475a6",
		"isRoleSynced": null,
		"isPublic": false,
		"groupId": "WD56qLmd",
		"createdAt": "2020-07-31T18:51:19.563Z",
		"createdBy": "EdVMVKR4",
		"updatedAt": "2020-07-31T18:51:19.563Z",
		"contentType": "chat",
		"archivedAt": null,
		"parentChannelId": null,
		"autoArchiveAt": null,
		"deletedAt": null,
		"archivedBy": null,
		"createdByWebhookId": null,
		"archivedByWebhookId": null,
		"userPermissions": null,
		"tournamentRoles": null,
		"voiceParticipants": [],
		"userStreams": []
	}`)

	if err := json.Unmarshal(data, v.Channel); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, v.Channel.ChannelID, "d2242862-401c-42f8-9ce2-0c75977475a6")
	assert.Equal(t, v.Channel.Type, ChannelTypeTeam)
	assert.Equal(t, v.Channel.Name, "general")
}
