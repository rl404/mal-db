package saver

import (
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

func (a *API) saveAnimeHistory(tx *gorm.DB, anime raw.Anime) error {
	return tx.Create(&raw.StatsHistory{
		MediaID:    anime.ID,
		Type:       constant.AnimeType,
		Score:      anime.Score,
		Voter:      anime.Voter,
		Rank:       anime.Rank,
		Popularity: anime.Popularity,
		Member:     anime.Member,
		Favorite:   anime.Favorite,
	}).Error
}

func (a *API) saveMangaHistory(tx *gorm.DB, manga raw.Manga) error {
	return tx.Create(&raw.StatsHistory{
		MediaID:    manga.ID,
		Type:       constant.MangaType,
		Score:      manga.Score,
		Voter:      manga.Voter,
		Rank:       manga.Rank,
		Popularity: manga.Popularity,
		Member:     manga.Member,
		Favorite:   manga.Favorite,
	}).Error
}

func (a *API) saveCharacterHistory(tx *gorm.DB, char raw.Character) error {
	return tx.Create(&raw.StatsHistory{
		MediaID:  char.ID,
		Type:     constant.CharacterType,
		Favorite: char.Favorite,
	}).Error
}

func (a *API) savePeopleHistory(tx *gorm.DB, people raw.People) error {
	return tx.Create(&raw.StatsHistory{
		MediaID:  people.ID,
		Type:     constant.PeopleType,
		Favorite: people.Favorite,
	}).Error
}
