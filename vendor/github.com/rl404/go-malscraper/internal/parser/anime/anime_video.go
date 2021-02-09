package anime

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type video struct {
	area       *goquery.Selection
	cleanVideo bool
	data       model.Video
}

// GetVideos to get anime video list.
func (p *parser) GetVideos(a *goquery.Selection) *model.Video {
	v := video{area: a, cleanVideo: p.cleanVid}
	v.setEpisode()
	v.setPromotion()
	return &v.data
}

func (v *video) setEpisode() {
	episodeList := []model.VideoEpisode{}
	v.area.Find(".episode-video").Find(".video-list-outer").Each(func(i int, eachEpisode *goquery.Selection) {
		linkArea := eachEpisode.Find("a").First()
		link, _ := linkArea.Attr("href")
		episodeList = append(episodeList, model.VideoEpisode{
			Title:   v.getEpisodeTitle(linkArea),
			Episode: v.getEpisodeNo(linkArea),
			Link:    link,
		})
	})
	v.data.Episodes = episodeList
}

func (v *video) getEpisodeTitle(linkArea *goquery.Selection) string {
	return linkArea.Find("span.title span").Text()
}

func (v *video) getEpisodeNo(linkArea *goquery.Selection) int {
	linkArea.Find("span.title").Find("span").Remove()
	episode := linkArea.Find("span.title").Text()
	episode = strings.Replace(episode, "Episode", "", -1)
	return utils.StrToNum(episode)
}

func (v *video) setPromotion() {
	promoList := []model.VideoPromo{}
	area := v.area.Find(".promotional-video")
	area.Find(".video-list-outer").Each(func(i int, eachPromo *goquery.Selection) {
		linkArea := eachPromo.Find("a").First()
		promoList = append(promoList, model.VideoPromo{
			Title: v.getPromoTitle(linkArea),
			Link:  v.getPromoLink(linkArea),
		})
	})
	v.data.Promotions = promoList
}

func (v *video) getPromoTitle(linkArea *goquery.Selection) string {
	return linkArea.Find("span.title").Text()
}

func (v *video) getPromoLink(linkArea *goquery.Selection) string {
	link, _ := linkArea.Attr("href")
	return utils.URLCleaner(link, "video", v.cleanVideo)
}
