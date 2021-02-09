package user

import (
	"fmt"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type manga struct {
	raws     []model.UserRawManga
	cleanImg bool
	data     []model.UserManga
}

// GetManga to get user manga list.
func (p *parser) GetManga(a []model.UserRawManga) []model.UserManga {
	v := manga{raws: a, cleanImg: p.cleanImg}
	v.setDetails()
	return v.data
}

func (m *manga) setDetails() {
	list := []model.UserManga{}
	for _, r := range m.raws {
		tmp := model.UserManga{
			ID:               r.MangaID,
			Title:            fmt.Sprintf("%v", r.MangaTitle),
			Image:            r.MangaImage,
			Score:            r.Score,
			Status:           r.Status,
			Type:             r.MangaType,
			ChapterProgress:  r.ReadChapter,
			VolumeProgress:   r.ReadVolume,
			Chapter:          r.MangaChapter,
			Volume:           r.MangaVolume,
			Tag:              fmt.Sprintf("%v", r.Tags),
			PublishingStatus: r.MangaStatus,
			IsRereading:      m.getIsRereading(r.IsRereading),
			Days:             r.Days,
			Priority:         r.Priority,
		}

		y1, m1, d1 := utils.StrToDate(r.MangaStartDate)
		tmp.PublishingStart = model.Date{Year: y1, Month: m1, Day: d1}

		y2, m2, d2 := utils.StrToDate(r.MangaEndDate)
		tmp.PublishingEnd = model.Date{Year: y2, Month: m2, Day: d2}

		t3, valid := utils.StrToTime(r.StartDate)
		if !valid {
			tmp.ReadStart = nil
		} else {
			tmp.ReadStart = &t3
		}

		t4, valid := utils.StrToTime(r.FinishDate)
		if !valid {
			tmp.ReadEnd = nil
		} else {
			tmp.ReadEnd = &t4
		}

		list = append(list, tmp)
	}
	m.data = list
}

func (m *manga) getIsRereading(value interface{}) bool {
	if v, ok := value.(float64); ok {
		return v == 1
	}
	return false
}
