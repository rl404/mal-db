package database

import (
	"net/http"
	"time"

	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
)

// GetEntryCount to get all entry count.
func (d *Database) GetEntryCount() (*model.Total, map[string]interface{}, int, error) {
	var animeCount, mangaCount, charCount, peopleCount int64
	if err := d.db.Model(&raw.Anime{}).Count(&animeCount).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	if err := d.db.Model(&raw.Manga{}).Count(&mangaCount).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	if err := d.db.Model(&raw.Character{}).Count(&charCount).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	if err := d.db.Model(&raw.People{}).Count(&peopleCount).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	return &model.Total{
		Anime:     int(animeCount),
		Manga:     int(mangaCount),
		Character: int(charCount),
		People:    int(peopleCount),
	}, nil, http.StatusOK, nil
}

// GetYearSummary to get yearly anime & manga summary.
func (d *Database) GetYearSummary() ([]model.YearSummary, map[string]interface{}, int, error) {
	minYear := time.Now().Year()
	animeSummary, mangaSummary := make(map[int]model.YearSummaryDetail), make(map[int]model.YearSummaryDetail)

	rows, err := d.db.Model(&raw.Anime{}).
		Select("start_year as year, count(*) as count, round(avg(nullif(score,0)),2) as avg_score, min(nullif(score,0)) as min_score, max(score) as max_score").
		Where("start_year != ? and start_year <= ?", 0, time.Now().Year()).
		Group("start_year").
		Order("start_year asc").
		Rows()
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	for rows.Next() {
		var tmp model.YearSummaryDetail
		if err = d.db.ScanRows(rows, &tmp); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		animeSummary[tmp.Year] = tmp
		if tmp.Year < minYear {
			minYear = tmp.Year
		}
	}

	rows, err = d.db.Model(&raw.Manga{}).
		Select("start_year as year, count(*) as count, round(avg(nullif(score,0)),2) as avg_score, min(nullif(score,0)) as min_score, max(score) as max_score").
		Where("start_year != ? and start_year <= ?", 0, time.Now().Year()).
		Group("start_year").
		Order("start_year asc").
		Rows()
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	for rows.Next() {
		var tmp model.YearSummaryDetail
		if err = d.db.ScanRows(rows, &tmp); err != nil {
			return nil, nil, http.StatusInternalServerError, err
		}
		mangaSummary[tmp.Year] = tmp
		if tmp.Year < minYear {
			minYear = tmp.Year
		}
	}

	var data []model.YearSummary
	for y := minYear; y <= time.Now().Year(); y++ {
		a, m := animeSummary[y], mangaSummary[y]
		if a.Year == 0 {
			a.Year = y
		}
		if m.Year == 0 {
			m.Year = y
		}
		data = append(data, model.YearSummary{
			Anime: a,
			Manga: m,
		})
	}

	meta := map[string]interface{}{
		"count": len(data),
	}

	return data, meta, http.StatusOK, nil
}
