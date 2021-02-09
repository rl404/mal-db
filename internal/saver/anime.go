package saver

import (
	"errors"
	"net/http"
	"strings"

	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

func (a *API) parseAnime(id int) error {
	// Check old entry.
	if a.isEntryNew(constant.AnimeType, id) {
		return _errors.ErrNewData
	}

	// Delete existing empty id.
	if err := a.deleteEmptyID(constant.AnimeType, id); err != nil {
		return err
	}

	// Parse anime.
	data, code, err := a.mal.GetAnime(id)

	// Empty id.
	if code == http.StatusNotFound {
		if err = a.insertEmptyID(constant.AnimeType, id); err != nil {
			return err
		}
		a.logger.Debug("empty anime %v", id)
		return nil
	}

	// Other error.
	if err != nil {
		return err
	}

	// Begin transaction.
	tx := a.db.Begin()
	defer tx.Rollback()

	// Get existing anime db.
	var anime raw.Anime
	err = tx.Where("id = ?", id).First(&anime).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Update data.
	anime.ID = data.ID
	anime.Title = data.Title
	anime.TitleEnglish = data.AlternativeTitles.English
	anime.TitleJapanese = data.AlternativeTitles.Japanese
	anime.TitleSynonym = data.AlternativeTitles.Synonym
	anime.ImageURL = data.Image
	anime.VideoURL = data.Video
	anime.Synopsis = data.Synopsis
	anime.Score = data.Score
	anime.Voter = data.Voter
	anime.Rank = data.Rank
	anime.Popularity = data.Popularity
	anime.Member = data.Member
	anime.Favorite = data.Favorite
	anime.AnimeTypeID = utils.GetMapKey(constant.Types[constant.AnimeType], data.Type)
	anime.Episode = data.Episode
	anime.AnimeStatusID = utils.GetMapKey(constant.Statuses[constant.AnimeType], data.Status)
	anime.StartYear = data.AiringDate.Start.Year
	anime.StartMonth = data.AiringDate.Start.Month
	anime.StartDay = data.AiringDate.Start.Day
	anime.EndYear = data.AiringDate.End.Year
	anime.EndMonth = data.AiringDate.End.Month
	anime.EndDay = data.AiringDate.End.Day
	anime.AiringDay, anime.AiringTime = a.getAiringDayTime(data.Broadcast)
	anime.Premiered = strings.ToLower(data.Premiered)
	anime.AnimeSourceID = utils.GetMapKey(constant.Sources, data.Source)
	anime.Duration = utils.GetDuration(data.Duration)
	anime.AnimeRatingID = utils.GetMapKey(constant.Ratings, data.Rating)

	// Save anime.
	if err = tx.Save(&anime).Error; err != nil {
		return err
	}

	// Save history.
	if err = a.saveAnimeHistory(tx, anime); err != nil {
		return err
	}

	// Save producers.
	if err = a.saveAnimeProducer(tx, data); err != nil {
		return err
	}

	// Save genres.
	if err = a.saveGenre(tx, constant.AnimeType, data.ID, data.Genres); err != nil {
		return err
	}

	// Save related.
	if err = a.saveRelated(tx, constant.AnimeType, data.ID, data.Related); err != nil {
		return err
	}

	// Save characters.
	if err = a.saveAnimeCharacter(tx, data); err != nil {
		return err
	}

	// Save staff.
	if err = a.saveAnimeStaff(tx, data); err != nil {
		return err
	}

	// Save song.
	if err = a.saveSong(tx, data); err != nil {
		return err
	}

	// Save stats.
	if err = a.saveStats(tx, constant.AnimeType, data.ID); err != nil {
		return err
	}

	// Commit.
	if err = tx.Commit().Error; err != nil {
		return err
	}

	// Delete cache.
	return a.cacher.Delete(constant.GetKey(constant.KeyAnime, id))
}

func (a *API) getAiringDayTime(daytime string) (string, string) {
	def := "00:00"
	if daytime == "" || daytime == "Not scheduled once per week" {
		return "", def
	}

	daytime = strings.Replace(daytime, "(JST)", "", -1)
	split := strings.Split(daytime, " at")

	if split[1] == "" {
		return split[0], def
	}

	split[0] = strings.ToLower(split[0])

	// Remove letter 's' at the end.
	if last := len(split[0]) - 1; last >= 0 && split[0][last] == 's' {
		split[0] = split[0][:last]
	}

	return split[0], strings.TrimSpace(split[1])
}
