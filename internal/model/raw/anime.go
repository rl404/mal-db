package raw

import "time"

// Anime is model for anime table.
type Anime struct {
	ID            int       `gorm:"primary_key"`
	Title         string    `gorm:"type:varchar"`
	TitleEnglish  string    `gorm:"type:varchar"`
	TitleJapanese string    `gorm:"type:varchar"`
	TitleSynonym  string    `gorm:"type:varchar"`
	ImageURL      string    `gorm:"type:varchar"`
	VideoURL      string    `gorm:"type:varchar"`
	Synopsis      string    `gorm:"type:varchar"`
	Score         float64   `gorm:"type:numeric"`
	Voter         int       `gorm:"type:int"`
	Rank          int       `gorm:"type:int"`
	Popularity    int       `gorm:"type:int"`
	Member        int       `gorm:"type:int"`
	Favorite      int       `gorm:"type:int"`
	AnimeTypeID   int       `gorm:"type:int"`
	Episode       int       `gorm:"type:int"`
	AnimeStatusID int       `gorm:"type:int"`
	StartYear     int       `gorm:"type:int"`
	StartMonth    int       `gorm:"type:int"`
	StartDay      int       `gorm:"type:int"`
	EndYear       int       `gorm:"type:int"`
	EndMonth      int       `gorm:"type:int"`
	EndDay        int       `gorm:"type:int"`
	AiringDay     string    `gorm:"type:varchar"`
	AiringTime    string    `gorm:"type:varchar"`
	Premiered     string    `gorm:"type:varchar"`
	AnimeSourceID int       `gorm:"type:int"`
	Duration      int       `gorm:"type:int"`
	AnimeRatingID int       `gorm:"type:int"`
	CreatedAt     time.Time `gorm:"type:timestamp"`
	UpdatedAt     time.Time `gorm:"type:timestamp"`
}

// TableName to get table name.
func (a Anime) TableName() string {
	return "anime"
}
