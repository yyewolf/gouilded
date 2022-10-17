package guilded

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type emoteTest struct {
	Emote *Emote
}

func TestEmote_UnmarshallJSON(t *testing.T) {
	v := emoteTest{
		Emote: &Emote{},
	}

	data := []byte(`{
		"id": 90000000,
		"name": "grinning",
		"url": "https://img.guildedcdn.com/asset/Emojis/grinning.webp"
	  }`)

	if err := json.Unmarshal(data, v.Emote); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, v.Emote.EmojiID, 90000000)
	assert.Equal(t, v.Emote.Name, "grinning")
}
