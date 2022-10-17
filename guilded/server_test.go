package guilded

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type serverTest struct {
	Server *Server `json:"server,omitempty"`
}

func TestServer_UnmarshallJSON(t *testing.T) {
	s := serverTest{
		Server: &Server{},
	}

	data := []byte(`{
		"id": "wlVr3Ggl",
		"type": "community",
		"name": "Guilded",
		"url": "Guilded-Official",
		"about": "The Official Guilded Team! For devs, friends, and fans alike!",
		"ownerId": "Ann6LewA",
		"createdAt": "2018-10-05T20:15:00.706Z",
		"isVerified": true,
		"timezone": "America/Los Angeles (PST/PDT)"
	  }`)

	if err := s.Server.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, s.Server.ServerID, "wlVr3Ggl")
	assert.Equal(t, s.Server.CreatedBy.UserID, "Ann6LewA")
}
