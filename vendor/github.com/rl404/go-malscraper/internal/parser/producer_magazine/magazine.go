package producermagazine

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type magazine struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.MangaItem
}

// GetManga to get magazine manga list.
func (p *parser) GetManga(a *goquery.Selection) []model.MangaItem {
	v := magazine{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (m *magazine) setDetail() {
	mangas := []model.MangaItem{}
	m.area.Find("div[class=\"seasonal-anime js-seasonal-anime\"]").Each(func(i int, eachArea *goquery.Selection) {
		nameArea := eachArea.Find("div.title")
		topArea := eachArea.Find("div.prodsrc")
		infoArea := eachArea.Find(".information")
		mangas = append(mangas, model.MangaItem{
			ID:             m.getID(nameArea),
			Image:          m.getImage(eachArea),
			Title:          m.getTitle(nameArea),
			Genres:         m.getGenres(eachArea),
			Synopsis:       m.getSynopsis(eachArea),
			Authors:        m.getAuthor(topArea),
			Volume:         m.getProgress(topArea),
			Serializations: m.getSerialization(eachArea),
			Type:           m.getType(topArea),
			StartDate:      m.getStartDate(infoArea),
			Member:         m.getMember(infoArea),
			Score:          m.getScore(infoArea),
		})
	})
	m.data = mangas
}

func (m *magazine) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("p a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (m *magazine) getImage(eachArea *goquery.Selection) string {
	image, _ := eachArea.Find("div.image img").Attr("data-src")
	return utils.URLCleaner(image, "image", m.cleanImg)
}

func (m *magazine) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("p a").Text()
}

func (m *magazine) getGenres(eachArea *goquery.Selection) []model.Item {
	genres := []model.Item{}
	genreArea := eachArea.Find("div[class=\"genres js-genre\"]")
	genreArea.Find("a").Each(func(i int, genre *goquery.Selection) {
		genreLink, _ := genre.Attr("href")
		genres = append(genres, model.Item{
			ID:   utils.StrToNum(utils.GetValueFromSplit(genreLink, "/", 3)),
			Name: genre.Text(),
		})
	})
	return genres
}

func (m *magazine) getSynopsis(eachArea *goquery.Selection) string {
	synopsis := strings.TrimSpace(eachArea.Find("div[class=\"synopsis js-synopsis\"] .preline").Text())
	if regexp.MustCompile(`No synopsis`).FindString(synopsis) != "" {
		return ""
	}
	return synopsis
}

func (m *magazine) getType(topArea *goquery.Selection) string {
	return strings.TrimSpace(topArea.Find("span.source").Text())
}

func (m *magazine) getAuthor(area *goquery.Selection) []model.Item {
	authors := []model.Item{}
	area.Find("span.producer").Find("a").Each(func(i int, each *goquery.Selection) {
		authors = append(authors, model.Item{
			ID:   m.getAuthorID(each),
			Name: m.getAuthorName(each),
		})
	})
	return authors
}

func (m *magazine) getAuthorID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(link, "/", 4))
}

func (m *magazine) getAuthorName(area *goquery.Selection) string {
	return area.Text()
}

func (m *magazine) getProgress(area *goquery.Selection) int {
	replacer := strings.NewReplacer("eps", "", "ep", "", "vols", "", "vol", "")
	return utils.StrToNum(replacer.Replace(area.Find("div.eps").Text()))
}

func (m *magazine) getSerialization(eachArea *goquery.Selection) []string {
	serialization := eachArea.Find("div[class=\"synopsis js-synopsis\"] .serialization a").Text()
	serial := utils.ArrayFilter(strings.Split(serialization, ","))
	if len(serial) == 0 {
		return []string{}
	}
	return serial
}

func (m *magazine) getStartDate(area *goquery.Selection) model.Date {
	y1, m1, d1 := utils.StrToDate(area.Find(".info .remain-time").Text())
	return model.Date{Year: y1, Month: m1, Day: d1}
}

func (m *magazine) getMember(area *goquery.Selection) int {
	return utils.StrToNum(area.Find(".scormem span[class^=member]").Text())
}

func (m *magazine) getScore(area *goquery.Selection) float64 {
	return utils.StrToFloat(area.Find(".scormem .score").Text())
}
