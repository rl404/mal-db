package model

// Manga represent main manga model.
type Manga struct {
	ID                int              `json:"id"`
	Title             string           `json:"title"`
	AlternativeTitles AlternativeTitle `json:"alternativeTitles"`
	Image             string           `json:"image"`
	Synopsis          string           `json:"synopsis"`
	Score             float64          `json:"score"`
	Voter             int              `json:"voter"`
	Rank              int              `json:"rank"`
	Popularity        int              `json:"popularity"`
	Member            int              `json:"member"`
	Favorite          int              `json:"favorite"`
	Type              string           `json:"type"`
	Volume            int              `json:"volume"`
	Chapter           int              `json:"chapter"`
	Status            string           `json:"status"`
	PublishingDate    StartEndDate     `json:"publishingDate"`
	Genres            []Item           `json:"genres"`
	Authors           []Item           `json:"authors"`
	Serializations    []Item           `json:"serializations"`
	Related           Related          `json:"related"`
}

// MangaItem represents simpler manga model for magazine.
type MangaItem struct {
	ID             int      `json:"id"`
	Title          string   `json:"title"`
	Image          string   `json:"image"`
	Volume         int      `json:"volume"`
	Type           string   `json:"type"`
	Member         int      `json:"member"`
	Score          float64  `json:"score"`
	StartDate      Date     `json:"startDate"`
	Synopsis       string   `json:"synopsis"`
	Authors        []Item   `json:"authors"`
	Serializations []string `json:"serializations"`
	Genres         []Item   `json:"genres"`
}

// TopManga represents model for top manga list.
type TopManga struct {
	Rank      int     `json:"rank"`
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	Type      string  `json:"type"`
	Volume    int     `json:"volume"`
	StartDate Date    `json:"startDate"`
	EndDate   Date    `json:"endDate"`
	Member    int     `json:"member"`
	Score     float64 `json:"score"`
}
