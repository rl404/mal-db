package raw

// EmptyID is model for empty_id table containing
// empty MyAnimeList entry id.
type EmptyID struct {
	Type string `gorm:"primary_key;type:varchar"`
	ID   int    `gorm:"primary_key"`
}
