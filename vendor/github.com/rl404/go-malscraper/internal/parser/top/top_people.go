package top

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type people struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.TopPeople
}

// GetPeople to get top people list.
func (p *parser) GetPeople(a *goquery.Selection) []model.TopPeople {
	v := people{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (p *people) setDetail() {
	topList := []model.TopPeople{}
	p.area.Find(".people-favorites-ranking-table").Find("tr.ranking-list").Each(func(i int, eachPeople *goquery.Selection) {
		nameArea := eachPeople.Find(".people")
		topList = append(topList, model.TopPeople{
			Rank:         p.getRank(eachPeople),
			ID:           p.getID(nameArea),
			Name:         p.getName(nameArea),
			JapaneseName: p.getJapaneseName(nameArea),
			Image:        p.getImage(nameArea),
			Birthday:     p.getBirthday(eachPeople),
			Favorite:     p.getFavorite(eachPeople),
		})
	})
	p.data = topList
}

func (p *people) getRank(eachPeople *goquery.Selection) int {
	return utils.StrToNum(eachPeople.Find("td").First().Find("span").Text())
}

func (p *people) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (p *people) getName(nameArea *goquery.Selection) string {
	return nameArea.Find(".information a").Text()
}

func (p *people) getJapaneseName(nameArea *goquery.Selection) string {
	japName := nameArea.Find(".information span").Text()
	if japName != "" {
		japName = japName[1 : len(japName)-1]
	}
	return japName
}

func (p *people) getImage(nameArea *goquery.Selection) string {
	image, _ := nameArea.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", p.cleanImg)
}

func (p *people) getBirthday(eachPeople *goquery.Selection) model.Date {
	day := eachPeople.Find(".birthday").Text()
	r := regexp.MustCompile(`\s+`)
	day = r.ReplaceAllString(strings.TrimSpace(day), " ")
	y, m, d := utils.StrToDate(day)
	return model.Date{Year: y, Month: m, Day: d}
}

func (p *people) getFavorite(eachPeople *goquery.Selection) int {
	return utils.StrToNum(eachPeople.Find(".favorites").Text())
}
