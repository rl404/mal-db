package raw

// PeopleManga is model for people_manga table containing
// people (author) and manga relation.
type PeopleManga struct {
	PeopleID   int `gorm:"primary_key"`
	MangaID    int `gorm:"primary_key"`
	PositionID int `gorm:"primary_key"`
}

// TableName to get table name.
func (pm PeopleManga) TableName() string {
	return "people_manga"
}
