package raw

// MangaMagazine is model for manga_magazine table containing
// manga and magazine (serialization) relation.
type MangaMagazine struct {
	MangaID    int `gorm:"primary_key"`
	MagazineID int `gorm:"primary_key"`
}

// TableName to get table name.
func (mm MangaMagazine) TableName() string {
	return "manga_magazine"
}
