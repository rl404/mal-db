package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetCharacter to get character detail information.
//
// Example: https://myanimelist.net/character/1/Spike_Spiegel.
func (m *Malscraper) GetCharacter(id int) (*model.Character, int, error) {
	return m.api.GetCharacter(id)
}

// GetCharacterArticle to get character featured article list.
//
// Example: https://myanimelist.net/character/1/Spike_Spiegel/featured.
func (m *Malscraper) GetCharacterArticle(id int) ([]model.ArticleItem, int, error) {
	return m.api.GetCharacterArticle(id)
}

// GetCharacterOgraphy to get character animeography/mangaography list.
//
// Param `_type` should be one of these constants.
//
//  AnimeType
//  MangaType
//
// Or just use method `GetCharacterAnime()` or `GetCharacterManga()`.
func (m *Malscraper) GetCharacterOgraphy(_type int, id int) ([]model.Role, int, error) {
	return m.api.GetCharacterOgraphy(mainTypes[_type], id)
}

// GetCharacterAnime to get character animeography list.
//
// Example: https://myanimelist.net/character/1/Spike_Spiegel.
func (m *Malscraper) GetCharacterAnime(id int) ([]model.Role, int, error) {
	return m.GetCharacterOgraphy(AnimeType, id)
}

// GetCharacterManga to get character mangaography list.
//
// Example: https://myanimelist.net/character/1/Spike_Spiegel.
func (m *Malscraper) GetCharacterManga(id int) ([]model.Role, int, error) {
	return m.GetCharacterOgraphy(MangaType, id)
}

// GetCharacterPicture to get character picture list.
//
// Example: https://myanimelist.net/character/1/Spike_Spiegel/pictures.
func (m *Malscraper) GetCharacterPicture(id int) ([]string, int, error) {
	return m.api.GetCharacterPicture(id)
}

// GetCharacterClub to get character club list.
//
// Example: https://myanimelist.net/character/1/Spike_Spiegel/clubs.
func (m *Malscraper) GetCharacterClub(id int) ([]model.ClubItem, int, error) {
	return m.api.GetCharacterClub(id)
}

// GetCharacterVA to get character voice actor list.
//
// Example: https://myanimelist.net/character/1/Spike_Spiegel.
func (m *Malscraper) GetCharacterVA(id int) ([]model.Role, int, error) {
	return m.api.GetCharacterVA(id)
}
