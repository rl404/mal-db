package raw

// Language is model for language table containing
// voice actor language constant values.
type Language struct {
	ID       int    `gorm:"primary_key"`
	Language string `gorm:"type:varchar"`
}

// TableName to get table name.
func (l Language) TableName() string {
	return "language"
}
