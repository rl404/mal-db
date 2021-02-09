package review

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type detail struct {
	area     *goquery.Selection
	cleanImg bool
	data     model.Review
}

// GetDetails to get review details.
func (p *parser) GetDetails(a *goquery.Selection) *model.Review {
	d := detail{area: a, cleanImg: p.cleanImg}
	d.setID()
	d.setSource()
	d.setUsername()
	d.setImage()
	d.setHelpful()
	d.setDate()
	d.setProgress()
	d.setScore()
	d.setReview()
	return &d.data
}

func (d *detail) setID() {
	link, _ := d.area.Find(".breadcrumb .di-ib:nth-of-type(5) a").Attr("href")
	d.data.ID = utils.StrToNum(utils.GetValueFromSplit(link, "?id=", 1))
}

func (d *detail) setSource() {
	sourceArea := d.area.Find(".borderDark .spaceit")
	topArea := sourceArea.Find(".mb8")
	bottomArea := sourceArea.Next()
	d.data.Source = model.Source{
		ID:    d.getSourceID(topArea),
		Type:  d.getSourceType(topArea),
		Title: d.getSourceTitle(topArea),
		Image: d.getSourceImage(bottomArea),
	}
}

func (d *detail) getSourceID(topArea *goquery.Selection) int {
	id, _ := topArea.Find("strong a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (d *detail) getSourceType(topArea *goquery.Selection) string {
	typeStr := topArea.Find("small").First().Text()
	typeStr = strings.Replace(typeStr, "(", "", -1)
	typeStr = strings.Replace(typeStr, ")", "", -1)
	return strings.ToLower(typeStr)
}

func (d *detail) getSourceTitle(topArea *goquery.Selection) string {
	return strings.TrimSpace(topArea.Find("strong").Text())
}

func (d *detail) getSourceImage(bottomArea *goquery.Selection) string {
	image, _ := bottomArea.Find(".picSurround img").Attr("data-src")
	return utils.URLCleaner(image, "image", d.cleanImg)
}

func (d *detail) setUsername() {
	d.data.Username = d.area.Find(".borderDark .spaceit table td:nth-of-type(2) a").First().Text()
}

func (d *detail) setImage() {
	image, _ := d.area.Find(".borderDark .spaceit table td img").Attr("src")
	d.data.Image = utils.URLCleaner(image, "image", d.cleanImg)
}

func (d *detail) setHelpful() {
	d.data.Helpful = utils.StrToNum(d.area.Find(".borderDark .spaceit table td:nth-of-type(2) strong").First().Text())
}

func (d *detail) setDate() {
	dateArea := d.area.Find(".borderDark .spaceit div div").First()
	date := dateArea.Text()
	time, _ := dateArea.Attr("title")
	t, _ := utils.StrToTime(date + " " + time)
	d.data.Date = t
}

func (d *detail) setProgress() {
	progress := d.area.Find(".borderDark .spaceit div div:nth-of-type(2)").First().Text()
	progress = strings.Replace(progress, "episodes seen", "", -1)
	progress = strings.Replace(progress, "chapters read", "", -1)
	progress = strings.TrimSpace(progress)

	if d.data.Source.Type == "anime" {
		d.data.Episode = progress
	} else if d.data.Source.Type == "manga" {
		d.data.Chapter = progress
	}
}

func (d *detail) setScore() {
	scoreArea := d.area.Find(".borderDark .spaceit").Next().Find("table")
	scoreArea.Find("tr").Each(func(i int, score *goquery.Selection) {
		scoreType := strings.ToLower(score.Find("td").First().Text())
		scoreValue := utils.StrToNum(score.Find("td:nth-of-type(2)").Text())
		switch scoreType {
		case "overall":
			d.data.Score.Overall = scoreValue
		case "story":
			d.data.Score.Story = scoreValue
		case "animation", "art":
			d.data.Score.Art = scoreValue
		case "sound":
			d.data.Score.Sound = scoreValue
		case "character":
			d.data.Score.Character = scoreValue
		case "enjoyment":
			d.data.Score.Enjoyment = scoreValue
		}
	})
}

func (d *detail) setReview() {
	reviewArea := d.area.Find(".borderDark .spaceit").First().Next()
	reviewArea.Find("div").Remove()
	d.data.Review = strings.TrimSpace(reviewArea.Text())
}
