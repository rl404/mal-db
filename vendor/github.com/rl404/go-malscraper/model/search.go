package model

import "time"

// Query represents anime & manga search query model.
type Query struct {
	Title        string
	Page         int
	Type         int
	Score        int
	Status       int
	ProducerID   int // for anime only
	MagazineID   int // for manga only
	Rating       int // for anime only
	StartDate    time.Time
	EndDate      time.Time
	ExcludeGenre bool
	GenreIDs     []int
	FirstLetter  string
}

// ClubQuery represents club search query model.
type ClubQuery struct {
	Name     string
	Page     int
	Category int
	Sort     int
}

// UserQuery represents user search query model.
type UserQuery struct {
	Username string
	Page     int
	Location string
	MinAge   int
	MaxAge   int
	Gender   int
}

// AnimeSearch represents anime search result model.
type AnimeSearch struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	Summary   string  `json:"summary"`
	Type      string  `json:"type"`
	Episode   int     `json:"episode"`
	Score     float64 `json:"score"`
	StartDate Date    `json:"startDate"`
	EndDate   Date    `json:"endDate"`
	Member    int     `json:"member"`
	Rated     string  `json:"rated"`
}

// MangaSearch represents manga search result model.
type MangaSearch struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	Summary   string  `json:"summary"`
	Type      string  `json:"type"`
	Volume    int     `json:"volume"`
	Chapter   int     `json:"chapter"`
	Score     float64 `json:"score"`
	StartDate Date    `json:"startDate"`
	EndDate   Date    `json:"endDate"`
	Member    int     `json:"member"`
}

// CharacterSearch represents character search result model.
type CharacterSearch struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Image    string `json:"image"`
}

// PeopleSearch represents people search result model.
type PeopleSearch struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Image    string `json:"image"`
}

// ClubSearch represents club search result model.
type ClubSearch struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	Summary string `json:"summary"`
	Creator string `json:"creator"`
	Member  int    `json:"member"`
}

// UserSearch represents user search result model.
type UserSearch struct {
	Username   string     `json:"username"`
	Image      string     `json:"image"`
	LastOnline *time.Time `json:"lastOnline"`
}
