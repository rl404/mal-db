package raw

// Type is model for media_type table containing
// anime & manga type constant value.
type Type struct {
	ID   int    `gorm:"primary_key"`
	Type string `gorm:"primary_key;type:varchar"` // anime/manga
	Name string `gorm:"type:varchar"`
}

// TableName to get table name.
func (t Type) TableName() string {
	return "type"
}
