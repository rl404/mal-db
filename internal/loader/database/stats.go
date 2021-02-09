package database

import (
	"errors"
	"net/http"

	"github.com/rl404/go-malscraper/pkg/utils"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
	"gorm.io/gorm"
)

// GetStats to get anime/manga stats.
func (d *Database) GetStats(t string, id int) (*model.Stats, map[string]interface{}, int, error) {
	// Is empty.
	if d.isEntryEmpty(t, id) {
		return nil, nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	// Retrieve from db.
	var statsRaw raw.Stats
	if errors.Is(d.db.Where("type = ? and id = ?", t, id).First(&statsRaw).Error, gorm.ErrRecordNotFound) {
		// Enqueue if not exists.
		if err := d.enqueue(t, id); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		return nil, nil, http.StatusAccepted, _errors.ErrQueueEntry
	}

	// Fill data.
	total := statsRaw.Score1 + statsRaw.Score2 + statsRaw.Score3 + statsRaw.Score4 + statsRaw.Score5 + statsRaw.Score6 + statsRaw.Score7 + statsRaw.Score8 + statsRaw.Score9 + statsRaw.Score10
	stats := &model.Stats{
		Summary: model.Summary{
			Current:   statsRaw.Current,
			Completed: statsRaw.Completed,
			OnHold:    statsRaw.OnHold,
			Dropped:   statsRaw.Dropped,
			Planned:   statsRaw.Planned,
		},
		Score: model.Score{
			Score1:  model.ScoreDetail{Vote: statsRaw.Score1, Percent: utils.GetPercent(statsRaw.Score1, total, 2)},
			Score2:  model.ScoreDetail{Vote: statsRaw.Score2, Percent: utils.GetPercent(statsRaw.Score2, total, 2)},
			Score3:  model.ScoreDetail{Vote: statsRaw.Score3, Percent: utils.GetPercent(statsRaw.Score3, total, 2)},
			Score4:  model.ScoreDetail{Vote: statsRaw.Score4, Percent: utils.GetPercent(statsRaw.Score4, total, 2)},
			Score5:  model.ScoreDetail{Vote: statsRaw.Score5, Percent: utils.GetPercent(statsRaw.Score5, total, 2)},
			Score6:  model.ScoreDetail{Vote: statsRaw.Score6, Percent: utils.GetPercent(statsRaw.Score6, total, 2)},
			Score7:  model.ScoreDetail{Vote: statsRaw.Score7, Percent: utils.GetPercent(statsRaw.Score7, total, 2)},
			Score8:  model.ScoreDetail{Vote: statsRaw.Score8, Percent: utils.GetPercent(statsRaw.Score8, total, 2)},
			Score9:  model.ScoreDetail{Vote: statsRaw.Score9, Percent: utils.GetPercent(statsRaw.Score9, total, 2)},
			Score10: model.ScoreDetail{Vote: statsRaw.Score10, Percent: utils.GetPercent(statsRaw.Score10, total, 2)},
		},
	}

	return stats, nil, http.StatusOK, nil
}
