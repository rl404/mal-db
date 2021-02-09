package raw

// AnimeStaff is model for anime_staff table containing
// relation between anime, staff, and position.
type AnimeStaff struct {
	AnimeID    int `gorm:"primary_key"`
	PeopleID   int `gorm:"primary_key"`
	PositionID int `gorm:"primary_key"`
}

// TableName to get table name.
func (as AnimeStaff) TableName() string {
	return "anime_staff"
}
