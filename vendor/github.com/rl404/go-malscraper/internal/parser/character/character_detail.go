package character

import (
	"html"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type detail struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.Character
}

// GetDetails to get character details.
func (p *parser) GetDetails(a *goquery.Selection) *model.Character {
	d := detail{area: a, cleanImg: p.cleanImg}

	if !d.isValid() {
		return nil
	}

	d.setID()
	d.setImage()
	d.setNickname()
	d.setJapaneseName()
	d.setName()
	d.setFavorite()
	d.setAbout()
	return &d.data
}

func (d detail) isValid() bool {
	return d.area.Find(".badresult").Text() == ""
}

func (d *detail) setID() {
	link, _ := d.area.Find("#horiznav_nav").Find("li").First().Find("a").Attr("href")
	d.data.ID = utils.StrToNum(utils.GetValueFromSplit(link, "/", 4))
}

func (d *detail) setImage() {
	image, _ := d.area.Find("#content table tr td div a").Html()
	if image == "" {
		d.data.Image = ""
	} else {
		image, _ := d.area.Find("#content table tr td div a img").Attr("data-src")
		d.data.Image = utils.URLCleaner(image, "image", d.cleanImg)
	}
}

func (d *detail) setNickname() {
	nickname := d.area.Find("h1")

	r := regexp.MustCompile(`\s+`)
	nickRegex := r.ReplaceAllString(nickname.Text(), " ")
	nickRegex = strings.TrimSpace(nickRegex)

	r = regexp.MustCompile(`\"([^"])*`)
	nickRegex = r.FindString(nickRegex)

	if nickRegex != "" {
		d.data.Nickname = nickRegex[1:]
	}
}

func (d *detail) setName() {
	area := d.area.Find("#content table tr td").Next()
	area = area.Find("h2.normal_header[style^=height]")

	r := regexp.MustCompile(`\s+`)
	name := r.ReplaceAllString(area.Text(), " ")

	area.Find("span").Remove()
	d.data.Name = strings.TrimSpace(name)
}

func (d *detail) setJapaneseName() {
	area := d.area.Find("#content table tr td").Next()
	area = area.Find("h2.normal_header small")

	r := regexp.MustCompile(`(\(|\))`)
	jpName := r.ReplaceAllString(area.Text(), "")

	d.data.JapaneseName = jpName
}

func (d *detail) setFavorite() {
	favorite := d.area.Find("#content table tr td").Text()

	r := regexp.MustCompile(`(Member Favorites: ).+`)
	regexFav := r.FindString(favorite)
	regexFav = strings.TrimSpace(regexFav)

	d.data.Favorite = utils.StrToNum(utils.GetValueFromSplit(regexFav, ": ", 1))
}

func (d *detail) setAbout() {
	aboutHTML, _ := d.area.Find("#content table tr td").Next().Html()

	r := regexp.MustCompile(`(?ms)(<h2 class="normal_header" style="height: 15px;">).*(<div class="normal_header">)`)
	regexAbout := r.FindString(aboutHTML)

	aboutGoQuery, _ := goquery.NewDocumentFromReader(strings.NewReader(regexAbout))
	aboutGoQuery.Find("h2.normal_header").Remove()
	cleanAbout := strings.TrimSpace(aboutGoQuery.Text())
	cleanAbout = strings.Replace(cleanAbout, "No biography written.", "", -1)
	d.data.About = strings.TrimSpace(html.UnescapeString(cleanAbout))
}
