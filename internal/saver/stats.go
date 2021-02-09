package saver

import (
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

func (a *API) saveStats(tx *gorm.DB, t string, id int) (err error) {
	// Parse anime stats.
	var stats *model.Stats
	switch t {
	case constant.AnimeType:
		stats, _, err = a.mal.GetAnimeStats(id)
	case constant.MangaType:
		stats, _, err = a.mal.GetMangaStats(id)
	}
	if err != nil {
		return err
	}

	// Delete existing data.
	if err := tx.Where("id = ? and type = ?", id, t).Delete(raw.Stats{}).Error; err != nil {
		return err
	}

	s := raw.Stats{
		ID:        id,
		Type:      t,
		Current:   stats.Summary.Current,
		Completed: stats.Summary.Completed,
		OnHold:    stats.Summary.OnHold,
		Dropped:   stats.Summary.Dropped,
		Planned:   stats.Summary.Planned,
		Score1:    stats.Score.Score1.Vote,
		Score2:    stats.Score.Score2.Vote,
		Score3:    stats.Score.Score3.Vote,
		Score4:    stats.Score.Score4.Vote,
		Score5:    stats.Score.Score5.Vote,
		Score6:    stats.Score.Score6.Vote,
		Score7:    stats.Score.Score7.Vote,
		Score8:    stats.Score.Score8.Vote,
		Score9:    stats.Score.Score9.Vote,
		Score10:   stats.Score.Score10.Vote,
	}

	return tx.Create(&s).Error
}
