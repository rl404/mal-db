package database

import (
	"fmt"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/join"
	"github.com/rl404/mal-db/internal/model/raw"
)

func (d *Database) initRelated() model.Related {
	return model.Related{
		Sequel:      []model.Source{},
		Prequel:     []model.Source{},
		AltSetting:  []model.Source{},
		AltVersion:  []model.Source{},
		SideStory:   []model.Source{},
		Summary:     []model.Source{},
		FullStory:   []model.Source{},
		ParentStory: []model.Source{},
		SpinOff:     []model.Source{},
		Adaptation:  []model.Source{},
		Character:   []model.Source{},
		Other:       []model.Source{},
	}
}

func (d *Database) getRelated(t string, id int) model.Related {
	var animeRelated, mangaRelated []join.MediaRelated
	err := d.db.Table(fmt.Sprintf("%s as mr", raw.MediaRelated{}.TableName())).
		Select("mr.related_id as id, a.title, a.image_url, mr.related_type_id, mr.related_type").
		Joins(fmt.Sprintf("left join %s as a on a.id = mr.related_id", raw.Anime{}.TableName())).
		Where("mr.media_id = ? and mr.media_type = ? and mr.related_type = ?", id, t, constant.AnimeType).
		Find(&animeRelated).Error
	if err != nil {
		d.log.Error(err.Error())
	}

	err = d.db.Table(fmt.Sprintf("%s as mr", raw.MediaRelated{}.TableName())).
		Select("mr.related_id as id, m.title, m.image_url, mr.related_type_id, mr.related_type").
		Joins(fmt.Sprintf("left join %s as m on m.id = mr.related_id", raw.Manga{}.TableName())).
		Where("mr.media_id = ? and mr.media_type = ? and mr.related_type = ?", id, t, constant.MangaType).
		Find(&mangaRelated).Error
	if err != nil {
		d.log.Error(err.Error())
	}

	result := d.initRelated()
	relatedRaw := append(animeRelated, mangaRelated...)
	if len(relatedRaw) == 0 {
		return result
	}

	for _, r := range relatedRaw {
		tmp := model.Source{
			ID:    r.ID,
			Title: r.Title,
			Image: r.ImageURL,
			Type:  r.RelatedType,
		}

		switch r.RelatedTypeID {
		case 1:
			result.Sequel = append(result.Sequel, tmp)
		case 2:
			result.Prequel = append(result.Prequel, tmp)
		case 3:
			result.AltSetting = append(result.AltSetting, tmp)
		case 4:
			result.AltVersion = append(result.AltVersion, tmp)
		case 6:
			result.SideStory = append(result.SideStory, tmp)
		case 7:
			result.Summary = append(result.Summary, tmp)
		case 8:
			result.FullStory = append(result.FullStory, tmp)
		case 9:
			result.ParentStory = append(result.ParentStory, tmp)
		case 10:
			result.SpinOff = append(result.SpinOff, tmp)
		case 11:
			result.Adaptation = append(result.Adaptation, tmp)
		case 12:
			result.Character = append(result.Character, tmp)
		case 13:
			result.Other = append(result.Other, tmp)
		default:
		}
	}

	return result
}
