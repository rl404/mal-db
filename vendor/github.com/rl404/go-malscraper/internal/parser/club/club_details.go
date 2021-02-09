package club

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type detail struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.Club
}

// GetDetails to get club details.
func (p *parser) GetDetails(a *goquery.Selection) *model.Club {
	d := detail{area: a, cleanImg: p.cleanImg}
	d.setID()
	d.setName()
	d.setImage()
	d.setInformation()
	d.setMember()
	d.setPicture()
	d.setCategory()
	d.setDate()
	d.setDetail()
	d.setType()
	return &d.data
}

func (d *detail) setID() {
	area := d.area.Find("#content table tr").Find("td.borderClass").First()
	area.Find("div.normal_header").EachWithBreak(func(i int, header *goquery.Selection) bool {
		link, _ := header.Find("a").Attr("href")
		header.Find("div").Remove()
		hText := strings.TrimSpace(header.Text())
		if hText == "Club Members" {
			u, _ := url.Parse(link)
			q, _ := url.ParseQuery(u.RawQuery)
			d.data.ID = utils.StrToNum(q.Get("id"))
			return false
		}
		return true
	})
}

func (d *detail) setName() {
	d.data.Name = strings.TrimSpace(d.area.Find("h1").First().Text())
}

func (d *detail) setImage() {
	area := d.area.Find("#content table tr").Find("td.borderClass").First().Next()
	image, _ := area.Find("img").First().Attr("data-src")
	d.data.Image = utils.URLCleaner(image, "image", d.cleanImg)
}

func (d *detail) setInformation() {
	area := d.area.Find("#content table tr").Find("td.borderClass").First()
	d.data.Information, _ = area.Find("div.normal_header").First().Next().Html()
}

func (d *detail) setMember() {
	countArea := d.area.Find("#content table tr").Find("td.borderClass").First().Next()
	countArea = countArea.Find(".normal_header").First().Next()
	countArea.Find("span").Remove()
	d.data.Member = utils.StrToNum(countArea.Text())
}

func (d *detail) setPicture() {
	countArea := d.area.Find("#content table tr").Find("td.borderClass").First().Next()
	countArea = countArea.Find(".normal_header").First().Next().Next()
	countArea.Find("span").Remove()
	d.data.Picture = utils.StrToNum(countArea.Text())
}

func (d *detail) setCategory() {
	catArea := d.area.Find("#content table tr").Find("td.borderClass").First().Next()
	catArea = catArea.Find(".normal_header").First().Next().Next().Next()
	catArea.Find("span").Remove()
	d.data.Category = strings.TrimSpace(catArea.Text())
}

func (d *detail) setDate() {
	dateArea := d.area.Find("#content table tr").Find("td.borderClass").First().Next()
	dateArea = dateArea.Find(".normal_header").First().Next().Next().Next().Next()
	dateArea.Find("span").Remove()
	t, _ := utils.StrToTime(dateArea.Text())
	d.data.CreatedDate = t
}

func (d *detail) setDetail() {
	d.area.Find("div.normal_header").EachWithBreak(func(i int, header *goquery.Selection) bool {
		if strings.TrimSpace(header.Text()) == "Club Admins and Officers" {
			d.data.Admins = d.getAdmin(header)
			return false
		}
		return true
	})
}

func (d *detail) getAdmin(adminArea *goquery.Selection) []model.ClubAdmin {
	admins := []model.ClubAdmin{}
	adminArea = adminArea.Next()
	for {
		className, _ := adminArea.Attr("class")
		if className != "borderClass" {
			break
		}
		admins = append(admins, model.ClubAdmin{
			Username: d.getAdminUser(adminArea),
			Roles:    d.getAdminRole(adminArea),
		})
		adminArea = adminArea.Next()
	}
	return admins
}

func (d *detail) getAdminUser(adminArea *goquery.Selection) string {
	user, _ := adminArea.Find("a").Attr("href")
	return utils.GetValueFromSplit(user, "/", 2)
}

func (d *detail) getAdminRole(adminArea *goquery.Selection) []string {
	adminArea.Find("a").Remove()
	role := strings.TrimSpace(adminArea.Text())
	return strings.Split(role[1:len(role)-1], ", ")
}

func (d *detail) setType() {
	r := regexp.MustCompile(`This is a \w+ club\.`)
	clubTypeArea, _ := d.area.Find("#content table tr").Find("td.borderClass").First().Next().Html()
	clubType := r.FindString(clubTypeArea)
	clubType = strings.Replace(clubType, "This is a ", "", -1)
	clubType = strings.Replace(clubType, " club.", "", -1)
	d.data.Type = clubType
}
