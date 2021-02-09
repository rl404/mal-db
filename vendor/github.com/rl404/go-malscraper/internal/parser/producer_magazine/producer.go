package producermagazine

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type producer struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.AnimeItem
}

// GetAnime to get producer anime list.
func (p *parser) GetAnime(a *goquery.Selection) []model.AnimeItem {
	v := producer{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (p *producer) setDetail() {
	producers := []model.AnimeItem{}
	p.area.Find("div[class=\"seasonal-anime js-seasonal-anime\"]").Each(func(i int, eachArea *goquery.Selection) {
		nameArea := eachArea.Find("div.title")
		topArea := eachArea.Find("div.prodsrc")
		infoArea := eachArea.Find(".information")
		producers = append(producers, model.AnimeItem{
			ID:        p.getID(nameArea),
			Image:     p.getImage(eachArea),
			Title:     p.getTitle(nameArea),
			Genres:    p.getGenres(eachArea),
			Synopsis:  p.getSynopsis(eachArea),
			Source:    p.getSource(topArea),
			Producers: p.getProducer(topArea),
			Episode:   p.getProgress(topArea),
			Licensors: p.getLicensors(eachArea),
			Type:      p.getType(infoArea),
			StartDate: p.getStartDate(infoArea),
			Member:    p.getMember(infoArea),
			Score:     p.getScore(infoArea),
		})
	})
	p.data = producers
}

func (p *producer) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("h2 a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (p *producer) getImage(eachArea *goquery.Selection) string {
	image, _ := eachArea.Find("div.image img").Attr("src")
	if image == "" {
		image, _ = eachArea.Find("div.image img").Attr("data-src")
	}
	return utils.URLCleaner(image, "image", p.cleanImg)
}

func (p *producer) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("h2 a").Text()
}

func (p *producer) getGenres(eachArea *goquery.Selection) []model.Item {
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

func (p *producer) getSynopsis(eachArea *goquery.Selection) string {
	synopsis := strings.TrimSpace(eachArea.Find("div[class=\"synopsis js-synopsis\"]").Text())
	if regexp.MustCompile(`No synopsis`).FindString(synopsis) != "" {
		return ""
	}
	return synopsis
}

func (p *producer) getSource(topArea *goquery.Selection) string {
	return strings.TrimSpace(topArea.Find("span.source").Text())
}

func (p *producer) getProducer(area *goquery.Selection) []model.Item {
	producers := []model.Item{}
	area.Find("span.producer").Find("a").Each(func(i int, each *goquery.Selection) {
		producers = append(producers, model.Item{
			ID:   p.getProducerID(each),
			Name: p.getProducerName(each),
		})
	})
	return producers
}

func (p *producer) getProducerID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(link, "/", 3))
}

func (p *producer) getProducerName(area *goquery.Selection) string {
	return area.Text()
}

func (p *producer) getProgress(area *goquery.Selection) int {
	replacer := strings.NewReplacer("eps", "", "ep", "", "vols", "", "vol", "")
	return utils.StrToNum(replacer.Replace(area.Find("div.eps").Text()))
}

func (p *producer) getLicensors(eachArea *goquery.Selection) []string {
	licensor, _ := eachArea.Find("div[class=\"synopsis js-synopsis\"] .licensors").Attr("data-licensors")
	return utils.ArrayFilter(strings.Split(licensor, ","))
}

func (p *producer) getType(area *goquery.Selection) string {
	return utils.GetValueFromSplit(area.Find(".info").Text(), "-", 0)
}

func (p *producer) getStartDate(area *goquery.Selection) model.Date {
	y, m, d := utils.StrToDate(area.Find(".info .remain-time").Text())
	return model.Date{Year: y, Month: m, Day: d}
}

func (p *producer) getMember(area *goquery.Selection) int {
	return utils.StrToNum(area.Find(".scormem span[class^=member]").Text())
}

func (p *producer) getScore(area *goquery.Selection) float64 {
	return utils.StrToFloat(area.Find(".scormem .score").Text())
}
