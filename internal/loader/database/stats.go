package database

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rl404/go-malscraper/pkg/utils"
	"github.com/rl404/mal-db/internal/constant"
	_errors "github.com/rl404/mal-db/internal/errors"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/join"
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

// GetStatsHistory to get entry stats history.
func (d *Database) GetStatsHistory(t string, id int, _ int, _ int) ([]model.StatsHistory, int, error) {
	// Is empty.
	if d.isEntryEmpty(t, id) {
		return nil, http.StatusNotFound, _errors.ErrInvalidID
	}

	var history []model.StatsHistory
	err := d.db.Model(&raw.StatsHistory{}).
		Select("date_part('year', created_at) as year, date_part('month', created_at) as month, round(avg(score),2) as score, floor(avg(voter)) as voter, floor(avg(rank)) as rank, floor(avg(popularity)) as popularity, floor(avg(member)) as member, floor(avg(favorite)) as favorite").
		Where("type = ? and media_id = ?", t, id).
		Group("date_part('year', created_at), date_part('month', created_at)").
		Order("date_part('year', created_at) desc, date_part('month', created_at) desc").
		Find(&history).
		Error
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return history, http.StatusOK, nil
}

// CompareScore to get entry score comparison.
func (d *Database) CompareScore(query model.CompareQuery) ([]model.ScoreComparison, map[string]interface{}, int, error) {
	if query.Order != "" {
		sort := " asc"
		if query.Order[0] == '-' {
			sort = " desc"
			query.Order = query.Order[1:]
		}
		query.Order = "m." + query.Order + sort
	} else {
		query.Order = "m.title asc"
	}

	subQuery := d.db.Table(fmt.Sprintf("%s as m", raw.Manga{}.TableName())).
		Select("m.id").
		Joins(fmt.Sprintf("left join %s as mr on m.id = mr.media_id", raw.MediaRelated{}.TableName())).
		Joins(fmt.Sprintf("left join %s as a on a.id = mr.related_id and mr.related_type = ?", raw.Anime{}.TableName()), constant.AnimeType).
		Joins(fmt.Sprintf("left join %s as m2 on m2.id = mr.related_id and mr.related_type = ?", raw.Manga{}.TableName()), constant.MangaType).
		Where("lower(m.title) like ? and mr.media_type = ? and (m.manga_type_id = ? or m.manga_type_id = ?) and (mr.related_type_id = ? or mr.related_type_id = ?)", "%"+query.Title+"%", constant.MangaType, 2, 8, 11, 4).
		Group("m.id, m.score, m.title").
		Order(query.Order + ", m.title").
		Limit(query.Limit).
		Offset(query.Limit * (query.Page - 1))

	rows, err := d.db.Table("(?) as mm", subQuery).
		Select("m.id as n_id, m.title as n_title, m.score as n_score, a.id as a_id, a.title as a_title, a.score as a_score, m2.id as m_id, m2.title as m_title, m2.score as m_score").
		Joins(fmt.Sprintf("left join %s as m on mm.id = m.id", raw.Manga{}.TableName())).
		Joins(fmt.Sprintf("left join %s as mr on m.id = mr.media_id", raw.MediaRelated{}.TableName())).
		Joins(fmt.Sprintf("left join %s as a on a.id = mr.related_id and mr.related_type = ?", raw.Anime{}.TableName()), constant.AnimeType).
		Joins(fmt.Sprintf("left join %s as m2 on m2.id = mr.related_id and mr.related_type = ?", raw.Manga{}.TableName()), constant.MangaType).
		Where("mr.media_type = ? and (mr.related_type_id = ? or mr.related_type_id = ?)", constant.MangaType, 11, 4).
		Order(query.Order + ", m.title, a.title, m2.title").
		Rows()
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	defer rows.Close()

	scores := []model.ScoreComparison{}
	novelMap := make(map[int]*model.ScoreComparison)
	animeMap := make(map[int][]model.EntryScore)
	mangaMap := make(map[int][]model.EntryScore)
	for rows.Next() {
		var tmp join.ScoreComparison
		if err = d.db.ScanRows(rows, &tmp); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}

		if novelMap[tmp.NID] == nil {
			nTmp := model.EntryScore{
				ID:    tmp.NID,
				Title: tmp.NTitle,
				Type:  constant.MangaType,
				Score: tmp.NScore,
			}
			novelMap[tmp.NID] = &model.ScoreComparison{
				Novel: []model.EntryScore{nTmp},
			}
			scores = append(scores, model.ScoreComparison{
				Novel: []model.EntryScore{nTmp},
				Anime: []model.EntryScore{},
				Manga: []model.EntryScore{},
			})
		}

		if tmp.AID != 0 {
			animeMap[tmp.NID] = append(animeMap[tmp.NID], model.EntryScore{
				ID:    tmp.AID,
				Title: tmp.ATitle,
				Type:  constant.AnimeType,
				Score: tmp.AScore,
			})
		}

		if tmp.MID != 0 {
			mangaMap[tmp.NID] = append(mangaMap[tmp.NID], model.EntryScore{
				ID:    tmp.MID,
				Title: tmp.MTitle,
				Type:  constant.MangaType,
				Score: tmp.MScore,
			})
		}
	}

	for i, s := range scores {
		if len(animeMap[s.Novel[0].ID]) > 0 {
			scores[i].Anime = animeMap[s.Novel[0].ID]
		}
		if len(mangaMap[s.Novel[0].ID]) > 0 {
			scores[i].Manga = mangaMap[s.Novel[0].ID]
		}
	}

	// Prepare meta.
	meta := map[string]interface{}{
		"count": len(scores),
	}

	return scores, meta, http.StatusOK, nil
}
