package club

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type member struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.ClubMember
}

// GetMembers to get club member list.
func (p *parser) GetMembers(a *goquery.Selection) []model.ClubMember {
	v := member{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (m *member) setDetail() {
	users := []model.ClubMember{}
	m.area.Find("table").First().Find("td.borderClass").Each(func(i int, eachUser *goquery.Selection) {
		users = append(users, model.ClubMember{
			Username: m.getUsername(eachUser),
			Image:    m.getImage(eachUser),
		})
	})
	m.data = users
}

func (m *member) getUsername(eachUser *goquery.Selection) string {
	username, _ := eachUser.Find("a").First().Attr("href")
	return utils.GetValueFromSplit(username, "/", 2)
}

func (m *member) getImage(eachUser *goquery.Selection) string {
	image, _ := eachUser.Find(".picSurround img").Attr("data-src")
	if !strings.Contains(image, cdnMyAnimeListURL) {
		image = cdnMyAnimeListURL + image
	}
	return utils.URLCleaner(image, "image", m.cleanImg)
}
