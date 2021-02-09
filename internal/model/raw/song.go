package raw

// Song is model for song table.
type Song struct {
	ID      int    `gorm:"primary_key"`
	AnimeID int    `gorm:"primary_key"`
	Type    int    `gorm:"primary_key"` // 1=op, 2=ed
	Song    string `gorm:"type:varchar"`
}

// TableName to get table name.
func (s Song) TableName() string {
	return "song"
}
