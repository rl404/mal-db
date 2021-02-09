package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"github.com/rl404/mal-db/internal/loader/api"
	"github.com/rl404/mal-db/internal/model"
	"github.com/rl404/mal-db/internal/pkg/utils"
)

type search struct {
	api api.API
}

func registerSearch(r chi.Router, api api.API) {
	s := search{api: api}
	r.Get("/search/anime", s.searchAnime)
	r.Get("/search/manga", s.searchManga)
	r.Get("/search/character", s.searchCharacter)
	r.Get("/search/people", s.searchPeople)
}

// @summary Search anime
// @tags search
// @accept json
// @produce json
// @param query query string false "Anime title"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @param score query integer false "Score" Enums(0,1,2,3,4,5,6,7,8,9,10)
// @param type query integer false "Anime type (1=TV, 2=OVA, 3=Movie, 4=Special, 5=ONA, 6=Music)" Enums(0,1,2,3,4,5,6)
// @param status query integer false "Anime airing status (1=airing, 2=finished, 3=upcoming)" Enums(0,1,2,3)
// @param rating query integer false "Anime rating (1=G, 2=PG, 3=PG13, 4=R17, 5=R, 6=RX)" Enums(0,1,2,3,4,5,6)
// @param source query integer false "Anime source (1=original, 2=manga, 3=4-koma, 4=web manga, 5=digital manga, 6=novel, 7=light novel, 8=visual novel, 9=game, 10=card game, 11=book, 12=picture book, 13=radio, 14=music)" Enums(1,2,3,4,5,6,7,8,9,10,11,12,13,14)
// @param year query integer false "Airing year"
// @param season query string false "Season" Enums(winter,spring,summer,fall)
// @param order query string false "Order (negative means descending)" Enums(member,-member,title,-title,score,-score)
// @param start_year query integer false "Start airing year"
// @param end_year query integer false "End airing year"
// @param start_episode query integer false "Minimum episode count"
// @param end_episode query integer false "Maximum episode count"
// @param start_duration query integer false "Minimum duration (in minutes)"
// @param end_duration query integer false "Maximum duration (in minutes)"
// @param producer query integer false "Producer ID"
// @param genre query []integer false "Genre ID (negative means to exclude the genre)"
// @success 200 {object} utils.Response{data=[]model.Media}
// @router /search/anime [get]
func (s *search) searchAnime(w http.ResponseWriter, r *http.Request) {
	var query model.AnimeQuery
	// Basic search.
	query.Title = r.URL.Query().Get("title")
	query.Score, _ = strconv.Atoi(r.URL.Query().Get("score"))
	query.Type, _ = strconv.Atoi(r.URL.Query().Get("type"))
	query.Status, _ = strconv.Atoi(r.URL.Query().Get("status"))
	query.Rating, _ = strconv.Atoi(r.URL.Query().Get("rating"))
	query.Source, _ = strconv.Atoi(r.URL.Query().Get("source"))
	query.Year, _ = strconv.Atoi(r.URL.Query().Get("year"))
	query.Season = r.URL.Query().Get("season")
	query.Limit, _ = strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Order = r.URL.Query().Get("order")

	// Advanced search.
	query.StartYear, _ = strconv.Atoi(r.URL.Query().Get("start_year"))
	query.EndYear, _ = strconv.Atoi(r.URL.Query().Get("end_year"))
	query.StartEpisode, _ = strconv.Atoi(r.URL.Query().Get("start_episode"))
	query.EndEpisode, _ = strconv.Atoi(r.URL.Query().Get("end_episode"))
	query.StartDuration, _ = strconv.Atoi(r.URL.Query().Get("start_duration"))
	query.EndDuration, _ = strconv.Atoi(r.URL.Query().Get("end_duration"))
	query.Producer, _ = strconv.Atoi(r.URL.Query().Get("producer"))

	genre := strings.Split(r.URL.Query().Get("genre"), ",")
	for _, g := range genre {
		gInt, _ := strconv.Atoi(g)
		if gInt != 0 {
			query.Genre = append(query.Genre, gInt)
		}
	}

	data, meta, code, err := s.api.SearchAnime(query)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Search manga
// @tags search
// @accept json
// @produce json
// @param query query string false "Manga title"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @param score query integer false "Score" Enums(0,1,2,3,4,5,6,7,8,9,10)
// @param type query integer false "Manga type (1=manga, 2=light novel, 3=one-shot, 4=doujinshi, 5=manhwa, 6=manhua, 7=oel, 8=novel)" Enums(0,1,2,3,4,5,6,7,8)
// @param status query integer false "Manga publishing status (1=publishing, 2=finished, 3=upcoming, 4=hiatus, 5=discontinued)" Enums(0,1,2,3,4,5)
// @param year query integer false "Publishing year"
// @param order query string false "Order (negative means descending)" Enums(member,-member,title,-title,score,-score)
// @param start_year query integer false "Start publishing year"
// @param end_year query integer false "End publishing year"
// @param start_chapter query integer false "Minimum chapter count"
// @param end_chapter query integer false "Maximum chapter count"
// @param start_volume query integer false "Minimum volume count"
// @param end_volume query integer false "Maximum volume count"
// @param magazine query integer false "Magazine ID"
// @param genre query []integer false "Genre ID (negative means to exclude the genre)"
// @success 200 {object} utils.Response{data=[]model.Media}
// @router /search/manga [get]
func (s *search) searchManga(w http.ResponseWriter, r *http.Request) {
	var query model.MangaQuery
	// Basic search.
	query.Title = r.URL.Query().Get("title")
	query.Score, _ = strconv.Atoi(r.URL.Query().Get("score"))
	query.Type, _ = strconv.Atoi(r.URL.Query().Get("type"))
	query.Status, _ = strconv.Atoi(r.URL.Query().Get("status"))
	query.Year, _ = strconv.Atoi(r.URL.Query().Get("year"))
	query.Limit, _ = strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Order = r.URL.Query().Get("order")

	// Advanced search.
	query.StartYear, _ = strconv.Atoi(r.URL.Query().Get("start_year"))
	query.EndYear, _ = strconv.Atoi(r.URL.Query().Get("end_year"))
	query.StartChapter, _ = strconv.Atoi(r.URL.Query().Get("start_chapter"))
	query.EndChapter, _ = strconv.Atoi(r.URL.Query().Get("end_chapter"))
	query.StartVolume, _ = strconv.Atoi(r.URL.Query().Get("start_volume"))
	query.EndVolume, _ = strconv.Atoi(r.URL.Query().Get("end_volume"))
	query.Magazine, _ = strconv.Atoi(r.URL.Query().Get("magazine"))

	genre := strings.Split(r.URL.Query().Get("genre"), ",")
	for _, g := range genre {
		gInt, _ := strconv.Atoi(g)
		if gInt != 0 {
			query.Genre = append(query.Genre, gInt)
		}
	}

	data, meta, code, err := s.api.SearchManga(query)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Search character
// @tags search
// @accept json
// @produce json
// @param query query string false "Character name"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @param order query string false "Order (negative means descending)" Enums(name,-name,favorite,-favorite)
// @success 200 {object} utils.Response{data=[]model.Entry}
// @router /search/character [get]
func (s *search) searchCharacter(w http.ResponseWriter, r *http.Request) {
	var query model.EntryQuery
	query.Name = r.URL.Query().Get("name")
	query.Limit, _ = strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Order = r.URL.Query().Get("order")
	data, meta, code, err := s.api.SearchCharacter(query)
	utils.ResponseWithJSON(w, code, data, err, meta)
}

// @summary Search people
// @tags search
// @accept json
// @produce json
// @param query query string false "People name"
// @param page query integer false "Page"
// @param limit query integer false "Limit"
// @param order query string false "Order (negative means descending)" Enums(name,-name,favorite,-favorite)
// @success 200 {object} utils.Response{data=[]model.Entry}
// @router /search/people [get]
func (s *search) searchPeople(w http.ResponseWriter, r *http.Request) {
	var query model.EntryQuery
	query.Name = r.URL.Query().Get("name")
	query.Limit, _ = strconv.Atoi(utils.GetQuery(r, "limit", "10"))
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Order = r.URL.Query().Get("order")
	data, meta, code, err := s.api.SearchPeople(query)
	utils.ResponseWithJSON(w, code, data, err, meta)
}
