package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/internal"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/service"
)

// Validator will intercept and validate request data
// before processing the request.
type Validator struct {
	api    service.API
	cacher service.Cacher
	logger service.Logger
}

// New to create new validator.
func New(api service.API, c service.Cacher, l service.Logger) service.API {
	return &Validator{
		api:    api,
		cacher: c,
		logger: l,
	}
}

func (v *Validator) isEmptyID(key string) (empty bool) {
	v.logger.Trace("[%s] checking empty id...", key)
	if v.cacher.Get(key, &empty) == nil {
		v.logger.Debug("[%s] found empty id", key)
		return empty
	}
	return false
}

func (v *Validator) saveEmptyID(code int, key string) {
	if code != http.StatusNotFound {
		return
	}

	v.logger.Trace("[%s] saving empty id...", key)
	if err := v.cacher.Set(key, true); err != nil {
		v.logger.Error("[%s] failed saving cache: %s", key, err.Error())
	}
}

func (v *Validator) isArticleTagValid(tag string) bool {
	if tag == "" {
		return true
	}

	var tags []model.ArticleTagItem
	v.logger.Trace("[%s] checking valid article tag...", internal.KeyArticleTag)
	if v.cacher.Get(internal.KeyArticleTag, &tags) == nil {
		for _, t := range tags {
			if t.Tag == tag {
				return true
			}
		}
		return false
	}
	return true
}

func (v *Validator) isAnimeGenreValid(id int) bool {
	if id == 0 {
		return false
	}

	var genres []model.ItemCount
	v.logger.Trace("[%s] checking valid anime genre...", internal.GetKey(internal.KeyGenres, AnimeType))
	if v.cacher.Get(internal.GetKey(internal.KeyGenres, AnimeType), &genres) == nil {
		for _, g := range genres {
			if g.ID == id {
				return true
			}
		}
		return false
	}
	return true
}

func (v *Validator) isMangaGenreValid(id int) bool {
	if id == 0 {
		return false
	}

	var genres []model.ItemCount
	v.logger.Trace("[%s] checking valid manga genre...", internal.GetKey(internal.KeyGenres, MangaType))
	if v.cacher.Get(internal.GetKey(internal.KeyGenres, MangaType), &genres) == nil {
		for _, g := range genres {
			if g.ID == id {
				return true
			}
		}
		return false
	}
	return true
}

func (v *Validator) isNewsTagValid(tag string) bool {
	if tag == "" {
		return true
	}

	var tags model.NewsTag
	v.logger.Trace("[%s] checking valid news tag...", internal.KeyNewsTag)
	if v.cacher.Get(internal.KeyNewsTag, &tags) == nil {
		for _, t := range tags.Anime {
			if t.Tag == tag {
				return true
			}
		}
		for _, t := range tags.Manga {
			if t.Tag == tag {
				return true
			}
		}
		for _, t := range tags.People {
			if t.Tag == tag {
				return true
			}
		}
		for _, t := range tags.Music {
			if t.Tag == tag {
				return true
			}
		}
		for _, t := range tags.Event {
			if t.Tag == tag {
				return true
			}
		}
		for _, t := range tags.Industry {
			if t.Tag == tag {
				return true
			}
		}
		return false
	}
	return true
}

func (v *Validator) isProducerValid(id int) bool {
	if id < 0 {
		return false
	}

	var producers []model.ItemCount
	v.logger.Trace("[%s] checking valid producer...", internal.KeyProducers)
	if v.cacher.Get(internal.KeyProducers, &producers) == nil {
		for _, p := range producers {
			if p.ID == id {
				return true
			}
		}
		return false
	}
	return true
}

func (v *Validator) isMagazineValid(id int) bool {
	if id < 0 {
		return false
	}

	var magazines []model.ItemCount
	v.logger.Trace("[%s] checking valid magazine...", internal.KeyMagazines)
	if v.cacher.Get(internal.KeyMagazines, &magazines) == nil {
		for _, m := range magazines {
			if m.ID == id {
				return true
			}
		}
		return false
	}
	return true
}
