package user

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type club struct {
	area *goquery.Selection
	data []model.Item
}

// GetClubs to get user club list.
func (p *parser) GetClubs(a *goquery.Selection) []model.Item {
	v := club{area: a}
	v.setDetails()
	return v.data
}

func (c *club) setDetails() {
	clubs := []model.Item{}
	c.area.Find("table td.pl8").Find("ol").Find("li").Each(func(i int, eachClub *goquery.Selection) {
		id, _ := eachClub.Find("a").Attr("href")
		clubs = append(clubs, model.Item{
			ID:   utils.StrToNum(utils.GetValueFromSplit(id, "cid=", 1)),
			Name: eachClub.Find("a").Text(),
		})
	})
	c.data = clubs
}
