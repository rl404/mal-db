package raw

// ProducerMagazine is model for producer_magazine table containing
// all anime producers and manga magazines.
type ProducerMagazine struct {
	ID   int    `gorm:"primary_key;"`
	Type string `gorm:"primary_key;type:varchar"` // anime/manga
	Name string `gorm:"type:varchar"`
}

// TableName to get table name.
func (p ProducerMagazine) TableName() string {
	return "producer_magazine"
}
