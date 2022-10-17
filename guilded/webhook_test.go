package guilded

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type webhookTest struct {
	Webhook *Webhook
}

func TestWebhook_UnmarshallJSON(t *testing.T) {
	w := webhookTest{
		Webhook: &Webhook{},
	}

	data := []byte(`{
		"id": "00000000-0000-0000-0000-000000000000",
		"serverId": "wlVr3Ggl",
		"channelId": "00000000-0000-0000-0000-000000000000",
		"name": "E-102 Gamma",
		"createdAt": "2021-06-15T20:15:00.706Z",
		"createdBy": "Ann6LewA"
	  }`)

	if err := w.Webhook.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, w.Webhook.WebhookID, "00000000-0000-0000-0000-000000000000")
	assert.Equal(t, w.Webhook.Name, "E-102 Gamma")
}
