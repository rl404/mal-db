package raw

// MediaRelated is model for media_related table containing
// relation between related anime & manga.
type MediaRelated struct {
	MediaID       int    `gorm:"primary_key"`
	MediaType     string `gorm:"primary_key;type:varchar"` // anime/manga
	RelatedTypeID int    `gorm:"primary_key"`
	RelatedID     int    `gorm:"primary_key"`
	RelatedType   string `gorm:"primary_key;type:varchar"` // anime/manga
}

// TableName to get table name.
func (mr MediaRelated) TableName() string {
	return "media_related"
}
