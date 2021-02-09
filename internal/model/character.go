package model

// Character represents character model retrieved from db.
type Character struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Nicknames    []string `json:"nicknames"`
	JapaneseName string   `json:"japaneseName"`
	Image        string   `json:"image"`
	Favorite     int      `json:"favorite"`
	About        string   `json:"about"`
}
