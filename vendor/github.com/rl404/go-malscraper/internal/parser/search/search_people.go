package search

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type people struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.PeopleSearch
}

// GetPeople to search people.
func (p *parser) GetPeople(a *goquery.Selection) []model.PeopleSearch {
	v := people{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (p *people) setDetail() {
	peopleList := []model.PeopleSearch{}
	p.area.Find("table").Find("tr").EachWithBreak(func(i int, eachSearch *goquery.Selection) bool {
		if i == 0 {
			return true
		}
		nameArea := eachSearch.Find("td:nth-of-type(2)")
		peopleList = append(peopleList, model.PeopleSearch{
			Image:    p.getImage(eachSearch),
			ID:       p.getID(nameArea),
			Name:     p.getName(nameArea),
			Nickname: p.getNickname(nameArea),
		})
		return true
	})
	p.data = peopleList
}

func (p *people) getImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td a img").Attr("data-src")
	return utils.URLCleaner(image, "image", p.cleanImg)
}

func (p *people) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 2))
}

func (p *people) getName(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

func (p *people) getNickname(nameArea *goquery.Selection) string {
	nick := nameArea.Find("small").First().Text()
	if nick != "" {
		nick = regexp.MustCompile(`\s+`).ReplaceAllString(nick, " ")
		return nick[1 : len(nick)-1]
	}
	return nick
}
