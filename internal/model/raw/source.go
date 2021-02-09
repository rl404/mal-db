package raw

// Source is model for anime_source table containing
// anime source constant values.
type Source struct {
	ID     int    `gorm:"primary_key"`
	Source string `gorm:"type:varchar"`
}

// TableName to get table name.
func (s Source) TableName() string {
	return "source"
}
