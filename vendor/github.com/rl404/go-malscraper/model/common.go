package model

// Item represents common struct containing id and name.
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ItemCount represents common struct containing id, name and count.
type ItemCount struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Date represents splitted date.
//
// Normal date usually has day, month, and year. But,
// in MyAnimeList, there are some dates that don't have them all 3.
// For example, (future) anime airing date which sometimes contains
// only year. So, it can't be parsed to time.Time. That's why this date
// has its own data type.
type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// Source represents simple anime/manga model.
type Source struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Type  string `json:"type"`
}

// Role represents simple model of character/va's role.
type Role struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Role  string `json:"role"`
}

// AlternativeTitle represents alternative english, synonym, and japanese title of anime/manga.
type AlternativeTitle struct {
	English  string `json:"english"`
	Synonym  string `json:"synonym"`
	Japanese string `json:"japanese"`
}

// StartEndDate represents anime/manga start and end of airing/publishing date.
type StartEndDate struct {
	Start Date `json:"start"`
	End   Date `json:"end"`
}

// Related represents related anime & manga model.
type Related struct {
	Sequel      []RelatedItem `json:"sequel"`
	Prequel     []RelatedItem `json:"prequel"`
	AltSetting  []RelatedItem `json:"alternativeSetting"`
	AltVersion  []RelatedItem `json:"alternativeVersion"`
	SideStory   []RelatedItem `json:"sideStory"`
	Summary     []RelatedItem `json:"summary"`
	FullStory   []RelatedItem `json:"fullStory"`
	ParentStory []RelatedItem `json:"parentStory"`
	SpinOff     []RelatedItem `json:"spinOff"`
	Adaptation  []RelatedItem `json:"adaptation"`
	Character   []RelatedItem `json:"character"`
	Other       []RelatedItem `json:"other"`
}

// RelatedItem represents related anime/manga model.
type RelatedItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

// Stats represents anime/manga stats model.
type Stats struct {
	Summary StatsSummary `json:"summary"`
	Score   StatsScore   `json:"score"`
}

// StatsSummary represents stats summary count.
type StatsSummary struct {
	Current   int `json:"current"`
	Completed int `json:"completed"`
	OnHold    int `json:"onHold"`
	Dropped   int `json:"dropped"`
	Planned   int `json:"planned"`
	Total     int `json:"total"`
}

// StatsScore represents stats each score.
type StatsScore struct {
	Score10 StatsScoreItem `json:"10"`
	Score9  StatsScoreItem `json:"9"`
	Score8  StatsScoreItem `json:"8"`
	Score7  StatsScoreItem `json:"7"`
	Score6  StatsScoreItem `json:"6"`
	Score5  StatsScoreItem `json:"5"`
	Score4  StatsScoreItem `json:"4"`
	Score3  StatsScoreItem `json:"3"`
	Score2  StatsScoreItem `json:"2"`
	Score1  StatsScoreItem `json:"1"`
}

// StatsScoreItem represents detail score model and its count.
type StatsScoreItem struct {
	Vote    int     `json:"vote"`
	Percent float64 `json:"percent"`
}
