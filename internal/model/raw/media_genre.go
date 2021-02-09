package raw

// MediaGenre is model for media_genre table containing
// relation between anime/manga and genre.
type MediaGenre struct {
	Type    string `gorm:"primary_key;type:varchar"` // anime/manga
	MediaID int    `gorm:"primary_key"`
	GenreID int    `gorm:"primary_key"`
}

// TableName to get table name.
func (mg MediaGenre) TableName() string {
	return "media_genre"
}
