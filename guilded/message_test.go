package guilded

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type messageTest struct {
	Message *Message `json:"message,omitempty"`
}

func TestMessage_UnmarshalJSON(t *testing.T) {
	v := messageTest{
		Message: &Message{},
	}

	data := []byte(`{
		"id": "00000000-0000-0000-0000-000000000000",
		"type": "default",
		"serverId": "wlVr3Ggl",
		"channelId": "00000000-0000-0000-0000-000000000000",
		"content": "Hello **world**!",
		"embeds": [
		  {
			"title": "embed title",
			"description": "embeds support a **different** __subset__ *of* markdown than other markdown fields. <@Ann6LewA>\n\n [links](https://www.guilded.gg) Stime!! ttyl",
			"url": "https://www.guilded.gg",
			"color": 6118369,
			"timestamp": "2022-04-12T22:14:36.737Z",
			"footer": {
			  "icon_url": "https://www.guilded.gg/asset/Logos/logomark/Color/Guilded_Logomark_Color.png",
			  "text": "footer text"
			},
			"thumbnail": {
			  "url": "https://www.guilded.gg/asset/Logos/logomark/Color/Guilded_Logomark_Color.png"
			},
			"image": {
			  "url": "https://www.guilded.gg/asset/Logos/logomark_wordmark/Color/Guilded_Logomark_Wordmark_Color.png"
			},
			"author": {
			  "name": "Gil",
			  "url": "https://www.guilded.gg",
			  "icon_url": "https://www.guilded.gg/asset/Default/Gil-md.png"
			},
			"fields": [
			  {
				"name": "hello",
				"value": "these are fields"
			  },
			  {
				"name": "~~help i have been crossed out~~",
				"value": "~~oh noes~~",
				"inline": true
			  },
			  {
				"name": "another inline",
				"value": "field",
				"inline": true
			  }
			]
		  }
		],
		"createdAt": "2021-06-15T20:15:00.706Z",
		"createdBy": "Ann6LewA"
	  }`)

	if err := json.Unmarshal(data, v.Message); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, v.Message.ID, "00000000-0000-0000-0000-000000000000")
}
