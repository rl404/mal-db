package model

// Recommendation represents main recommendation model.
type Recommendation struct {
	Source      Source               `json:"source"`
	Recommended Source               `json:"recommended"`
	Users       []RecommendationUser `json:"users"`
}

// RecommendationUser represents user recommendation model.
type RecommendationUser struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}
