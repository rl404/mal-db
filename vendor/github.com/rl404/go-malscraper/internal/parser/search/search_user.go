package search

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type user struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.UserSearch
}

// GetUser to search user.
func (p *parser) GetUser(a *goquery.Selection) []model.UserSearch {
	v := user{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (u *user) setDetail() {
	users := []model.UserSearch{}
	u.area.Find("td.borderClass").Each(func(i int, eachUser *goquery.Selection) {
		users = append(users, model.UserSearch{
			Username:   u.getName(eachUser),
			Image:      u.getImage(eachUser),
			LastOnline: u.getLastOnline(eachUser),
		})
	})
	u.data = users
}

func (u *user) getName(eachUser *goquery.Selection) string {
	return eachUser.Find("a").First().Text()
}

func (u *user) getImage(eachUser *goquery.Selection) string {
	image, _ := eachUser.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", u.cleanImg)
}

func (u *user) getLastOnline(eachUser *goquery.Selection) *time.Time {
	t, valid := utils.StrToTime(strings.TrimSpace(eachUser.Find("small").First().Text()))
	if !valid {
		return nil
	}
	return &t
}
