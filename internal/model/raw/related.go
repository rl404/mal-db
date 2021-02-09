package raw

// Related is model for related table containing
// anime & manga related type constant values.
type Related struct {
	ID      int    `gorm:"primary_key"`
	Related string `gorm:"type:varchar"`
}

// TableName to get table name.
func (r Related) TableName() string {
	return "related"
}
