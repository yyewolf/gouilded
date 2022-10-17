package guilded

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type calendarEventTest struct {
	CalendarEvent *CalendarEvent
}

func TestCalendarEvent_UnmarshallJSON(t *testing.T) {
	c := calendarEventTest{
		CalendarEvent: &CalendarEvent{},
	}

	data := []byte(`{
		"id": 1,
		"serverId": "wlVr3Ggl",
		"channelId": "00000000-0000-0000-0000-000000000000",
		"name": "Surprise LAN party for my wife 🤫",
		"description": "**Don't say anything to her!** She's gonna love playing Call of Duty all night",
		"location": "My house!",
		"url": "https://www.surprisepartygame.com/",
		"duration": 60,
		"color": 16106496,
		"startsAt": "2022-06-16T00:00:00.000Z",
		"createdAt": "2021-06-15T20:15:00.706Z",
		"createdBy": "Ann6LewA"
	  }`)

	if err := c.CalendarEvent.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, c.CalendarEvent.CalendarID, 1)
	assert.Equal(t, c.CalendarEvent.Server.ServerID, "wlVr3Ggl")
	assert.Equal(t, c.CalendarEvent.Channel.ChannelID, "00000000-0000-0000-0000-000000000000")
}

type cancellationTest struct {
	Cancellation *Cancellation
}

func TestCancellation_UnmarshallJSON(t *testing.T) {
	c := cancellationTest{
		Cancellation: &Cancellation{},
	}

	data := []byte(`{
		"description": "this is a test"
	}`)

	if err := c.Cancellation.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, c.Cancellation.Description, "this is a test")
}

type calendarEventRsvpTest struct {
	CalendarEventRsvp *CalendarEventRsvp
}

func TestCalendarEventRsvp_UnmarshallJSON(t *testing.T) {
	c := calendarEventRsvpTest{
		CalendarEventRsvp: &CalendarEventRsvp{},
	}

	data := []byte(`{
		"calendarEventId": 1,
		"channelId": "00000000-0000-0000-0000-000000000000",
		"serverId": "wlVr3Ggl",
		"userId": "Ann6LewA",
		"status": "going",
		"createdAt": "2021-06-15T20:15:00.706Z",
		"createdBy": "Ann6LewA"
	  }`)

	if err := c.CalendarEventRsvp.UnmarshalJSON(data); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, c.CalendarEventRsvp.CalendarEvent.CalendarID, 1)
	assert.Equal(t, c.CalendarEventRsvp.Server.ServerID, "wlVr3Ggl")
	assert.Equal(t, c.CalendarEventRsvp.User.UserID, "Ann6LewA")
}
