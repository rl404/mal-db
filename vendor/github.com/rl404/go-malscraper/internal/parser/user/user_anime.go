package user

import (
	"fmt"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type anime struct {
	raws     []model.UserRawAnime
	cleanImg bool
	data     []model.UserAnime
}

// GetAnime to get user anime list.
func (p *parser) GetAnime(a []model.UserRawAnime) []model.UserAnime {
	v := anime{raws: a, cleanImg: p.cleanImg}
	v.setDetails()
	return v.data
}

func (a *anime) setDetails() {
	list := []model.UserAnime{}
	for _, r := range a.raws {
		tmp := model.UserAnime{
			ID:           r.AnimeID,
			Title:        fmt.Sprintf("%v", r.AnimeTitle),
			Image:        utils.URLCleaner(r.AnimeImage, "image", a.cleanImg),
			Score:        r.Score,
			Status:       r.Status,
			Type:         r.AnimeType,
			Progress:     r.WatchedEpisode,
			Episode:      r.AnimeEpisode,
			Tag:          fmt.Sprintf("%v", r.Tags),
			Rating:       r.AnimeRating,
			AiringStatus: r.AnimeAiringStatus,
			IsRewatching: a.getIsRewatching(r.IsRewatching),
			Days:         r.Days,
			Storage:      r.Storage,
			Priority:     r.Priority,
		}

		y1, m1, d1 := utils.StrToDate(r.AnimeStartDate)
		tmp.AiringStart = model.Date{Year: y1, Month: m1, Day: d1}

		y2, m2, d2 := utils.StrToDate(r.AnimeEndDate)
		tmp.AiringEnd = model.Date{Year: y2, Month: m2, Day: d2}

		t3, valid := utils.StrToTime(r.StartDate)
		if !valid {
			tmp.WatchStart = nil
		} else {
			tmp.WatchStart = &t3
		}

		t4, valid := utils.StrToTime(r.FinishDate)
		if !valid {
			tmp.WatchEnd = nil
		} else {
			tmp.WatchEnd = &t4
		}

		list = append(list, tmp)
	}
	a.data = list
}

func (a *anime) getIsRewatching(value interface{}) bool {
	if v, ok := value.(float64); ok {
		return v == 1
	}
	return false
}
