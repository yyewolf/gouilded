package guilded

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type serverMemberTest struct {
	ServerMember *ServerMember `json:"serverMember,omitempty"`
}

func TestServerMember_UnmarshallJSON(t *testing.T) {
	sm := serverMemberTest{
		ServerMember: &ServerMember{},
	}

	data := []byte(`{
		"user": {
		  "id": "Ann6LewA",
		  "type": "user",
		  "name": "Leopold Stotch",
		  "createdAt": "2021-06-15T20:15:00.706Z"
		},
		"roleIds": [],
		"nickname": "Professor Chaos",
		"joinedAt": "2021-07-15T20:15:00.706Z"
	  }`)

	if err := json.Unmarshal(data, sm.ServerMember); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, sm.ServerMember.User.UserID, "Ann6LewA")
	assert.Equal(t, sm.ServerMember.Nickname, "Professor Chaos")
}

type serverMemberBanTest struct {
	ServerMemberBan *ServerMemberBan
}

func TestServerMemberBan_UnmarshallJSON(t *testing.T) {
	smb := serverMemberBanTest{
		ServerMemberBan: &ServerMemberBan{},
	}

	data := []byte(`{
		"user": {
		  "id": "Ann6LewA",
		  "type": "user",
		  "name": "Leopold Stotch"
		},
		"reason": "More toxic than a poison Pokémon",
		"createdAt": "2021-06-15T20:15:00.706Z",
		"createdBy": "Ann6LewA"
	  }`)

	if err := smb.ServerMemberBan.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, smb.ServerMemberBan.CreatedBy.UserID, "Ann6LewA")
	assert.Equal(t, smb.ServerMemberBan.User.UserID, "Ann6LewA")
	assert.Equal(t, smb.ServerMemberBan.User.Name, "Leopold Stotch")
}
