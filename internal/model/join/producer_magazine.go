package join

// AnimeProducer represents joined anime's producer model.
type AnimeProducer struct {
	ID         int
	Name       string
	IsLicensor bool
	IsStudio   bool
}
