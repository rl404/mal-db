package raw

// MangaCharacter is model for manga_character table containing
// manga and character relation.
type MangaCharacter struct {
	MangaID     int    `gorm:"primary_key"`
	CharacterID int    `gorm:"primary_key"`
	Role        string `gorm:"type:varchar"` // main/supporting
}

// TableName to get table name.
func (mc MangaCharacter) TableName() string {
	return "manga_character"
}
