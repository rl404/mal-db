package join

// MediaRelated represents joined related anime/manga model.
type MediaRelated struct {
	ID            int
	Title         string
	ImageURL      string
	RelatedType   string
	RelatedTypeID int
}
