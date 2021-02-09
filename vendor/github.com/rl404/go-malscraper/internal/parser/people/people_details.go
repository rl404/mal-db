package people

import (
	"html"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	strip "github.com/grokify/html-strip-tags-go"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type detail struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.People
}

// GetDetails to get people details.
func (p *parser) GetDetails(a *goquery.Selection) *model.People {
	d := detail{area: a, cleanImg: p.cleanImg}
	d.setID()
	d.setName()
	d.setImage()
	d.cleanBiodata()
	d.setBiodata()
	d.setMore()
	return &d.data
}

func (d *detail) setID() {
	area := d.area.Find("#horiznav_nav").Find("li").First()
	link, _ := area.Find("a").Attr("href")
	d.data.ID = utils.StrToNum(utils.GetValueFromSplit(link, "/", 4))
}

func (d *detail) setName() {
	d.data.Name = d.area.Find("h1.title-name").Text()
}

func (d *detail) setImage() {
	image, _ := d.area.Find("#content table tr td div a").Html()
	if image == "" {
		d.data.Image = ""
	} else {
		image, _ := d.area.Find("#content table tr td a img").Attr("data-src")
		d.data.Image = utils.URLCleaner(image, "image", d.cleanImg)
	}
}

func (d *detail) cleanBiodata() {
	area := d.area.Find("#content table tr td")
	area.Find("div").EachWithBreak(func(i int, trash *goquery.Selection) bool {
		if i < 5 {
			trash.Remove()
			return true
		}
		return false
	})
}

func (d *detail) setBiodata() {
	d.data.GivenName, _ = d.getBiodata("Given name")
	d.data.FamilyName, _ = d.getBiodata("Family name")
	_, d.data.AlternativeNames = d.getBiodata("Alternate names")
	d.data.Birthday = d.getBirthday()
	d.data.Website = d.getBioWeb()
	d.data.Favorite = d.getBioFavorite()
}

func (d *detail) getBiodata(t string) (string, []string) {
	area, _ := d.area.Find("#content table tr td").Html()

	r := regexp.MustCompile(`(` + t + `:<\/span>)[^<]*`)
	bioRegex := r.FindString(area)
	bioRegex = strip.StripTags(bioRegex)

	if bioRegex == "" {
		return "", []string{}
	}

	splitBio := strings.Split(bioRegex, ": ")
	splitBio[1] = strings.TrimSpace(splitBio[1])

	r = regexp.MustCompile(`\s+`)
	splitBio[1] = r.ReplaceAllString(splitBio[1], " ")

	if t == "Alternate names" {
		splitName := strings.Split(splitBio[1], ", ")
		if len(splitName) == 0 {
			return "", []string{}
		}
		for i := range splitName {
			splitName[i] = html.UnescapeString(splitName[i])
		}
		return "", splitName
	}

	if t == "Member Favorites" {
		splitBio[1] = strings.Replace(splitBio[1], ",", "", -1)
	}

	return html.UnescapeString(splitBio[1]), []string{}
}

func (d *detail) getBirthday() model.Date {
	date, _ := d.getBiodata("Birthday")
	y1, m1, d1 := utils.StrToDate(date)
	return model.Date{Year: y1, Month: m1, Day: d1}
}

func (d *detail) getBioWeb() string {
	area, _ := d.area.Find("#content table tr td").Html()

	r := regexp.MustCompile(`(Website:<\/span> <a)[^<]*`)
	bioRegex := r.FindString(html.UnescapeString(area))

	r = regexp.MustCompile(`".+"`)
	bioRegex = r.FindString(bioRegex)

	if bioRegex != "\"http://\"" {
		return strings.Replace(bioRegex, "\"", "", -1)
	}

	return ""
}

func (d *detail) getBioFavorite() int {
	area, _ := d.area.Find("#content table tr td").Html()

	r := regexp.MustCompile(`(Member Favorites:<\/span>)[^<]*`)
	bioRegex := r.FindString(area)
	bioRegex = strip.StripTags(bioRegex)

	splitBio := utils.GetValueFromSplit(bioRegex, ": ", 1)
	return utils.StrToNum(splitBio)
}

func (d *detail) setMore() {
	d.data.More = html.UnescapeString(d.area.Find("#content table tr td div[class^=people-informantion-more]").Text())
}
