package model

// People represents main people model.
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

// PeopleCharacter represents people's anime and character role.
type PeopleCharacter struct {
	Anime     Role `json:"anime"`
	Character Role `json:"character"`
}

// TopPeople represents model for top people list.
type TopPeople struct {
	Rank         int    `json:"rank"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	JapaneseName string `json:"japaneseName"`
	Image        string `json:"image"`
	Birthday     Date   `json:"birthday"`
	Favorite     int    `json:"favorite"`
}
