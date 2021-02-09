package model

// Anime represents anime model retrieved from db.
type Anime struct {
	ID                int              `json:"id"`
	Title             string           `json:"title"`
	AlternativeTitles AlternativeTitle `json:"alternativeTitles"`
	Image             string           `json:"image"`
	Video             string           `json:"video"`
	Synopsis          string           `json:"synopsis"`
	Score             float64          `json:"score"`
	Voter             int              `json:"voter"`
	Rank              int              `json:"rank"`
	Popularity        int              `json:"popularity"`
	Member            int              `json:"member"`
	Favorite          int              `json:"favorite"`
	Type              int              `json:"type"`
	Episode           int              `json:"episode"`
	Status            int              `json:"status"`
	Airing            Airing           `json:"airing"`
	Duration          string           `json:"duration"`
	Premiered         string           `json:"premiered"`
	Source            int              `json:"source"`
	Rating            int              `json:"rating"`
	Producers         []Item           `json:"producers"`
	Licensors         []Item           `json:"licensors"`
	Studios           []Item           `json:"studios"`
	Genres            []Item           `json:"genres"`
	Related           Related          `json:"related"`
	Songs             Song             `json:"songs"`
}

// Airing represents anime airing details.
type Airing struct {
	Start Date   `json:"start"`
	End   Date   `json:"end"`
	Day   string `json:"day"`
	Time  string `json:"time"`
}

// Song represents anime opening and ending songs.
type Song struct {
	Opening []string `json:"opening"`
	Ending  []string `json:"ending"`
}

// AnimeCharacter represents anime's characters and their voice actors.
type AnimeCharacter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Role        string `json:"role"`
	VoiceActors []Role `json:"voiceActors"`
}
