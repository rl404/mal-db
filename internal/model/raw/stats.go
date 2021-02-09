package raw

// Stats is model for stats table containing
// anime & manga progress count and score.
type Stats struct {
	ID        int    `gorm:"primary_key"`
	Type      string `gorm:"primary_key;type:varchar"` // anime/manga
	Current   int    `gorm:"type:int"`
	Completed int    `gorm:"type:int"`
	OnHold    int    `gorm:"type:int"`
	Dropped   int    `gorm:"type:int"`
	Planned   int    `gorm:"type:int"`
	Score1    int    `gorm:"column:score_1"`
	Score2    int    `gorm:"column:score_2"`
	Score3    int    `gorm:"column:score_3"`
	Score4    int    `gorm:"column:score_4"`
	Score5    int    `gorm:"column:score_5"`
	Score6    int    `gorm:"column:score_6"`
	Score7    int    `gorm:"column:score_7"`
	Score8    int    `gorm:"column:score_8"`
	Score9    int    `gorm:"column:score_9"`
	Score10   int    `gorm:"column:score_10"`
}

// TableName to get table name.
func (s Stats) TableName() string {
	return "stats"
}
