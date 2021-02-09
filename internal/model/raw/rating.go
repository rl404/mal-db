package raw

// Rating is model for rating table containing
// anime rating constant values.
type Rating struct {
	ID          int    `gorm:"primary_key"`
	Rating      string `gorm:"type:varchar"`
	Description string `gorm:"type:varchar"`
}

// TableName to get table name.
func (r Rating) TableName() string {
	return "rating"
}
