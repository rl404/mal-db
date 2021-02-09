package model

import "time"

// Anime represent main anime model.
type Anime struct {
	ID                int              `json:"id"`
	Title             string           `json:"title"`
	AlternativeTitles AlternativeTitle `json:"alternativeTitles"`
	Image             string           `json:"image"`
	Video             string           `json:"video"`
	Synopsis          string           `json:"synopsis"`
	Score             float64          `json:"score"`
	Voter             int              `json:"voter"`
	Rank              int              `json:"rank"`
	Popularity        int              `json:"popularity"`
	Member            int              `json:"member"`
	Favorite          int              `json:"favorite"`
	Type              string           `json:"type"`
	Episode           int              `json:"episode"`
	Status            string           `json:"status"`
	AiringDate        StartEndDate     `json:"airingDate"`
	Premiered         string           `json:"premiered"`
	Broadcast         string           `json:"broadcast"`
	Source            string           `json:"source"`
	Duration          string           `json:"duration"`
	Rating            string           `json:"rating"`
	Producers         []Item           `json:"producers"`
	Licensors         []Item           `json:"licensors"`
	Studios           []Item           `json:"studios"`
	Genres            []Item           `json:"genres"`
	Related           Related          `json:"related"`
	Song              Song             `json:"song"`
}

// Song represents list of opening and ending anime songs.
type Song struct {
	Opening []string `json:"opening"`
	Ending  []string `json:"ending"`
}

// Episode represents anime episode model.
type Episode struct {
	Episode       int        `json:"episode"`
	Title         string     `json:"title"`
	JapaneseTitle string     `json:"japaneseTitle"`
	AiredDate     *time.Time `json:"airedDate"`
	Link          string     `json:"link"`
	Tag           string     `json:"tag"`
}

// Video represents anime video model.
type Video struct {
	Episodes   []VideoEpisode `json:"episodes"`
	Promotions []VideoPromo   `json:"promotions"`
}

// VideoEpisode represents anime episode model.
type VideoEpisode struct {
	Episode int    `json:"episode"`
	Title   string `json:"title"`
	Link    string `json:"link"`
}

// VideoPromo represents promotion video model.
type VideoPromo struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

// AnimeItem represents simpler anime model for producer/seasonal anime.
type AnimeItem struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Image     string   `json:"image"`
	Source    string   `json:"source"`
	Episode   int      `json:"episode"`
	Type      string   `json:"type"`
	Member    int      `json:"member"`
	Score     float64  `json:"score"`
	StartDate Date     `json:"startDate"`
	Synopsis  string   `json:"synopsis"`
	Genres    []Item   `json:"genres"`
	Producers []Item   `json:"producers"`
	Licensors []string `json:"licensors"`
}

// TopAnime represents model for top anime list.
type TopAnime struct {
	Rank      int     `json:"rank"`
	Title     string  `json:"title"`
	ID        int     `json:"id"`
	Image     string  `json:"image"`
	Type      string  `json:"type"`
	Episode   int     `json:"episode"`
	StartDate Date    `json:"startDate"`
	EndDate   Date    `json:"endDate"`
	Member    int     `json:"member"`
	Score     float64 `json:"score"`
}
