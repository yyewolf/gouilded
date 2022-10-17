package guilded

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type docsTest struct {
	Docs *Docs
}

func TestDocs_UnmarshallJSON(t *testing.T) {
	d := docsTest{
		Docs: &Docs{},
	}

	data := []byte(`{
		"id": 0,
		"serverId": "wlVr3Ggl",
		"channelId": "00000000-0000-0000-0000-000000000000",
		"title": "HOW-TO: Smoke These Meats with Sweet Baby Ray's",
		"content": "Spicy jalapeno bacon ipsum dolor amet sirloin ground round short loin, meatball brisket capicola tri-tip ham pork belly biltong corned beef chuck. Chicken ham brisket shank rump buffalo t-bone. Short loin sausage buffalo porchetta pork belly rump tri-tip frankfurter tail pork chop cow sirloin. Pancetta porchetta tail ball tip chislic beef ribs. Buffalo andouille leberkas jerky. Fatback shankle andouille beef. Cow kielbasa buffalo pork loin chislic meatloaf short loin rump meatball prosciutto.",
		"createdAt": "2021-06-15T20:15:00.706Z",
		"createdBy": "Ann6LewA",
		"updatedAt": "2021-07-15T22:20:00.706Z",
		"updatedBy": "Ann6LewA"
	  }`)

	if err := d.Docs.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, d.Docs.DocsID, 0)
	assert.Equal(t, d.Docs.Server.ServerID, "wlVr3Ggl")
	assert.Equal(t, d.Docs.Channel.ChannelID, "00000000-0000-0000-0000-000000000000")
}
