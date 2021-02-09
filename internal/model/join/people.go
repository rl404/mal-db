package join

// PeopleVA represents joined model for people's voice actor roles.
type PeopleVA struct {
	AID      int    `gorm:"column:a_id"`
	ATitle   string `gorm:"column:a_title"`
	AImage   string `gorm:"column:a_image"`
	Role     string `gorm:"column:role"`
	Language string `gorm:"column:language"`
	CID      int    `gorm:"column:c_id"`
	CName    string `gorm:"column:c_name"`
	CImage   string `gorm:"column:c_image"`
}
