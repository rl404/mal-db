package model

// Manga represents manga model retrieved from db.
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
	Type              int              `json:"type"`
	Volume            int              `json:"volume"`
	Chapter           int              `json:"chapter"`
	Status            int              `json:"status"`
	Publishing        Publishing       `json:"publishing"`
	Genres            []Item           `json:"genres"`
	Authors           []Role           `json:"authors"`
	Serializations    []Item           `json:"serializations"`
	Related           Related          `json:"related"`
}

// Publishing represents manga publishing dates.
type Publishing struct {
	Start Date `json:"start"`
	End   Date `json:"end"`
}
