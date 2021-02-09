package anime

import (
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type review struct {
	area     *goquery.Selection
	cleanImg bool
	data     []model.Review
}

// GetReviews to get anime review list.
func (p *parser) GetReviews(a *goquery.Selection) []model.Review {
	v := review{area: a, cleanImg: p.cleanImg}
	v.setDetail()
	return v.data
}

func (r *review) setDetail() {
	reviews := []model.Review{}

	source := model.Source{
		ID:    r.getSourceID(),
		Title: r.getSourceTitle(),
		Image: r.getSourceImage(),
		Type:  "anime",
	}

	r.area.Find(".borderDark").Each(func(i int, review *goquery.Selection) {
		topArea := review.Find(".spaceit").First()
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()
		reviews = append(reviews, model.Review{
			ID:       r.getID(veryBottomArea),
			Username: r.getUsername(topArea),
			Image:    r.getImage(topArea),
			Source:   source,
			Helpful:  r.getHelpful(topArea),
			Date:     r.getDate(topArea),
			Episode:  r.getProgress(topArea),
			Score:    r.getScore(bottomArea),
			Review:   r.getReview(bottomArea),
		})
	})
	r.data = reviews
}

func (r *review) getID(veryBottomArea *goquery.Selection) int {
	id, _ := veryBottomArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "?id=", 1))
}

func (r *review) getUsername(topArea *goquery.Selection) string {
	return topArea.Find("table td:nth-of-type(2)").Find("a").First().Text()
}

func (r *review) getImage(topArea *goquery.Selection) string {
	image, _ := topArea.Find("table td:nth-of-type(1)").Find("img").Attr("src")
	return utils.URLCleaner(image, "image", r.cleanImg)
}

func (r *review) getSourceID() int {
	area := r.area.Find("#horiznav_nav").Find("li").First()
	link, _ := area.Find("a").Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(link, "/", 4))
}

func (r *review) getSourceImage() string {
	img, _ := r.area.Find("img.ac").Attr("data-src")
	return utils.URLCleaner(img, "img", r.cleanImg)
}

func (r *review) getSourceTitle() string {
	titleArea := r.area.Parent().Find("h1.title-name").First()
	titleArea.Find("span.title-english").Remove()
	return titleArea.Text()
}

func (r *review) getHelpful(topArea *goquery.Selection) int {
	helpful := topArea.Find("table td:nth-of-type(2) strong").First().Text()
	return utils.StrToNum(helpful)
}

func (r *review) getDate(topArea *goquery.Selection) time.Time {
	dateArea := topArea.Find("div").Find("div").First()
	date := dateArea.Text()
	time, _ := dateArea.Attr("title")
	result, _ := utils.StrToTime(date + " " + time)
	return result
}

func (r *review) getProgress(topArea *goquery.Selection) string {
	progress := topArea.Find("div").First().Find("div:nth-of-type(2)").Text()
	progress = strings.Replace(progress, "episodes seen", "", -1)
	progress = strings.Replace(progress, "chapters read", "", -1)
	return strings.TrimSpace(progress)
}

func (r *review) getScore(bottomArea *goquery.Selection) (rs model.ReviewScore) {
	bottomArea.Find("table").Find("tr").Each(func(i int, eachScore *goquery.Selection) {
		scoreType := strings.ToLower(eachScore.Find("td:nth-of-type(1)").Text())
		scoreValue := utils.StrToNum(eachScore.Find("td:nth-of-type(2)").Text())
		switch scoreType {
		case "overall":
			rs.Overall = scoreValue
		case "story":
			rs.Story = scoreValue
		case "animation":
			rs.Art = scoreValue
		case "sound":
			rs.Sound = scoreValue
		case "character":
			rs.Character = scoreValue
		case "enjoyment":
			rs.Enjoyment = scoreValue
		}
	})
	return rs
}

func (r *review) getReview(bottomArea *goquery.Selection) string {
	bottomArea.Find("a").Remove()
	bottomArea.Find("div[id^=score]").Remove()

	rex := regexp.MustCompile(`[^\S\r\n]+`)
	reviewContent := rex.ReplaceAllString(bottomArea.Text(), " ")

	rex = regexp.MustCompile(`(\n\n \n)`)
	reviewContent = rex.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}
