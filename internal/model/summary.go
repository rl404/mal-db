package model

// Total represents total main type count model.
type Total struct {
	Anime     int `json:"anime"`
	Manga     int `json:"manga"`
	Character int `json:"character"`
	People    int `json:"people"`
}

// YearSummary represents anime & manga score summary.
type YearSummary struct {
	Anime YearSummaryDetail `json:"anime"`
	Manga YearSummaryDetail `json:"manga"`
}

// YearSummaryDetail represents detail score summary.
type YearSummaryDetail struct {
	Year     int     `json:"year"`
	Count    int     `json:"count"`
	AvgScore float64 `json:"avgScore"`
	MinScore float64 `json:"minScore"`
	MaxScore float64 `json:"maxScore"`
}
