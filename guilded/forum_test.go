package guilded

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type forumTopicTest struct {
	ForumTopic *ForumTopic
}

func TestForumTopic_UnmarshallJSON(t *testing.T) {
	f := forumTopicTest{
		ForumTopic: &ForumTopic{},
	}

	data := []byte(`{
		"id": 123456,
		"serverId": "wlVr3Ggl",
		"channelId": "00000000-0000-0000-0000-000000000000",
		"title": "Welcome new members!!",
		"createdAt": "2021-06-15T20:15:00.706Z",
		"createdBy": "Ann6LewA",
		"content": "Please introduce yourself in this topic!!!"
	  }`)

	if err := f.ForumTopic.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, f.ForumTopic.ForumID, 123456)
	assert.Equal(t, f.ForumTopic.Server.ServerID, "wlVr3Ggl")
}
