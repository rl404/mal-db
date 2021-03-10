package model

// People represents people model retrieved from db.
type People struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	Image            string   `json:"image"`
	GivenName        string   `json:"givenName"`
	FamilyName       string   `json:"familyName"`
	AlternativeNames []string `json:"alternativeNames"`
	Birthday         Date     `json:"birthday"`
	Website          string   `json:"website"`
	Favorite         int      `json:"favorite"`
	More             string   `json:"more"`
}

// VoiceActor represents voice actor model with their anime and character role.
type VoiceActor struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Role  string `json:"role"`
	Anime []Role `json:"anime"`
}
