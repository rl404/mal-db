package saver

import (
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model/raw"
)

func (a *API) insertEmptyID(t string, id int) error {
	err := a.db.Create(&raw.EmptyID{
		Type: t,
		ID:   id,
	}).Error
	if err != nil {
		return err
	}
	return a.cleanEmptyRelation(t, id)
}

func (a *API) deleteEmptyID(t string, id int) error {
	return a.db.Delete(&raw.EmptyID{
		Type: t,
		ID:   id,
	}).Error
}

func (a *API) cleanEmptyRelation(t string, id int) (err error) {
	switch t {
	case constant.AnimeType:
		if err = a.db.Where("anime_id = ?", id).Delete(raw.AnimeCharacter{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("anime_id = ?", id).Delete(raw.AnimeProducer{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("anime_id = ?", id).Delete(raw.AnimeStaff{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("id = ?", id).Delete(raw.Anime{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("media_id = ? and type = ?", id, t).Delete(raw.MediaGenre{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("(id = ? and type = ?) or (related_id = ? and related_type = ?)", id, t, id, t).Delete(raw.MediaRelated{}).Error; err != nil {
			return err
		}
	case constant.MangaType:
		if err = a.db.Where("manga_id = ?", id).Delete(raw.MangaCharacter{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("manga_id = ?", id).Delete(raw.MangaMagazine{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("id = ?", id).Delete(raw.Manga{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("media_id = ? and type = ?", id, t).Delete(raw.MediaGenre{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("(id = ? and type = ?) or (related_id = ? and related_type = ?)", id, t, id, t).Delete(raw.MediaRelated{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("manga_id = ?", id).Delete(raw.PeopleManga{}).Error; err != nil {
			return err
		}
	case constant.CharacterType:
		if err = a.db.Where("character_id = ?", id).Delete(raw.AnimeCharacter{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("id = ?", id).Delete(raw.Character{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("character_id = ?", id).Delete(raw.MangaCharacter{}).Error; err != nil {
			return err
		}
	case constant.PeopleType:
		if err = a.db.Where("people_id = ?", id).Delete(raw.AnimeCharacter{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("people_id = ?", id).Delete(raw.AnimeStaff{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("people_id = ?", id).Delete(raw.PeopleManga{}).Error; err != nil {
			return err
		}
		if err = a.db.Where("id = ?", id).Delete(raw.People{}).Error; err != nil {
			return err
		}
	}
	return nil
}
