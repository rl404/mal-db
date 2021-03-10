package raw

// AnimeProducer is many-to-many relation model between anime and producer.
type AnimeProducer struct {
	AnimeID    int  `gorm:"primary_key"`
	ProducerID int  `gorm:"primary_key"`
	IsLicensor bool `gorm:"primary_key;type:bool;default:false"`
	IsStudio   bool `gorm:"primary_key;type:bool;default:false"`
}

// TableName to get table name.
func (ap AnimeProducer) TableName() string {
	return "anime_producer"
}
