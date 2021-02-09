package model

// Character represents main character model.
type Character struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Nickname     string `json:"nickname"`
	JapaneseName string `json:"japaneseName"`
	Image        string `json:"image"`
	Favorite     int    `json:"favorite"`
	About        string `json:"about"`
}

// CharacterItem represents character role model in anime/manga.
type CharacterItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Role        string `json:"role"`
	VoiceActors []Role `json:"voiceActors"`
}

// TopCharacter represents model for top character list.
type TopCharacter struct {
	Rank         int    `json:"rank"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	JapaneseName string `json:"japaneseName"`
	Image        string `json:"image"`
	Favorite     int    `json:"favorite"`
}
