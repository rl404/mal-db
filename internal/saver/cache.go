package saver

import "github.com/rl404/mal-db/internal/constant"

func (a *API) deleteCache(t string, id int) (err error) {
	keys := map[string][]string{
		constant.AnimeType: {
			constant.GetKey(constant.KeyAnime, id),
			constant.GetKey(constant.KeyAnimeCharacter, id),
			constant.GetKey(constant.KeyAnimeStaff, id),
			constant.GetKey(constant.KeyStats, constant.AnimeType, id),
		},
		constant.MangaType: {
			constant.GetKey(constant.KeyManga, id),
			constant.GetKey(constant.KeyMangaCharacter, id),
			constant.GetKey(constant.KeyStats, constant.MangaType, id),
		},
		constant.CharacterType: {
			constant.GetKey(constant.KeyCharacter, id),
			constant.GetKey(constant.KeyCharacterOgraphy, id, constant.AnimeType),
			constant.GetKey(constant.KeyCharacterOgraphy, id, constant.MangaType),
			constant.GetKey(constant.KeyCharacterVA, id),
		},
		constant.PeopleType: {
			constant.GetKey(constant.KeyPeople, id),
			constant.GetKey(constant.KeyPeopleVA, id),
			constant.GetKey(constant.KeyPeopleStaff, id),
			constant.GetKey(constant.KeyPeopleManga, id),
		},
	}

	for _, k := range keys[t] {
		if err = a.cacher.Delete(k); err != nil {
			return err
		}
	}

	return nil
}
