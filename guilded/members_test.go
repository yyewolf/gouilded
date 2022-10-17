package guilded

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type serverMemberTest struct {
	ServerMember *ServerMember `json:"serverMember,omitempty"`
}

type serverMemberSummaryTest struct {
	ServerMemberSummary *ServerMemberSummary `json:"serverMemberSummary,omitempty"`
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

func TestServerMemberSummary_UnmarshallJSON(t *testing.T) {
	sms := serverMemberSummaryTest{
		ServerMemberSummary: &ServerMemberSummary{},
	}

	data := []byte(`{
		"user": {
		  "id": "Ann6LewA",
		  "type": "user",
		  "name": "Leopold Stotch"
		},
		"roleIds": []
	  }`)

	if err := json.Unmarshal(data, sms.ServerMemberSummary); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, sms.ServerMemberSummary.User.UserID, "Ann6LewA")
	assert.Equal(t, sms.ServerMemberSummary.User.Name, "Leopold Stotch")
}
