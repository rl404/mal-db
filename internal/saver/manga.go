package saver

import (
	"errors"
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) parseManga(id int) error {
	// Check old entry.
	if a.isEntryNew(constant.MangaType, id) {
		return _errors.ErrNewData
	}

	// Delete existing empty id.
	if err := a.deleteEmptyID(constant.MangaType, id); err != nil {
		return err
	}

	// Parse manga.
	data, code, err := a.mal.GetManga(id)

	// Empty id.
	if code == http.StatusNotFound {
		if err = a.insertEmptyID(constant.MangaType, id); err != nil {
			return err
		}
		a.logger.Debug("empty manga %v", id)
		return nil
	}

	// Other error.
	if err != nil {
		return err
	}

	// Begin transaction.
	tx := a.db.Begin()
	defer tx.Rollback()

	// Get existing manga db.
	var manga raw.Manga
	err = tx.Where("id = ?", id).First(&manga).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Update data.
	manga.ID = data.ID
	manga.Title = data.Title
	manga.TitleEnglish = data.AlternativeTitles.English
	manga.TitleJapanese = data.AlternativeTitles.Japanese
	manga.TitleSynonym = data.AlternativeTitles.Synonym
	manga.ImageURL = data.Image
	manga.Synopsis = data.Synopsis
	manga.Score = data.Score
	manga.Voter = data.Voter
	manga.Rank = data.Rank
	manga.Popularity = data.Popularity
	manga.Member = data.Member
	manga.Favorite = data.Favorite
	manga.MangaTypeID = utils.GetMapKey(constant.Types[constant.MangaType], data.Type)
	manga.Volume = data.Volume
	manga.Chapter = data.Chapter
	manga.MangaStatusID = utils.GetMapKey(constant.Statuses[constant.MangaType], data.Status)
	manga.StartYear = data.PublishingDate.Start.Year
	manga.StartMonth = data.PublishingDate.Start.Month
	manga.StartDay = data.PublishingDate.Start.Day
	manga.EndYear = data.PublishingDate.End.Year
	manga.EndMonth = data.PublishingDate.End.Month
	manga.EndDay = data.PublishingDate.End.Day

	// Save manga.
	if err = tx.Save(&manga).Error; err != nil {
		return err
	}

	// Save history.
	if err = a.saveMangaHistory(tx, manga); err != nil {
		return err
	}

	// Save magazines.
	if err = a.saveMangaMagazine(tx, data); err != nil {
		return err
	}

	// Save genres.
	if err = a.saveGenre(tx, constant.MangaType, data.ID, data.Genres); err != nil {
		return err
	}

	// Save related.
	if err = a.saveRelated(tx, constant.MangaType, data.ID, data.Related); err != nil {
		return err
	}

	// Save characters.
	if err = a.saveMangaCharacter(tx, data); err != nil {
		return err
	}

	// Enqueue people.
	if err = a.enqueueMangaPeople(data); err != nil {
		return err
	}

	// Save stats.
	if err = a.saveStats(tx, constant.MangaType, data.ID); err != nil {
		return err
	}

	// Commit.
	if err = tx.Commit().Error; err != nil {
		return err
	}

	// Delete cache.
	return a.cacher.Delete(constant.GetKey(constant.KeyManga, id))
}
