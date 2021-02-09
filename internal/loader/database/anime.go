package database

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/join"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

// GetAnime to get anime from db.
func (d *Database) GetAnime(id int) (*model.Anime, map[string]interface{}, int, error) {
	// Is empty.
	if d.isEntryEmpty(constant.AnimeType, id) {
		return nil, nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	// Retrieve from db.
	var animeRaw raw.Anime
	if errors.Is(d.db.Where("id = ?", id).First(&animeRaw).Error, gorm.ErrRecordNotFound) {
		// Enqueue if not exists.
		if err := d.enqueue(constant.AnimeType, id); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		return nil, nil, http.StatusAccepted, _errors.ErrQueueEntry
	}

	// Fill data.
	anime := &model.Anime{
		ID:    animeRaw.ID,
		Title: animeRaw.Title,
		AlternativeTitles: model.AlternativeTitle{
			English:  animeRaw.TitleEnglish,
			Japanese: animeRaw.TitleJapanese,
			Synonym:  animeRaw.TitleSynonym,
		},
		Image:      animeRaw.ImageURL,
		Video:      animeRaw.VideoURL,
		Synopsis:   animeRaw.Synopsis,
		Score:      animeRaw.Score,
		Voter:      animeRaw.Voter,
		Rank:       animeRaw.Rank,
		Popularity: animeRaw.Popularity,
		Member:     animeRaw.Member,
		Favorite:   animeRaw.Favorite,
		Type:       animeRaw.AnimeTypeID,
		Episode:    animeRaw.Episode,
		Status:     animeRaw.AnimeStatusID,
		Airing: model.Airing{
			Start: model.Date{
				Year:  animeRaw.StartYear,
				Month: animeRaw.StartMonth,
				Day:   animeRaw.StartDay,
			},
			End: model.Date{
				Year:  animeRaw.EndYear,
				Month: animeRaw.EndMonth,
				Day:   animeRaw.EndDay,
			},
			Day:  animeRaw.AiringDay,
			Time: animeRaw.AiringTime,
		},
		Duration:  utils.SecondToString(animeRaw.Duration),
		Premiered: animeRaw.Premiered,
		Source:    animeRaw.AnimeSourceID,
		Rating:    animeRaw.AnimeRatingID,
		Genres:    d.getMediaGenre(constant.AnimeType, id),
		Related:   d.getRelated(constant.AnimeType, id),
		Songs:     d.getAnimeSong(id),
	}

	anime.Producers, anime.Studios, anime.Licensors = d.getAnimeProducer(id)

	// Prepare meta.
	meta := map[string]interface{}{
		"parsedAt": animeRaw.UpdatedAt,
	}

	return anime, meta, http.StatusOK, nil
}

func (d *Database) getAnimeProducer(id int) (producers, studios, licensors []model.Item) {
	var animeProducer []join.AnimeProducer
	err := d.db.Table(fmt.Sprintf("%s as ap", raw.AnimeProducer{}.TableName())).
		Select("pm.id, pm.name, ap.is_licensor, ap.is_studio").
		Joins(fmt.Sprintf("left join %s as pm on pm.id = ap.producer_id", raw.ProducerMagazine{}.TableName())).
		Where("ap.anime_id = ? and pm.type = ?", id, constant.AnimeType).
		Find(&animeProducer).Error
	if err != nil {
		d.log.Error(err.Error())
	}

	producers, studios, licensors = []model.Item{}, []model.Item{}, []model.Item{}
	for _, p := range animeProducer {
		if !p.IsLicensor && !p.IsStudio {
			producers = append(producers, model.Item{
				ID:   p.ID,
				Name: p.Name,
			})
		} else if !p.IsLicensor && p.IsStudio {
			studios = append(studios, model.Item{
				ID:   p.ID,
				Name: p.Name,
			})
		} else if p.IsLicensor && !p.IsStudio {
			licensors = append(licensors, model.Item{
				ID:   p.ID,
				Name: p.Name,
			})
		}
	}
	return producers, studios, licensors
}

func (d *Database) getAnimeSong(id int) model.Song {
	var songs []raw.Song
	if err := d.db.Select("type, song").Where("anime_id = ?", id).Find(&songs).Error; err != nil {
		d.log.Error(err.Error())
	}

	result := model.Song{
		Opening: []string{},
		Ending:  []string{},
	}

	for _, s := range songs {
		if s.Type == constant.OpeningSong {
			result.Opening = append(result.Opening, s.Song)
		}

		if s.Type == constant.EndingSong {
			result.Ending = append(result.Ending, s.Song)
		}
	}

	return result
}
