package guilded

import (
	"encoding/json"
	"time"
)

type CalendarEvent struct {
	// The ID of the calendar event (min 1)
	CalendarID int `json:"id"`

	// The server
	Server *Server `json:"-"`

	// The channel
	Channel *Channel `json:"-"`

	// The name of the event (min length 1; max length 60)
	Name string `json:"name"`

	// The description of the event (min length 1; max length 8000)
	Description string `json:"description"`

	// The location of the event (min length 1; max length 8000)
	Location string `json:"location"`

	// A URL to associate with the event
	Url string `json:"url"`

	// The color of the event when viewing in the calendar (min 0; max 16777215)
	Color int `json:"color"`

	// The number of RSVPs to allow before waitlisting RSVPs (min 1)
	RSVPLimit int `json:"rsvpLimit"`

	// The ISO 8601 timestamp that the event starts at
	StartsAt time.Time `json:"startsAt"`

	// The duration of the event in minutes (min 1)
	Duration int `json:"duration"`

	IsPrivate bool `json:"isPrivate"`

	Mentions *Mention `json:"mentions"`

	// The ISO 8601 timestamp that the event was created at
	CreatedAt time.Time `json:"createdAt"`

	// The user who created this event
	CreatedBy *User `json:"-"`

	Cancellation *Cancellation
}

type Cancellation struct {
	// The description of event cancellation (min length 1; max length 140)
	Description string `json:"description"`

	// The user who created this event cancellation
	CreatedBy *User `json:"-"`
}

type CalendarEventRsvpStatus string

const (
	CalendarEventRsvpStatusGoing        CalendarEventRsvpStatus = "going"
	CalendarEventRsvpStatusMaybe        CalendarEventRsvpStatus = "maybe"
	CalendarEventRsvpStatusDeclined     CalendarEventRsvpStatus = "declined"
	CalendarEventRsvpStatusInvited      CalendarEventRsvpStatus = "invited"
	CalendarEventRsvpStatusWaitListed   CalendarEventRsvpStatus = "waitlisted"
	CalendarEventRsvpStatusNotResponded CalendarEventRsvpStatus = "not responded"
)

type CalendarEventRsvp struct {
	// The calendar event (min 1)
	CalendarEvent *CalendarEvent `json:"-"`

	// The server
	Server *Server `json:"-"`

	// The channel
	Channel *Channel `json:"-"`

	// The user
	User *User `json:"-"`

	// The status of the RSVP
	Status CalendarEventRsvpStatus `json:"status"`

	// The user who created this RSVP
	CreatedBy *User `json:"-"`

	// The ISO 8601 timestamp that the RSVP was created at
	CreatedAt time.Time `json:"createdAt"`

	// The user who updated this RSVP
	UpdatedBy *User `json:"-"`

	// The ISO 8601 timestamp that the RSVP was updated at, if relevant
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *CalendarEvent) UnmarshalJSON(data []byte) error {
	type Alias CalendarEvent
	aux := &struct {
		*Alias
		ServerID  string `json:"serverId"`
		ChannelID string `json:"channelId"`
		CreatedBy string `json:"createdBy"`
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.Server = &Server{ServerID: aux.ServerID}
	c.Channel = &Channel{ChannelID: aux.ChannelID}
	c.CreatedBy = &User{UserID: aux.CreatedBy}
	return nil
}

func (c *Cancellation) UnmarshalJSON(data []byte) error {
	type Alias Cancellation
	aux := &struct {
		*Alias
		CreatedBy string `json:"createdBy"`
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.CreatedBy = &User{UserID: aux.CreatedBy}
	return nil
}

func (c *CalendarEventRsvp) UnmarshalJSON(data []byte) error {
	type Alias CalendarEventRsvp
	aux := &struct {
		*Alias
		CalendarEventId int    `json:"calendarEventId"`
		ServerID        string `json:"serverId"`
		ChannelID       string `json:"channelId"`
		UserID          string `json:"userId"`
		CreatedBy       string `json:"createdBy"`
		UpdatedBy       string `json:"updatedBy"`
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	c.CalendarEvent = &CalendarEvent{CalendarID: aux.CalendarEventId}
	c.Server = &Server{ServerID: aux.ServerID}
	c.Channel = &Channel{ChannelID: aux.ChannelID}
	c.User = &User{UserID: aux.UserID}
	c.CreatedBy = &User{UserID: aux.CreatedBy}
	c.UpdatedBy = &User{UserID: aux.UpdatedBy}
	return nil
}
