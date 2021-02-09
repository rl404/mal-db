package raw

import "time"

// Manga is model for manga table.
type Manga struct {
	ID            int       `gorm:"primary_key"`
	Title         string    `gorm:"type:varchar"`
	TitleEnglish  string    `gorm:"type:varchar"`
	TitleJapanese string    `gorm:"type:varchar"`
	TitleSynonym  string    `gorm:"type:varchar"`
	ImageURL      string    `gorm:"type:varchar"`
	Synopsis      string    `gorm:"type:varchar"`
	Score         float64   `gorm:"type:numeric"`
	Voter         int       `gorm:"type:int"`
	Rank          int       `gorm:"type:int"`
	Popularity    int       `gorm:"type:int"`
	Member        int       `gorm:"type:int"`
	Favorite      int       `gorm:"type:int"`
	MangaTypeID   int       `gorm:"type:int"`
	Volume        int       `gorm:"type:int"`
	Chapter       int       `gorm:"type:int"`
	MangaStatusID int       `gorm:"type:int"`
	StartYear     int       `gorm:"type:int"`
	StartMonth    int       `gorm:"type:int"`
	StartDay      int       `gorm:"type:int"`
	EndYear       int       `gorm:"type:int"`
	EndMonth      int       `gorm:"type:int"`
	EndDay        int       `gorm:"type:int"`
	CreatedAt     time.Time `gorm:"type:timestamp"`
	UpdatedAt     time.Time `gorm:"type:timestamp"`
}

// TableName to get table name.
func (m Manga) TableName() string {
	return "manga"
}
