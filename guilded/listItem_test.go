package guilded

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type listItemTest struct {
	ListItem *ListItem
}

func TestListItem_UnmarshallJSON(t *testing.T) {
	l := listItemTest{
		ListItem: &ListItem{},
	}

	data := []byte(`{
		"id": "00000000-0000-0000-0000-000000000000",
		"serverId": "wlVr3Ggl",
		"channelId": "00000000-0000-0000-0000-000000000000",
		"message": "Remember to say hello **world**!",
		"createdAt": "2021-06-15T20:15:00.706Z",
		"createdBy": "Ann6LewA",
		"note": {
		  "createdAt": "2021-06-15T20:15:00.706Z",
		  "createdBy": "Ann6LewA",
		  "content": "Duly noted"
		}
	  }`)

	if err := l.ListItem.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, l.ListItem.Channel.ChannelID, "00000000-0000-0000-0000-000000000000")
	assert.Equal(t, l.ListItem.Server.ServerID, "wlVr3Ggl")
}

type noteTest struct {
	Note *Note
}

func TestNote_UnmarshallJSON(t *testing.T) {
	n := noteTest{
		Note: &Note{},
	}

	data := []byte(`{
		"createdAt": "2021-06-15T20:15:00.706Z",
		"createdBy": "Ann6LewA",
		"content": "Duly noted"
	  }`)

	if err := n.Note.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, n.Note.CreatedBy.UserID, "Ann6LewA")
	assert.Equal(t, n.Note.Content, "Duly noted")
}
