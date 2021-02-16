package database

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/model/raw"
)

// SearchAnime to search anime.
func (d *Database) SearchAnime(query model.AnimeQuery) ([]model.Media, map[string]interface{}, int, error) {
	// Prepare query.
	baseQuery := d.db.Table(fmt.Sprintf("%s as a", raw.Anime{}.TableName()))
	if len(query.Title) >= 3 {
		columns := []string{"lower(title) like ?", "lower(title_english) like ?", "lower(title_japanese) like ?", "lower(title_synonym) like ?"}
		titleSplit := strings.Split(query.Title, " ")

		var whereConds []string
		var whereValues []interface{}

		for _, column := range columns {
			tmp := []string{}
			for _, title := range titleSplit {
				title = "%" + strings.ToLower(title) + "%"
				whereValues = append(whereValues, title)
				tmp = append(tmp, column)
			}
			whereConds = append(whereConds, "("+strings.Join(tmp, " and ")+")")
		}
		baseQuery = baseQuery.Where(strings.Join(whereConds, " or "), whereValues...)
	}
	if query.Score != 0 {
		baseQuery = baseQuery.Where("floor(score) = ?", query.Score)
	}
	if query.Type != 0 {
		baseQuery = baseQuery.Where("anime_type_id = ?", query.Type)
	}
	if query.Status != 0 {
		baseQuery = baseQuery.Where("anime_status_id = ?", query.Status)
	}
	if query.Rating != 0 {
		baseQuery = baseQuery.Where("anime_rating_id = ?", query.Rating)
	}
	if query.Source != 0 {
		baseQuery = baseQuery.Where("anime_source_id = ?", query.Source)
	}
	if query.Year != 0 {
		baseQuery = baseQuery.Where("start_year = ?", query.Year)
	}
	if query.Season != "" {
		if query.Year == 0 {
			query.Year = time.Now().Year()
		}

		seasonStart, seasonEnd := 0, 0
		switch query.Season {
		case "winter":
			seasonStart, seasonEnd = 1, 4
		case "spring":
			seasonStart, seasonEnd = 4, 7
		case "summer":
			seasonStart, seasonEnd = 7, 10
		case "fall":
			seasonStart, seasonEnd = 10, 13
		}

		baseQuery = baseQuery.Where("((premiered = '' and start_year = ? and start_month >= ? and start_month < ?) or (premiered != '' and split_part(premiered, ' ', 1) = ? and split_part(premiered, ' ', 2) = ?))", query.Year, seasonStart, seasonEnd, query.Season, strconv.Itoa(query.Year))
	}
	if query.StartYear != 0 || query.EndYear != 0 {
		if query.EndYear == 0 {
			query.EndYear = time.Now().Year()
		}
		if query.StartYear == 0 {
			query.StartYear = query.EndYear
		}
		baseQuery = baseQuery.Where("start_year >= ? and start_year <= ?", query.StartYear, query.EndYear)
	}
	if query.StartEpisode != 0 || query.EndEpisode != 0 {
		if query.StartEpisode == 0 && query.EndEpisode != 0 {
			baseQuery = baseQuery.Where("episode <= ?", query.EndEpisode)
		} else if query.StartEpisode != 0 && query.EndEpisode == 0 {
			baseQuery = baseQuery.Where("episode >= ?", query.StartEpisode)
		} else {
			baseQuery = baseQuery.Where("episode >= ? and episode <= ?", query.StartEpisode, query.EndEpisode)
		}
	}
	if query.StartDuration != 0 || query.EndDuration != 0 {
		if query.StartDuration == 0 && query.EndDuration != 0 {
			baseQuery = baseQuery.Where("duration <= ?", query.EndDuration*60)
		} else if query.StartDuration != 0 && query.EndDuration == 0 {
			baseQuery = baseQuery.Where("duration >= ?", query.StartDuration*60)
		} else {
			baseQuery = baseQuery.Where("duration >= ? and duration <= ?", query.StartDuration*60, query.EndDuration*60)
		}
	}
	if query.Order != "" {
		var orderList []string
		for _, order := range strings.Split(query.Order, ",") {
			order = strings.TrimSpace(order)
			sort := " asc"
			if order[0] == '-' {
				sort = " desc"
				order = order[1:]
			}
			orderList = append(orderList, order+sort)
		}
		query.Order = strings.Join(orderList, ",")
	} else {
		query.Order = "member desc"
	}
	if query.Producer != 0 {
		baseQuery = d.db.Table(fmt.Sprintf("%s as ap", raw.AnimeProducer{}.TableName())).
			Joins("left join (?) as a on a.id = ap.anime_id", baseQuery.Model(&raw.Anime{})).
			Where("ap.producer_id = ? and a.id is not null", query.Producer)
	}
	if len(query.Genre) > 0 {
		var genreConditions []string
		for _, g := range query.Genre {
			if g > 0 {
				genreConditions = append(genreConditions, fmt.Sprintf("%v = any(g.arr)", g))
			}
			if g < 0 {
				genreConditions = append(genreConditions, fmt.Sprintf("%v != all(g.arr)", g*-1))
			}
		}

		genreQuery := d.db.Model(&raw.MediaGenre{}).
			Select("media_id as anime_id, array_agg(genre_id) as arr").
			Where("type = ?", constant.AnimeType).
			Group("media_id")

		baseQuery = baseQuery.
			Joins("right join (?) as g on a.id = g.anime_id", genreQuery).
			Where(strings.Join(genreConditions, " and "))
	}

	var animeRaw []model.Media
	err := baseQuery.
		Select("a.id, a.title, a.image_url as image, a.score, a.voter, a.rank, a.popularity, a.member, a.favorite, a.anime_type_id as type, a.anime_status_id as status").
		Order(query.Order).
		Limit(query.Limit).
		Offset(query.Limit * (query.Page - 1)).
		Find(&animeRaw).Error
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	// Prepare meta.
	var count int64
	if err = baseQuery.Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	meta := map[string]interface{}{
		"count": count,
	}
	return animeRaw, meta, http.StatusOK, nil
}

// SearchManga to search manga.
func (d *Database) SearchManga(query model.MangaQuery) ([]model.Media, map[string]interface{}, int, error) {
	// Prepare query.
	baseQuery := d.db.Table(fmt.Sprintf("%s as m", raw.Manga{}.TableName()))
	if len(query.Title) >= 3 {
		columns := []string{"lower(title) like ?", "lower(title_english) like ?", "lower(title_japanese) like ?", "lower(title_synonym) like ?"}
		titleSplit := strings.Split(query.Title, " ")

		var whereConds []string
		var whereValues []interface{}

		for _, column := range columns {
			tmp := []string{}
			for _, title := range titleSplit {
				title = "%" + strings.ToLower(title) + "%"
				whereValues = append(whereValues, title)
				tmp = append(tmp, column)
			}
			whereConds = append(whereConds, "("+strings.Join(tmp, " and ")+")")
		}
		baseQuery = baseQuery.Where(strings.Join(whereConds, " or "), whereValues...)
	}
	if query.Score != 0 {
		baseQuery = baseQuery.Where("floor(score) = ?", query.Score)
	}
	if query.Type != 0 {
		baseQuery = baseQuery.Where("manga_type_id = ?", query.Type)
	}
	if query.Status != 0 {
		baseQuery = baseQuery.Where("manga_status_id = ?", query.Status)
	}
	if query.Year != 0 {
		baseQuery = baseQuery.Where("start_year = ?", query.Year)
	}
	if query.StartYear != 0 || query.EndYear != 0 {
		if query.EndYear == 0 {
			query.EndYear = time.Now().Year()
		}
		if query.StartYear == 0 {
			query.StartYear = query.EndYear
		}
		baseQuery = baseQuery.Where("start_year >= ? and start_year <= ?", query.StartYear, query.EndYear)
	}
	if query.StartChapter != 0 || query.EndChapter != 0 {
		if query.StartChapter == 0 && query.EndChapter != 0 {
			baseQuery = baseQuery.Where("chapter <= ?", query.EndChapter)
		} else if query.StartChapter != 0 && query.EndChapter == 0 {
			baseQuery = baseQuery.Where("chapter >= ?", query.StartChapter)
		} else {
			baseQuery = baseQuery.Where("chapter >= ? and chapter <= ?", query.StartChapter, query.EndChapter)
		}
	}
	if query.StartVolume != 0 || query.EndVolume != 0 {
		if query.StartVolume == 0 && query.EndVolume != 0 {
			baseQuery = baseQuery.Where("volume <= ?", query.EndVolume)
		} else if query.StartVolume != 0 && query.EndVolume == 0 {
			baseQuery = baseQuery.Where("volume >= ?", query.StartVolume)
		} else {
			baseQuery = baseQuery.Where("volume >= ? and volume <= ?", query.StartVolume, query.EndVolume)
		}
	}
	if query.Order != "" {
		var orderList []string
		for _, order := range strings.Split(query.Order, ",") {
			order = strings.TrimSpace(order)
			sort := " asc"
			if order[0] == '-' {
				sort = " desc"
				order = order[1:]
			}
			orderList = append(orderList, order+sort)
		}
		query.Order = strings.Join(orderList, ",")
	} else {
		query.Order = "member desc"
	}
	if query.Magazine != 0 {
		baseQuery = d.db.Table(fmt.Sprintf("%s as mm", raw.MangaMagazine{}.TableName())).
			Joins("left join (?) as m on m.id = mm.manga_id", baseQuery.Model(&raw.Manga{})).
			Where("mm.magazine_id = ? and m.id is not null", query.Magazine)
	}
	if len(query.Genre) > 0 {
		var genreConditions []string
		for _, g := range query.Genre {
			if g > 0 {
				genreConditions = append(genreConditions, fmt.Sprintf("%v = any(g.arr)", g))
			}
			if g < 0 {
				genreConditions = append(genreConditions, fmt.Sprintf("%v != all(g.arr)", g*-1))
			}
		}

		genreQuery := d.db.Model(&raw.MediaGenre{}).
			Select("media_id as manga_id, array_agg(genre_id) as arr").
			Where("type = ?", constant.MangaType).
			Group("media_id")

		baseQuery = baseQuery.
			Joins("right join (?) as g on m.id = g.manga_id", genreQuery).
			Where(strings.Join(genreConditions, " and "))
	}

	var mangaRaw []model.Media
	err := baseQuery.
		Select("m.id, m.title, m.image_url as image, m.score, m.voter, m.rank, m.popularity, m.member, m.favorite, m.manga_type_id as type, m.manga_status_id as status").
		Order(query.Order).
		Limit(query.Limit).
		Offset(query.Limit * (query.Page - 1)).
		Find(&mangaRaw).Error
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	// Prepare meta.
	var count int64
	if err = baseQuery.Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	meta := map[string]interface{}{
		"count": count,
	}
	return mangaRaw, meta, http.StatusOK, nil
}

// SearchCharacter to search character.
func (d *Database) SearchCharacter(query model.EntryQuery) ([]model.Entry, map[string]interface{}, int, error) {
	// Prepare query.
	baseQuery := d.db.Table(raw.Character{}.TableName())
	if len(query.Name) >= 3 {
		columns := []string{"lower(name) like ?", "lower(nickname) like ?", "lower(japanese_name) like ?"}
		querySplit := strings.Split(query.Name, " ")

		var whereConds []string
		var whereValues []interface{}

		for _, column := range columns {
			tmp := []string{}
			for _, name := range querySplit {
				name = "%" + strings.ToLower(name) + "%"
				whereValues = append(whereValues, name)
				tmp = append(tmp, column)
			}
			whereConds = append(whereConds, "("+strings.Join(tmp, " and ")+")")
		}
		baseQuery = baseQuery.Where(strings.Join(whereConds, " or "), whereValues...)
	}
	if query.Order != "" {
		var orderList []string
		for _, o := range strings.Split(query.Order, ",") {
			o = strings.TrimSpace(o)
			sort := " asc"
			if o[0] == '-' {
				sort = " desc"
				o = o[1:]
			}
			orderList = append(orderList, o+sort)
		}
		query.Order = strings.Join(orderList, ",")
	} else {
		query.Order = "favorite desc"
	}

	var charRaw []model.Entry
	err := baseQuery.
		Select("id, name, image_url as image, 'character' as type").
		Order(query.Order).
		Limit(query.Limit).
		Offset(query.Limit * (query.Page - 1)).
		Find(&charRaw).Error
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	// Prepare meta.
	var count int64
	if err = baseQuery.Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	meta := map[string]interface{}{
		"count": count,
	}
	return charRaw, meta, http.StatusOK, nil
}

// SearchPeople to search people.
func (d *Database) SearchPeople(query model.EntryQuery) ([]model.Entry, map[string]interface{}, int, error) {
	// Prepare query.
	baseQuery := d.db.Table(raw.People{}.TableName())
	if len(query.Name) >= 3 {
		columns := []string{"lower(name) like ?", "lower(given_name) like ?", "lower(family_name) like ?", "lower(alternative_name) like ?"}
		querySplit := strings.Split(query.Name, " ")

		var whereConds []string
		var whereValues []interface{}

		for _, column := range columns {
			tmp := []string{}
			for _, name := range querySplit {
				name = "%" + strings.ToLower(name) + "%"
				whereValues = append(whereValues, name)
				tmp = append(tmp, column)
			}
			whereConds = append(whereConds, "("+strings.Join(tmp, " and ")+")")
		}
		baseQuery = baseQuery.Where(strings.Join(whereConds, " or "), whereValues...)
	}
	if query.Order != "" {
		var orderList []string
		for _, o := range strings.Split(query.Order, ",") {
			o = strings.TrimSpace(o)
			sort := " asc"
			if o[0] == '-' {
				sort = " desc"
				o = o[1:]
			}
			orderList = append(orderList, o+sort)
		}
		query.Order = strings.Join(orderList, ",")
	} else {
		query.Order = "favorite desc"
	}

	var charRaw []model.Entry
	err := baseQuery.
		Select("id, name, image_url as image, 'people' as type").
		Order(query.Order).
		Limit(query.Limit).
		Offset(query.Limit * (query.Page - 1)).
		Find(&charRaw).Error
	if err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}

	// Prepare meta.
	var count int64
	if err = baseQuery.Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, nil, http.StatusInternalServerError, err
	}
	meta := map[string]interface{}{
		"count": count,
	}
	return charRaw, meta, http.StatusOK, nil
}
