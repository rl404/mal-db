package character

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type club struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.ClubItem
}

// GetClubs to get character club list.
func (p *parser) GetClubs(a *goquery.Selection) []model.ClubItem {
	d := club{area: a, cleanImg: p.cleanImg}

	if !d.isValid() {
		return nil
	}

	d.setDetail()
	return d.data
}

func (c *club) isValid() bool {
	return c.area.Find(".badresult").Text() == ""
}

func (c *club) setDetail() {
	clubs := []model.ClubItem{}
	c.area.Find("table tr td:nth-of-type(2)").Find(".borderClass").Each(func(i int, eachClub *goquery.Selection) {
		clubs = append(clubs, model.ClubItem{
			ID:     c.getID(eachClub),
			Name:   c.getName(eachClub),
			Member: c.getMember(eachClub),
		})
	})
	c.data = clubs
}

func (c *club) getID(eachClub *goquery.Selection) int {
	link, _ := eachClub.Find("a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(link, "cid=", 1))
}

func (c *club) getName(eachClub *goquery.Selection) string {
	return eachClub.Find("a").Text()
}

func (c *club) getMember(eachClub *goquery.Selection) int {
	return utils.StrToNum(strings.Replace(eachClub.Find("small").Text(), "members", "", -1))
}
