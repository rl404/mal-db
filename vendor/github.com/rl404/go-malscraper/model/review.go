package model

import "time"

// Review represents main review model.
type Review struct {
	ID       int         `json:"id"`
	Username string      `json:"username"`
	Image    string      `json:"image"`
	Source   Source      `json:"source"`
	Helpful  int         `json:"helpful"`
	Date     time.Time   `json:"date"`
	Episode  string      `json:"episode"` // for anime
	Chapter  string      `json:"chapter"` // for manga
	Score    ReviewScore `json:"score"`
	Review   string      `json:"review"`
}

// ReviewScore represents review's detail score.
type ReviewScore struct {
	Overall   int `json:"overall"`   // for both
	Story     int `json:"story"`     // for both
	Art       int `json:"art"`       // for both
	Sound     int `json:"sound"`     // for anime only
	Character int `json:"character"` // for both
	Enjoyment int `json:"enjoyment"` // for both
}
