package join

// AnimeCharacter represents joined anime's characters model.
type AnimeCharacter struct {
	CID        int    `gorm:"column:c_id"`
	CName      string `gorm:"column:c_name"`
	CImage     string `gorm:"column:c_image"`
	Role       string `gorm:"column:role"`
	LanguageID int    `gorm:"column:language_id"`
	PID        int    `gorm:"column:p_id"`
	PName      string `gorm:"column:p_name"`
	PImage     string `gorm:"column:p_image"`
}
