package model

// Media represents simple anime/manga model.
type Media struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Image      string  `json:"image"`
	Score      float64 `json:"score"`
	Voter      int     `json:"voter"`
	Rank       int     `json:"rank"`
	Popularity int     `json:"popularity"`
	Member     int     `json:"member"`
	Favorite   int     `json:"favorite"`
	Type       int     `json:"type"`
	Status     int     `json:"status"`
}

// AnimeQuery represents anime search query model.
type AnimeQuery struct {
	// Basic search.
	Title  string
	Score  int
	Type   int
	Status int
	Rating int
	Source int
	Year   int
	Season string
	Limit  int
	Page   int
	Order  string

	// Advanced search.
	StartYear     int
	EndYear       int
	StartEpisode  int
	EndEpisode    int
	StartDuration int // Minutes.
	EndDuration   int // Minutes.
	Producer      int
	Genre         []int
}

// MangaQuery represents manga search query model.
type MangaQuery struct {
	// Basic search.
	Title  string
	Score  int
	Type   int
	Status int
	Year   int
	Limit  int
	Page   int
	Order  string

	// Advanced search.
	StartYear    int
	EndYear      int
	StartVolume  int
	EndVolume    int
	StartChapter int
	EndChapter   int
	Magazine     int
	Genre        []int
}

// EntryQuery represents character/people search query model.
type EntryQuery struct {
	Name  string
	Limit int
	Page  int
	Order string
}
