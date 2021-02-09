package user

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type stats struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.UserStats
}

// GetStats to get user stats detail.
func (p *parser) GetStats(a *goquery.Selection) *model.UserStats {
	v := stats{area: a, cleanImg: p.cleanImg}
	v.setAnimeStats()
	v.setMangaStats()
	return &v.data
}

func (s *stats) setAnimeStats() {
	statArea := s.area.Find(".container-right").Find(".user-statistics")
	statArea = statArea.Find("div[class=\"user-statistics-stats mt16\"]:nth-of-type(1)")
	scoreArea := statArea.Find(".stat-score")

	leftStatArea := statArea.Find("ul.stats-status")
	rightStatArea := statArea.Find("ul.stats-data")

	s.data.Anime.Days = s.getDaysScore(scoreArea, 1)
	s.data.Anime.MeanScore = s.getDaysScore(scoreArea, 2)
	s.data.Anime.Completed = s.getStatStatusCount(leftStatArea, 2, 1)
	s.data.Anime.OnHold = s.getStatStatusCount(leftStatArea, 3, 1)
	s.data.Anime.Dropped = s.getStatStatusCount(leftStatArea, 4, 1)
	s.data.Anime.Total = s.getStatStatusCount(rightStatArea, 1, 2)
	s.data.Anime.Current = s.getStatStatusCount(leftStatArea, 1, 1)
	s.data.Anime.Planned = s.getStatStatusCount(leftStatArea, 5, 1)
	s.data.Anime.Rewatched = s.getStatStatusCount(rightStatArea, 2, 2)
	s.data.Anime.Episode = s.getStatStatusCount(rightStatArea, 3, 2)
}

func (s *stats) setMangaStats() {
	statArea := s.area.Find(".container-right").Find(".user-statistics")
	statArea = statArea.Find("div[class=\"user-statistics-stats mt16\"]:nth-of-type(2)")
	scoreArea := statArea.Find(".stat-score")

	leftStatArea := statArea.Find("ul.stats-status")
	rightStatArea := statArea.Find("ul.stats-data")

	s.data.Manga.Days = s.getDaysScore(scoreArea, 1)
	s.data.Manga.MeanScore = s.getDaysScore(scoreArea, 2)
	s.data.Manga.Completed = s.getStatStatusCount(leftStatArea, 2, 1)
	s.data.Manga.OnHold = s.getStatStatusCount(leftStatArea, 3, 1)
	s.data.Manga.Dropped = s.getStatStatusCount(leftStatArea, 4, 1)
	s.data.Manga.Total = s.getStatStatusCount(rightStatArea, 1, 2)
	s.data.Manga.Current = s.getStatStatusCount(leftStatArea, 1, 1)
	s.data.Manga.Planned = s.getStatStatusCount(leftStatArea, 5, 1)
	s.data.Manga.Reread = s.getStatStatusCount(rightStatArea, 2, 2)
	s.data.Manga.Chapter = s.getStatStatusCount(rightStatArea, 3, 2)
	s.data.Manga.Volume = s.getStatStatusCount(rightStatArea, 4, 2)
}

func (s *stats) getDaysScore(scoreArea *goquery.Selection, nth int) float64 {
	area := scoreArea.Find("div:nth-of-type(" + strconv.Itoa(nth) + ")")
	area.Find("span").First().Remove()
	return utils.StrToFloat(area.Text())
}

func (s *stats) getStatStatusCount(statArea *goquery.Selection, liNo int, spanNo int) int {
	return utils.StrToNum(statArea.Find("li:nth-of-type(" + strconv.Itoa(liNo) + ") span:nth-of-type(" + strconv.Itoa(spanNo) + ")").Text())
}
