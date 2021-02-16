package api

import "github.com/rl404/mal-db/internal/model"

// API is loader interface.
type API interface {
	GetAnime(id int) (*model.Anime, map[string]interface{}, int, error)
	GetAnimeCharacter(id int, page int, limit int) ([]model.AnimeCharacter, map[string]interface{}, int, error)
	GetAnimeStaff(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error)
	GetManga(id int) (*model.Manga, map[string]interface{}, int, error)
	GetMangaCharacter(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error)
	GetStats(_type string, id int) (*model.Stats, map[string]interface{}, int, error)
	GetCharacter(id int) (*model.Character, map[string]interface{}, int, error)
	GetCharacterOgraphy(id int, _type string, page int, limit int) ([]model.Role, map[string]interface{}, int, error)
	GetCharacterVA(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error)
	GetPeople(id int) (*model.People, map[string]interface{}, int, error)
	GetPeopleVA(id int, page int, limit int) ([]model.VoiceActor, map[string]interface{}, int, error)
	GetPeopleStaff(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error)
	GetPeopleManga(id int, page int, limit int) ([]model.Role, map[string]interface{}, int, error)
	GetProducerMagazine(t string) ([]model.Item, map[string]interface{}, int, error)
	GetGenres(t string) ([]model.Item, map[string]interface{}, int, error)
	SearchAnime(query model.AnimeQuery) ([]model.Media, map[string]interface{}, int, error)
	SearchManga(query model.MangaQuery) ([]model.Media, map[string]interface{}, int, error)
	SearchCharacter(query model.EntryQuery) ([]model.Entry, map[string]interface{}, int, error)
	SearchPeople(query model.EntryQuery) ([]model.Entry, map[string]interface{}, int, error)
	GetEntryCount() (*model.Total, map[string]interface{}, int, error)
	GetYearSummary() ([]model.YearSummary, map[string]interface{}, int, error)
	Enqueue(_type string, id int) (int, error)
}
