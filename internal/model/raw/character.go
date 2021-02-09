package raw

import "time"

// Character is model for character table.
type Character struct {
	ID           int       `gorm:"primary_key"`
	Name         string    `gorm:"type:varchar"`
	Nickname     string    `gorm:"type:varchar"`
	JapaneseName string    `gorm:"type:varchar"`
	ImageURL     string    `gorm:"type:varchar"`
	Favorite     int       `gorm:"type:int"`
	About        string    `gorm:"type:varchar"`
	CreatedAt    time.Time `gorm:"type:timestamp"`
	UpdatedAt    time.Time `gorm:"type:timestamp"`
}

// TableName to get table name.
func (c Character) TableName() string {
	return "character"
}
