package user

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type detail struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.User
}

// GetDetails to get user detail.
func (p *parser) GetDetails(a *goquery.Selection) *model.User {
	v := detail{area: a, cleanImg: p.cleanImg}
	v.setUsername()
	v.setImage()
	v.setInfo()
	v.setMoreInfo()
	v.setSns()
	v.setFriend()
	v.setAbout()
	return &v.data
}

func (d *detail) setUsername() {
	user := strings.TrimSpace(d.area.Parent().Find("h1.h1 span").Text())
	d.data.Username = user[:len(user)-10]
}

func (d *detail) setImage() {
	imageSrc, _ := d.area.Find(".container-left .user-profile .user-image img").Attr("data-src")
	d.data.Image = utils.URLCleaner(imageSrc, "image", d.cleanImg)
}

func (d *detail) setInfo() {
	infoList := make(map[string]string)
	infoArea := d.area.Find(".container-left .user-profile .user-status:nth-of-type(1)")
	infoArea.Find("li").Each(func(i int, info *goquery.Selection) {
		infoType := info.Find("span:nth-of-type(1)").Text()
		infoValue := info.Find("span:nth-of-type(2)").Text()
		infoList[infoType] = infoValue
	})

	d.data.LastOnline, _ = utils.StrToTime(infoList["Last Online"])
	d.data.Gender = infoList["Gender"]
	y1, m1, d1 := utils.StrToDate(infoList["Birthday"])
	d.data.Birthday = model.Date{Year: y1, Month: m1, Day: d1}
	d.data.Location = infoList["Location"]
	d.data.JoinedDate, _ = utils.StrToTime(infoList["Joined"])
}

func (d *detail) setMoreInfo() {
	infoArea := d.area.Find(".container-left .user-profile .user-status:nth-of-type(3)")
	d.data.ForumPost = utils.StrToNum(infoArea.Find("li:nth-of-type(1)").Find("span:nth-of-type(2)").Text())
	d.data.Review = utils.StrToNum(infoArea.Find("li:nth-of-type(2)").Find("span:nth-of-type(2)").Text())
	d.data.Recommendation = utils.StrToNum(infoArea.Find("li:nth-of-type(3)").Find("span:nth-of-type(2)").Text())
	d.data.BlogPost = utils.StrToNum(infoArea.Find("li:nth-of-type(4)").Find("span:nth-of-type(2)").Text())
	d.data.Club = utils.StrToNum(infoArea.Find("li:nth-of-type(5)").Find("span:nth-of-type(2)").Text())
}

func (d *detail) setSns() {
	snsArea := d.area.Find(".container-left .user-profile .user-profile-sns")
	snsArea.Find("a").Each(func(i int, sns *goquery.Selection) {
		snsClass, _ := sns.Attr("class")
		if snsClass != "di-ib mb8" {
			snsHref, _ := sns.Attr("href")
			d.data.Sns = append(d.data.Sns, snsHref)
		}
	})
}

func (d *detail) setFriend() {
	friendArea := d.area.Find(".container-left .user-profile .user-friends")
	friendCount := friendArea.Prev().Find("a").Text()

	r := regexp.MustCompile(`\(\d+\)`)
	friendCount = r.FindString(friendCount)

	replacer := strings.NewReplacer("(", "", ")", "")
	d.data.Friend = utils.StrToNum(replacer.Replace(friendCount))
}

func (d *detail) setAbout() {
	aboutArea := d.area.Find(".container-right table tr td div[class=word-break]")
	aboutContent, _ := aboutArea.Html()
	d.data.About = strings.TrimSpace(aboutContent)
}
