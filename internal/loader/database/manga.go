package database

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

// GetManga to get manga.
func (d *Database) GetManga(id int) (*model.Manga, map[string]interface{}, int, error) {
	// Is empty.
	if d.isEntryEmpty(constant.MangaType, id) {
		return nil, nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	// Retrieve from db.
	var mangaRaw raw.Manga
	if errors.Is(d.db.Where("id = ?", id).First(&mangaRaw).Error, gorm.ErrRecordNotFound) {
		if err := d.enqueue(constant.MangaType, id); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		return nil, nil, http.StatusAccepted, _errors.ErrQueueEntry
	}

	// Fill data.
	manga := &model.Manga{
		ID:    mangaRaw.ID,
		Title: mangaRaw.Title,
		AlternativeTitles: model.AlternativeTitle{
			English:  mangaRaw.TitleEnglish,
			Japanese: mangaRaw.TitleJapanese,
			Synonym:  mangaRaw.TitleSynonym,
		},
		Image:      mangaRaw.ImageURL,
		Synopsis:   mangaRaw.Synopsis,
		Score:      mangaRaw.Score,
		Voter:      mangaRaw.Voter,
		Rank:       mangaRaw.Rank,
		Popularity: mangaRaw.Popularity,
		Member:     mangaRaw.Member,
		Favorite:   mangaRaw.Favorite,
		Type:       mangaRaw.MangaTypeID,
		Volume:     mangaRaw.Volume,
		Chapter:    mangaRaw.Chapter,
		Status:     mangaRaw.MangaStatusID,
		Publishing: model.Publishing{
			Start: model.Date{
				Year:  mangaRaw.StartYear,
				Month: mangaRaw.StartMonth,
				Day:   mangaRaw.StartDay,
			},
			End: model.Date{
				Year:  mangaRaw.EndYear,
				Month: mangaRaw.EndMonth,
				Day:   mangaRaw.EndDay,
			},
		},
		Genres:         d.getMediaGenre(constant.MangaType, id),
		Authors:        d.getMangaAuthor(id),
		Serializations: d.getMangaMagazine(id),
		Related:        d.getRelated(constant.MangaType, id),
	}

	// Prepare meta.
	meta := map[string]interface{}{
		"parsedAt": mangaRaw.UpdatedAt,
	}

	return manga, meta, http.StatusOK, nil
}

func (d *Database) getMangaAuthor(id int) (role []model.Role) {
	err := d.db.Table(fmt.Sprintf("%s as pm", raw.PeopleManga{}.TableName())).
		Select("pm.people_id as id, p.image_url as image, p.name, p2.position as role").
		Joins(fmt.Sprintf("left join %s as p on p.id = pm.people_id", raw.People{}.TableName())).
		Joins(fmt.Sprintf("left join %s as p2 on p2.id = pm.position_id", raw.Position{}.TableName())).
		Where("pm.manga_id = ?", id).
		Find(&role).Error
	if err != nil {
		d.log.Error(err.Error())
	}
	return role
}

func (d *Database) getMangaMagazine(id int) (mags []model.Item) {
	err := d.db.Table(fmt.Sprintf("%s as mm", raw.MangaMagazine{}.TableName())).
		Select("pm.id, pm.name").
		Joins(fmt.Sprintf("left join %s as pm on pm.id = mm.magazine_id", raw.ProducerMagazine{}.TableName())).
		Where("mm.manga_id = ? and pm.type = ?", id, constant.MangaType).
		Find(&mags).Error
	if err != nil {
		d.log.Error(err.Error())
	}
	return mags
}
