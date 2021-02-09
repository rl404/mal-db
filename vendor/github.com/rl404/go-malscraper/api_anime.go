package malscraper

import "github.com/rl404/go-malscraper/model"

// GetAnime to get anime detail information.
//
// Example: https://myanimelist.net/anime/1.
func (m *Malscraper) GetAnime(id int) (*model.Anime, int, error) {
	return m.api.GetAnime(id)
}

// GetAnimeCharacter to get anime character list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/characters.
func (m *Malscraper) GetAnimeCharacter(id int) ([]model.CharacterItem, int, error) {
	return m.api.GetAnimeCharacter(id)
}

// GetAnimeStaff to get anime staff list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/characters.
func (m *Malscraper) GetAnimeStaff(id int) ([]model.Role, int, error) {
	return m.api.GetAnimeStaff(id)
}

// GetAnimeVideo to get anime video list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/video.
func (m *Malscraper) GetAnimeVideo(id int, page ...int) (*model.Video, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetAnimeVideo(id, p)
}

// GetAnimeEpisode to get anime episode list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/episode.
func (m *Malscraper) GetAnimeEpisode(id int, page ...int) ([]model.Episode, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetAnimeEpisode(id, p)
}

// GetAnimeStats to get anime stats.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/stats.
func (m *Malscraper) GetAnimeStats(id int) (*model.Stats, int, error) {
	return m.api.GetAnimeStats(id)
}

// GetAnimeReview to get anime review list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/reviews.
func (m *Malscraper) GetAnimeReview(id int, page ...int) ([]model.Review, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetAnimeReview(id, p)
}

// GetAnimeRecommendation to get anime recommendation list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/userrecs
func (m *Malscraper) GetAnimeRecommendation(id int) ([]model.Recommendation, int, error) {
	return m.api.GetAnimeRecommendation(id)
}

// GetAnimeNews to get anime news list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/news.
func (m *Malscraper) GetAnimeNews(id int) ([]model.NewsItem, int, error) {
	return m.api.GetAnimeNews(id)
}

// GetAnimeArticle to get anime featured article list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/featured.
func (m *Malscraper) GetAnimeArticle(id int) ([]model.ArticleItem, int, error) {
	return m.api.GetAnimeArticle(id)
}

// GetAnimeClub to get anime club list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/clubs.
func (m *Malscraper) GetAnimeClub(id int) ([]model.ClubItem, int, error) {
	return m.api.GetAnimeClub(id)
}

// GetAnimePicture to get anime picture list.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/pics.
func (m *Malscraper) GetAnimePicture(id int) ([]string, int, error) {
	return m.api.GetAnimePicture(id)
}

// GetAnimeMoreInfo to get anime more info.
//
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/moreinfo.
func (m *Malscraper) GetAnimeMoreInfo(id int) (string, int, error) {
	return m.api.GetAnimeMoreInfo(id)
}
