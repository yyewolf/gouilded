package guilded

import (
	"encoding/json"
	"time"
)

type MessageEmbed struct {
	// Title is the title of the embed
	Title string `json:"title"`

	// Description is the description of the embed
	Description string `json:"description"`

	// URL is used to linkify the title
	URL string `json:"url"`

	// Color is the color of the embed
	Color int `json:"color"`

	// Footer is the footer of the embed
	Footer *MessageEmbedFooter `json:"footer"`

	// Timestamp is the timestamp of the embed
	Timestamp time.Time `json:"timestamp"`

	// Thumbnail is the thumbnail of the embed
	Thumbnail *MessageEmbedThumbnail `json:"thumbnail"`

	// Image is the image of the embed
	Image *MessageEmbedImage `json:"image"`

	// Author is the author of the embed
	Author *MessageEmbedAuthor `json:"author"`

	// Fields is the list of fields in the embed
	Fields []*MessageEmbedField `json:"fields"`
}

type MessageEmbedFooter struct {
	// IconURL is the URL of the footer icon
	IconURL string `json:"icon_url"`

	// Text is the text of the footer
	Text string `json:"text"`
}

type MessageEmbedThumbnail struct {
	// URL is the URL of the thumbnail
	URL string `json:"url"`
}

type MessageEmbedImage struct {
	// URL is the URL of the image
	URL string `json:"url"`
}

type MessageEmbedAuthor struct {
	// Name is the name of the author
	Name string `json:"name"`

	// URL is the URL of the author
	URL string `json:"url"`

	// IconURL is the URL of the author icon
	IconURL string `json:"icon_url"`
}

type MessageEmbedField struct {
	// Name is the name of the field
	Name string `json:"name"`

	// Value is the value of the field
	Value string `json:"value"`

	// Inline is whether the field is inline
	Inline bool `json:"inline"`
}

func (e *MessageEmbed) MarshalJSON() ([]byte, error) {
	type Alias MessageEmbed
	aux := &struct {
		*Alias
		Timestamp string `json:"timestamp"`
	}{
		Alias:     (*Alias)(e),
		Timestamp: e.Timestamp.Format(time.RFC3339),
	}

	return json.Marshal(aux)
}
