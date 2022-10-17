package guilded

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type userTest struct {
	User *User `json:"user,omitempty"`
}

type userSummaryTest struct {
	UserSummary *UserSummary `json:"userSummary,omitempty"`
}

func TestUser_UnmarshallJSON(t *testing.T) {
	u := userTest{
		User: &User{},
	}

	data := []byte(`{
		"id": "Ann6LewA",
		"type": "user",
		"name": "Leopold Stotch",
		"createdAt": "2021-06-15T20:15:00.706Z"
	  }`)

	if err := json.Unmarshal(data, u.User); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, u.User.UserID, "Ann6LewA")
	assert.Equal(t, u.User.Name, "Leopold Stotch")
}

func TestUserSummary_UnmarshallJSON(t *testing.T) {
	u := userSummaryTest{
		UserSummary: &UserSummary{},
	}

	data := []byte(`{
		"id": "Ann6LewA",
		"type": "user",
		"name": "Leopold Stotch"
	  }`)

	if err := json.Unmarshal(data, u.UserSummary); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, u.UserSummary.UserID, "Ann6LewA")
	assert.Equal(t, u.UserSummary.Name, "Leopold Stotch")
}
