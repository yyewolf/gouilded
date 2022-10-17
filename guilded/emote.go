package guilded

type Emote struct {
	// The ID of the emote
	EmojiID int `json:"id"`

	// The name of the emote
	Name string `json:"name"`

	// The URL of the emote image
	Url string `json:"url"`
}
