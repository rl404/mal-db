package user

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type favorite struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.UserFavorite
}

// GetFavorites to get user favorite list.
func (p *parser) GetFavorites(a *goquery.Selection) *model.UserFavorite {
	v := favorite{area: a, cleanImg: p.cleanImg}
	v.setDetails()
	return &v.data
}

func (f *favorite) setDetails() {
	favoriteArea := f.area.Find(".container-right .user-favorites-outer")
	f.data.Anime = f.getList(favoriteArea, "anime")
	f.data.Manga = f.getList(favoriteArea, "manga")
	f.data.Character = f.getList(favoriteArea, "characters")
	f.data.People = f.getList(favoriteArea, "people")
}

func (f *favorite) getList(favoriteArea *goquery.Selection, t string) []model.UserFavoriteItem {
	list := []model.UserFavoriteItem{}
	favoriteArea = favoriteArea.Find("ul.favorites-list." + t)
	if favoriteArea.Text() != "" {
		favoriteArea.Find("li").Each(func(i int, favorite *goquery.Selection) {
			list = append(list, model.UserFavoriteItem{
				ID:    f.getID(favorite),
				Name:  f.getName(favorite),
				Image: f.getImage(favorite),
			})
		})
	}
	return list
}

func (f *favorite) getID(favorite *goquery.Selection) int {
	href, _ := favorite.Find("a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(href, "/", 4))
}

func (f *favorite) getName(favorite *goquery.Selection) string {
	return favorite.Find(".data a").First().Text()
}

func (f *favorite) getImage(favorite *goquery.Selection) string {
	image, _ := favorite.Find("img").First().Attr("src")
	return utils.URLCleaner(image, "image", f.cleanImg)
}
