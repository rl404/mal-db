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

// AlternativeTitle represents anime & manga alternative titles.
type AlternativeTitle struct {
	English  string `json:"english"`
	Japanese string `json:"japanese"`
	Synonym  string `json:"synonym"`
}

// Related represents related anime & manga model.
type Related struct {
	Sequel      []Source `json:"sequel"`
	Prequel     []Source `json:"prequel"`
	AltSetting  []Source `json:"alternativeSetting"`
	AltVersion  []Source `json:"alternativeVersion"`
	SideStory   []Source `json:"sideStory"`
	Summary     []Source `json:"summary"`
	FullStory   []Source `json:"fullStory"`
	ParentStory []Source `json:"parentStory"`
	SpinOff     []Source `json:"spinOff"`
	Adaptation  []Source `json:"adaptation"`
	Character   []Source `json:"character"`
	Other       []Source `json:"other"`
}

// Role represents simple model of character/va's role.
type Role struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Role  string `json:"role"`
}

// Stats represents anime & manga stats data.
type Stats struct {
	Summary Summary `json:"summary"`
	Score   Score   `json:"score"`
}

// Summary represents anime & manga status count.
type Summary struct {
	Current   int `json:"current"`
	Completed int `json:"completed"`
	OnHold    int `json:"onHold"`
	Dropped   int `json:"dropped"`
	Planned   int `json:"planned"`
}

// Score represents anime & manga stats for each score.
type Score struct {
	Score1  ScoreDetail `json:"1"`
	Score2  ScoreDetail `json:"2"`
	Score3  ScoreDetail `json:"3"`
	Score4  ScoreDetail `json:"4"`
	Score5  ScoreDetail `json:"5"`
	Score6  ScoreDetail `json:"6"`
	Score7  ScoreDetail `json:"7"`
	Score8  ScoreDetail `json:"8"`
	Score9  ScoreDetail `json:"9"`
	Score10 ScoreDetail `json:"10"`
}

// ScoreDetail represents anime & manga each score count.
type ScoreDetail struct {
	Vote    int     `json:"vote"`
	Percent float64 `json:"percent"`
}

// Entry represents simple entry model.
type Entry struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// StatsHistory represents entry stats history.
type StatsHistory struct {
	Year       int     `json:"year"`
	Month      int     `json:"month"`
	Score      float64 `json:"score"`
	Voter      int     `json:"voter"`
	Rank       int     `json:"rank"`
	Popularity int     `json:"popularity"`
	Member     int     `json:"member"`
	Favorite   int     `json:"favorite"`
}
