package review

import (
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type reviews struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.Review
}

// GetReviews to get anime/manga review list.
func (p *parser) GetReviews(a *goquery.Selection) []model.Review {
	v := reviews{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (r *reviews) setDetail() {
	reviews := []model.Review{}
	r.area.Find(".borderDark").Each(func(i int, area *goquery.Selection) {
		topArea := area.Find(".spaceit").First()
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()
		tmp := model.Review{
			ID:       r.getID(veryBottomArea),
			Source:   r.getSource(topArea, bottomArea),
			Username: r.getUsername(topArea),
			Image:    r.getImage(topArea),
			Helpful:  r.getHelpful(topArea),
			Date:     r.getDate(topArea),
			Score:    r.getScore(bottomArea),
			Review:   r.getReview(bottomArea),
		}
		if tmp.Source.Type == "anime" {
			tmp.Episode = r.getProgress(topArea)
		} else if tmp.Source.Type == "manga" {
			tmp.Chapter = r.getProgress(topArea)
		}
		reviews = append(reviews, tmp)
	})
	r.data = reviews
}

func (r *reviews) getID(veryBottomArea *goquery.Selection) int {
	id, _ := veryBottomArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "?id=", 1))
}

func (r *reviews) getSource(topArea *goquery.Selection, bottomArea *goquery.Selection) model.Source {
	sourceArea := topArea.Find(".mb8:nth-of-type(2)")
	return model.Source{
		ID:    r.getSourceID(sourceArea),
		Type:  r.getSourceType(sourceArea),
		Title: r.getSourceTitle(sourceArea),
		Image: r.getSourceImage(bottomArea),
	}
}

func (r *reviews) getSourceID(sourceArea *goquery.Selection) int {
	id, _ := sourceArea.Find("strong a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (r *reviews) getSourceType(sourceArea *goquery.Selection) string {
	typ := sourceArea.Find("small").First().Text()
	typ = strings.Replace(typ, "(", "", -1)
	typ = strings.Replace(typ, ")", "", -1)
	return strings.ToLower(typ)
}

func (r *reviews) getSourceTitle(sourceArea *goquery.Selection) string {
	return strings.TrimSpace(sourceArea.Find("strong").First().Text())
}

func (r *reviews) getSourceImage(bottomArea *goquery.Selection) string {
	image, _ := bottomArea.Find(".picSurround img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", r.cleanImg)
}

func (r *reviews) getUsername(topArea *goquery.Selection) string {
	return topArea.Find("table").First().Find("td:nth-of-type(2)").Find("a").First().Text()
}

func (r *reviews) getImage(topArea *goquery.Selection) string {
	image, _ := topArea.Find("table td img").First().Attr("src")
	return utils.URLCleaner(image, "image", r.cleanImg)
}

func (r *reviews) getHelpful(topArea *goquery.Selection) int {
	return utils.StrToNum(topArea.Find("table td:nth-of-type(2) strong").Text())
}

func (r *reviews) getDate(topArea *goquery.Selection) time.Time {
	area := topArea.Find("div").First().Find("div").First()
	date := area.Text()
	time, _ := area.Attr("title")
	t, _ := utils.StrToTime(date + " " + time)
	return t
}

func (r *reviews) getProgress(topArea *goquery.Selection) string {
	area := topArea.Find("div").First().Find("div:nth-of-type(2)").Text()
	value := strings.Replace(area, "episodes seen", "", -1)
	value = strings.Replace(value, "chapters read", "", -1)
	return strings.TrimSpace(value)
}

func (r *reviews) getScore(bottomArea *goquery.Selection) (score model.ReviewScore) {
	area := bottomArea.Find("table").First()
	area.Find("tr").Each(func(i int, eachScore *goquery.Selection) {
		scoreType := strings.ToLower(eachScore.Find("td:nth-of-type(1)").Text())
		scoreValue := utils.StrToNum(eachScore.Find("td:nth-of-type(2)").Text())
		switch scoreType {
		case "overall":
			score.Overall = scoreValue
		case "story":
			score.Story = scoreValue
		case "animation", "art":
			score.Art = scoreValue
		case "sound":
			score.Sound = scoreValue
		case "character":
			score.Character = scoreValue
		case "enjoyment":
			score.Enjoyment = scoreValue
		}
	})
	return score
}

func (r *reviews) getReview(bottomArea *goquery.Selection) string {
	bottomArea.Find("div").Remove()
	bottomArea.Find("a").Remove()

	rex := regexp.MustCompile(`[^\S\r\n]+`)
	reviewContent := rex.ReplaceAllString(bottomArea.Text(), " ")

	rex = regexp.MustCompile(`(\n\n \n)`)
	reviewContent = rex.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}
