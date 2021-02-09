package club

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type related struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.ClubRelated
}

// GetRelated to get club related list.
func (p *parser) GetRelated(a *goquery.Selection) *model.ClubRelated {
	d := related{area: a, cleanImg: p.cleanImg}
	d.setDetail()
	return &d.data
}

func (r *related) setDetail() {
	r.area.Find("div.normal_header").Each(func(i int, header *goquery.Selection) {
		hText := strings.TrimSpace(header.Text())
		if hText == "Anime Relations" {
			r.data.Anime = r.getRelation(header)
		} else if hText == "Manga Relations" {
			r.data.Manga = r.getRelation(header)
		} else if hText == "Character Relations" {
			r.data.Character = r.getRelation(header)
		}
	})
}

func (r *related) getRelation(relArea *goquery.Selection) []model.Item {
	relations := []model.Item{}
	relArea = relArea.Next()
	for {
		className, _ := relArea.Attr("class")
		if className != "borderClass" {
			break
		}
		id, _ := relArea.Find("a").Attr("href")
		relations = append(relations, model.Item{
			ID:   utils.StrToNum(utils.GetValueFromSplit(id, "/", 1)),
			Name: relArea.Find("a").Text(),
		})
		relArea = relArea.Next()
	}
	return relations
}
