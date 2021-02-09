package user

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

// GetReviews to get user review list.
func (p *parser) GetReviews(a *goquery.Selection) []model.Review {
	v := review{area: a, cleanImg: p.cleanImg}
	v.setDetails()
	return v.data
}

func (r *review) setDetails() {
	reviews := []model.Review{}
	username := r.getUsername()
	image := r.getImage()
	r.area.Find(".borderDark").Each(func(i int, area *goquery.Selection) {
		topArea := area.Find(".spaceit").First()
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()
		tmp := model.Review{
			ID:       r.getID(veryBottomArea),
			Username: username,
			Image:    image,
			Source:   r.getSource(topArea, bottomArea),
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

func (r *review) getUsername() string {
	user := strings.TrimSpace(r.area.Parent().Find("h1.h1 span").Text())
	return user[:len(user)-10]
}

func (r *review) getImage() string {
	imageSrc, _ := r.area.Find(".user-profile .user-image img").Attr("data-src")
	return utils.URLCleaner(imageSrc, "image", r.cleanImg)
}

func (r *review) getID(veryBottomArea *goquery.Selection) int {
	id, _ := veryBottomArea.Find("a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "?id=", 1))
}

func (r *review) getSource(topArea *goquery.Selection, bottomArea *goquery.Selection) model.Source {
	sourceArea := topArea.Find("div:nth-of-type(2)")
	return model.Source{
		ID:    r.getSourceID(sourceArea),
		Type:  r.getSourceType(sourceArea),
		Title: r.getSourceTitle(sourceArea),
		Image: r.getSourceImage(bottomArea),
	}
}

func (r *review) getSourceID(sourceArea *goquery.Selection) int {
	id, _ := sourceArea.Find("strong a").First().Attr("href")
	return utils.StrToNum(utils.GetValueFromSplit(id, "/", 4))
}

func (r *review) getSourceType(sourceArea *goquery.Selection) string {
	typ := sourceArea.Find("small").First().Text()
	typ = strings.Replace(typ, "(", "", -1)
	typ = strings.Replace(typ, ")", "", -1)
	return strings.ToLower(typ)
}

func (r *review) getSourceTitle(sourceArea *goquery.Selection) string {
	return strings.TrimSpace(sourceArea.Find("strong").First().Text())
}

func (r *review) getSourceImage(bottomArea *goquery.Selection) string {
	image, _ := bottomArea.Find(".picSurround img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", r.cleanImg)
}

func (r *review) getHelpful(topArea *goquery.Selection) int {
	return utils.StrToNum(topArea.Find("span[id^=\"rhelp\"]").Text())
}

func (r *review) getDate(topArea *goquery.Selection) time.Time {
	area := topArea.Find("div").First().Find("div").First()
	date := area.Text()
	time, _ := area.Attr("title")
	t, _ := utils.StrToTime(date + " " + time)
	return t
}

func (r *review) getProgress(topArea *goquery.Selection) string {
	area := topArea.Find("div").First().Find("div:nth-of-type(2)").Text()
	value := strings.Replace(area, "episodes seen", "", -1)
	value = strings.Replace(value, "chapters read", "", -1)
	return strings.TrimSpace(value)
}

func (r *review) getScore(bottomArea *goquery.Selection) (score model.ReviewScore) {
	bottomArea.Find("table").First().Find("tr").Each(func(i int, eachScore *goquery.Selection) {
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

func (r *review) getReview(bottomArea *goquery.Selection) string {
	bottomArea.Find("div").Remove()
	bottomArea.Find("a").Remove()

	rex := regexp.MustCompile(`[^\S\r\n]+`)
	reviewContent := rex.ReplaceAllString(bottomArea.Text(), " ")

	rex = regexp.MustCompile(`(\n\n \n)`)
	reviewContent = rex.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}
