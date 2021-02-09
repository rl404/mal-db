package raw

// AnimeCharacter is model for anime_character table containing
// relation between anime, character, people (voice actor) and language.
type AnimeCharacter struct {
	AnimeID     int    `gorm:"primary_key"`
	CharacterID int    `gorm:"primary_key"`
	PeopleID    int    `gorm:"primary_key"`
	Role        string `gorm:"type:varchar"` // main/supporting
	LanguageID  int    `gorm:"type:int"`
}

// TableName to get table name.
func (ac AnimeCharacter) TableName() string {
	return "anime_character"
}
