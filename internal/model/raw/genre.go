package raw

// Genre is model for genre table containing
// anime & manga complete genre list.
type Genre struct {
	ID    int    `gorm:"primary_key"`
	Type  string `gorm:"primary_key;type:varchar"` // anime/manga
	Genre string `gorm:"type:varchar"`
}

// TableName to get table name.
func (g Genre) TableName() string {
	return "genre"
}
