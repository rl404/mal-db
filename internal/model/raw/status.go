package raw

// Status is model for status table containing
// anime & manga status constant value.
type Status struct {
	ID   int    `gorm:"primary_key"`
	Type string `gorm:"primary_key;type:varchar"` // anime/manga
	Name string `gorm:"type:varchar"`
}

// TableName to get table name.
func (s Status) TableName() string {
	return "status"
}
