package join

// MediaRelated represents joined related anime/manga model.
type MediaRelated struct {
	ID            int
	Title         string
	ImageURL      string
	RelatedType   string
	RelatedTypeID int
}

// ScoreComparison represents joined related anime/manga model
// for comparison.
type ScoreComparison struct {
	NID    int     `gorm:"column:n_id"`
	NTitle string  `gorm:"column:n_title"`
	NScore float64 `gorm:"column:n_score"`
	AID    int     `gorm:"column:a_id"`
	ATitle string  `gorm:"column:a_title"`
	AScore float64 `gorm:"column:a_score"`
	MID    int     `gorm:"column:m_id"`
	MTitle string  `gorm:"column:m_title"`
	MScore float64 `gorm:"column:m_score"`
}
