package raw

// Position is model for position table containing
// anime staff position constant values.
type Position struct {
	ID       int    `gorm:"primary_key"`
	Position string `gorm:"type:varchar"`
}

// TableName to get table name.
func (p Position) TableName() string {
	return "position"
}
