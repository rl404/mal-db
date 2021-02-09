package anime

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type episode struct {
	area       *goquery.Selection
	cleanVideo bool
	data       []model.Episode
}

// GetEpisode to get anime episode list.
func (p *parser) GetEpisodes(a *goquery.Selection) []model.Episode {
	v := episode{area: a, cleanVideo: p.cleanVid}
	v.setDetail()
	return v.data
}

func (e *episode) setDetail() {
	epList := []model.Episode{}
	e.area.Find("table.episode_list").First().Find(".episode-list-data").Each(func(i int, eachEp *goquery.Selection) {
		epList = append(epList, model.Episode{
			Episode:       e.getEpisode(eachEp),
			Link:          e.getLink(eachEp),
			Title:         e.getTitle(eachEp),
			JapaneseTitle: e.getJapaneseTitle(eachEp),
			AiredDate:     e.getAired(eachEp),
			Tag:           e.getTag(eachEp),
		})
	})
	e.data = epList
}

func (e *episode) getEpisode(eachEp *goquery.Selection) int {
	return utils.StrToNum(eachEp.Find(".episode-number").Text())
}

func (e *episode) getLink(eachEp *goquery.Selection) string {
	link, _ := eachEp.Find(".episode-video a").First().Attr("href")
	return link
}

func (e *episode) getTitle(eachEp *goquery.Selection) string {
	return eachEp.Find(".episode-title").First().Find("a").First().Text()
}

func (e *episode) getJapaneseTitle(eachEp *goquery.Selection) string {
	return strings.TrimSpace(eachEp.Find(".episode-title span:last-child").Text())
}

func (e *episode) getAired(eachEp *goquery.Selection) *time.Time {
	aired := eachEp.Find(".episode-aired").Text()
	aired = strings.Replace(aired, "N/A", "", -1)
	at, valid := utils.StrToTime(aired)
	if !valid {
		return nil
	}
	return &at
}

func (e *episode) getTag(eachEp *goquery.Selection) string {
	return eachEp.Find("span.icon-episode-type-bg").Text()
}
