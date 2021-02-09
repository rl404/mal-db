package search

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type club struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.ClubSearch
}

// GetClub to search club.
func (p *parser) GetClub(a *goquery.Selection) []model.ClubSearch {
	v := club{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (c *club) setDetail() {
	clubs := []model.ClubSearch{}
	c.area.Find("tr.table-data").Each(func(i int, eachClub *goquery.Selection) {
		clubs = append(clubs, model.ClubSearch{
			ID:      c.getID(eachClub),
			Name:    c.getName(eachClub),
			Image:   c.getImage(eachClub),
			Summary: c.getSummary(eachClub),
			Creator: c.getUsername(eachClub),
			Member:  c.getMember(eachClub),
		})
	})
	c.data = clubs
}

func (c *club) getID(eachClub *goquery.Selection) int {
	id, _ := eachClub.Find(".informantion").Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "cid=", 1))
}

func (c *club) getName(eachClub *goquery.Selection) string {
	return strings.TrimSpace(eachClub.Find(".informantion").Find("a").First().Text())
}

func (c *club) getImage(eachClub *goquery.Selection) string {
	image, _ := eachClub.Find(".picSurround img").Attr("data-src")
	return utils.URLCleaner(image, "image", c.cleanImg)
}

func (c *club) getUsername(eachClub *goquery.Selection) string {
	user, _ := eachClub.Find(".informantion a:nth-of-type(2)").Attr("href")
	return utils.GetValueFromSplit(user, "/", 4)
}

func (c *club) getSummary(eachClub *goquery.Selection) string {
	return eachClub.Find(".informantion div").Text()
}

func (c *club) getMember(eachClub *goquery.Selection) int {
	return utils.StrToNum(eachClub.Find("td:nth-of-type(2)").Text())
}
