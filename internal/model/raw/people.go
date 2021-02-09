package raw

import "time"

// People is model for people table.
type People struct {
	ID              int       `gorm:"primary_key"`
	Name            string    `gorm:"type:varchar"`
	GivenName       string    `gorm:"type:varchar"`
	FamilyName      string    `gorm:"type:varchar"`
	AlternativeName string    `gorm:"type:varchar"`
	ImageURL        string    `gorm:"type:varchar"`
	BirthdayYear    int       `gorm:"type:int"`
	BirthdayMonth   int       `gorm:"type:int"`
	BirthdayDay     int       `gorm:"type:int"`
	Website         string    `gorm:"type:varchar"`
	Favorite        int       `gorm:"type:int"`
	More            string    `gorm:"type:varchar"`
	CreatedAt       time.Time `gorm:"type:timestamp"`
	UpdatedAt       time.Time `gorm:"type:timestamp"`
}

// TableName to get table name.
func (p People) TableName() string {
	return "people"
}
