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
		"id": "b943384a-d951-4323-8e26-8e3e6b7c431a",
		"content": {
		  "object": "value",
		  "document": {
			"object": "document",
			"data": {},
			"nodes": [
			  {
				"object": "block",
				"type": "paragraph",
				"data": {},
				"nodes": [
				  {
					"object": "text",
					"leaves": [
					  {
						"object": "leaf",
						"text": "cool message with an ",
						"marks": []
					  }
					]
				  },
				  {
					"object": "inline",
					"type": "reaction",
					"data": {
					  "reaction": {
						"id": 90001815,
						"customReactionId": 90001815
					  }
					},
					"nodes": [
					  {
						"object": "text",
						"leaves": [
						  {
							"object": "leaf",
							"text": ":tada:",
							"marks": []
						  }
						]
					  }
					]
				  },
				  {
					"object": "text",
					"leaves": [
					  {
						"object": "leaf",
						"text": " emoji, some ",
						"marks": []
					  },
					  {
						"object": "leaf",
						"text": "bold formatting",
						"marks": [
						  {
							"object": "mark",
							"type": "italic",
							"data": {}
						  }
						]
					  },
					  {
						"object": "leaf",
						"text": ", maybe even a ",
						"marks": []
					  }
					]
				  },
				  {
					"object": "inline",
					"type": "reaction",
					"data": {
					  "reaction": {
						"id": 294637,
						"customReactionId": 294637
					  }
					},
					"nodes": [
					  {
						"object": "text",
						"leaves": [
						  {
							"object": "leaf",
							"text": ":thinkingwithblobs:",
							"marks": []
						  }
						]
					  }
					]
				  },
				  {
					"object": "text",
					"leaves": [
					  {
						"object": "leaf",
						"text": " custom emoji?",
						"marks": []
					  }
					]
				  }
				]
			  }
			]
		  }
		}
	  }`)

	if err := json.Unmarshal(data, v.Message); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, v.Message.ID, "b943384a-d951-4323-8e26-8e3e6b7c431a")
}
