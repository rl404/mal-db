package manga

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

// GetClubs to get manga club list.
func (p *parser) GetClubs(a *goquery.Selection) []model.ClubItem {
	v := club{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (c *club) setDetail() {
	clubs := []model.ClubItem{}
	c.area.Find(".borderClass").Each(func(i int, eachClub *goquery.Selection) {
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
