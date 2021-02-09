package raw

import "time"

// StatsHistory is model for history of entries stats.
type StatsHistory struct {
	ID         int       `gorm:"primary_key"`
	MediaID    int       `gorm:"primary_key"`
	Type       string    `gorm:"primary_key;type:varchar"` // anime/manga
	Score      float64   `gorm:"type:numeric"`
	Voter      int       `gorm:"type:int"`
	Rank       int       `gorm:"type:int"`
	Popularity int       `gorm:"type:int"`
	Member     int       `gorm:"type:int"`
	Favorite   int       `gorm:"type:int"`
	CreatedAt  time.Time `gorm:"type:timestamp"`
}
