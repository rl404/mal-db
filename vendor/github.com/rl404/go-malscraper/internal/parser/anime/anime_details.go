package anime

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
	cleanVid bool
	data     model.Anime
}

// GetDetails to get anime details.
func (p *parser) GetDetails(a *goquery.Selection) *model.Anime {
	d := detail{area: a, cleanImg: p.cleanImg, cleanVid: p.cleanVid}
	d.setID()
	d.setImage()
	d.setTitle()
	d.setAltTitle()
	d.setVideo()
	d.setSynopsis()
	d.setScore()
	d.setVoter()
	d.setRank()
	d.setPopularity()
	d.setMember()
	d.setFavorite()
	d.setOtherInfo()
	d.setRelated()
	d.setSong()
	return &d.data
}

func (d *detail) setID() {
	area := d.area.Find("#horiznav_nav").Find("li").First()
	link, _ := area.Find("a").Attr("href")
	d.data.ID = utils.StrToNum(utils.GetValueFromSplit(link, "/", 4))
}

func (d *detail) setImage() {
	img, _ := d.area.Find("img.ac").Attr("data-src")
	d.data.Image = utils.URLCleaner(img, "img", d.cleanImg)
}

func (d *detail) setTitle() {
	titleArea := d.area.Parent().Find("h1.title-name").First()
	titleArea.Find("span.title-english").Remove()
	d.data.Title = titleArea.Text()
}

func (d *detail) setAltTitle() {
	area := d.area.Find("td.borderClass").First()
	d.data.AlternativeTitles.English = d.getAltTitle(area, "English")
	d.data.AlternativeTitles.Synonym = d.getAltTitle(area, "Synonyms")
	d.data.AlternativeTitles.Japanese = d.getAltTitle(area, "Japanese")
}

func (d *detail) getAltTitle(area *goquery.Selection, t string) string {
	r := regexp.MustCompile(`(` + t + `:</span>)([^<]*)`)
	altTitle, _ := area.Html()
	altTitle = r.FindString(altTitle)
	altTitle = strings.Replace(altTitle, t+":</span>", "", -1)
	return strings.TrimSpace(html.UnescapeString(altTitle))
}

func (d *detail) setVideo() {
	video, _ := d.area.Find(".video-promotion a").Attr("href")
	d.data.Video = utils.URLCleaner(video, "video", d.cleanVid)
}

func (d *detail) setSynopsis() {
	r := regexp.MustCompile(`\n[^\S\n]*`)
	synopsis := d.area.Find("p[itemprop=description]").Text()
	synopsis = r.ReplaceAllString(synopsis, "\n")
	if regexp.MustCompile(`No synopsis`).FindString(synopsis) != "" {
		d.data.Synopsis = ""
		return
	}
	d.data.Synopsis = strings.TrimSpace(synopsis)
}

func (d *detail) setScore() {
	d.data.Score = utils.StrToFloat(d.area.Find("div[class=\"fl-l score\"]").Text())
}

func (d *detail) setVoter() {
	replacer := strings.NewReplacer("users", "", "user", "", ",", "")
	voter, _ := d.area.Find("div[class=\"fl-l score\"]").Attr("data-user")
	d.data.Voter = utils.StrToNum(replacer.Replace(voter))
}

func (d *detail) setRank() {
	rank := d.area.Find("span[class=\"numbers ranked\"] strong").Text()
	d.data.Rank = utils.StrToNum(strings.Replace(rank, "#", "", -1))
}

func (d *detail) setPopularity() {
	popularity := d.area.Find("span[class=\"numbers popularity\"] strong").Text()
	d.data.Popularity = utils.StrToNum(strings.Replace(popularity, "#", "", -1))
}

func (d *detail) setMember() {
	d.data.Member = utils.StrToNum(d.area.Find("span[class=\"numbers members\"] strong").Text())
}

func (d *detail) setFavorite() {
	favoriteArea := d.area.Find("div[data-id=info2]").Next().Next().Next()
	favoriteArea.Find("span").Remove()
	d.data.Favorite = utils.StrToNum(favoriteArea.Text())
}

func (d *detail) setOtherInfo() {
	d.area.Find("td.borderClass").First().Find("h2").Each(func(i int, area *goquery.Selection) {
		if area.Text() == "Information" {
			area = area.Next()
			for {
				infoType := area.Find("span").First().Text()
				infoType = strings.ToLower(infoType)
				infoType = strings.Replace(infoType, ":", "", -1)
				value := d.getCleanInfo(area)

				switch infoType {
				case "type":
					d.data.Type = value
				case "episodes":
					d.data.Episode = utils.StrToNum(value)
				case "status":
					d.data.Status = value
				case "premiered":
					d.data.Premiered = value
				case "broadcast":
					d.data.Broadcast = value
				case "source":
					d.data.Source = value
				case "duration":
					d.data.Duration = value
				case "rating":
					d.data.Rating = value
				case "aired":
					d.data.AiringDate.Start, d.data.AiringDate.End = d.getAiringInfo(value)
				case "producers":
					d.data.Producers = d.getItemInfo(area, infoType, value)
				case "licensors":
					d.data.Licensors = d.getItemInfo(area, infoType, value)
				case "studios":
					d.data.Studios = d.getItemInfo(area, infoType, value)
				case "genres":
					d.data.Genres = d.getItemInfo(area, infoType, value)
				}

				area = area.Next()
				if goquery.NodeName(area) == "h2" || goquery.NodeName(area) == "br" {
					break
				}
			}
			return
		}
	})
}

func (d *detail) getCleanInfo(area *goquery.Selection) string {
	area.Find("span:nth-of-type(1)").Remove()
	replacer := strings.NewReplacer(", add some", "", "?", "", "Unknown", "", "per ep.", "")
	return strings.TrimSpace(replacer.Replace(area.Text()))
}

func (d *detail) getAiringInfo(value string) (t1 model.Date, t2 model.Date) {
	value = strings.Replace(value, "Not available", "", -1)
	if value == "" {
		return t1, t2
	}

	r := regexp.MustCompile(`\s+`)
	value = r.ReplaceAllString(value, " ")

	y1, m1, d1 := utils.StrToDate(utils.GetValueFromSplit(value, "to", 0))
	y2, m2, d2 := utils.StrToDate(utils.GetValueFromSplit(value, "to", 1))

	return model.Date{Year: y1, Month: m1, Day: d1}, model.Date{Year: y2, Month: m2, Day: d2}
}

func (d *detail) getItemInfo(infoArea *goquery.Selection, infoType string, value string) []model.Item {
	itemList := []model.Item{}

	value = strings.Replace(value, "None found", "", -1)
	if value == "" {
		return []model.Item{}
	}

	infoArea.Find("a").Each(func(i int, name *goquery.Selection) {
		link, _ := name.Attr("href")
		link = utils.GetValueFromSplit(link, "/", 3)
		itemList = append(itemList, model.Item{
			ID:   utils.StrToNum(link),
			Name: name.Text(),
		})
	})

	return itemList
}

func (d *detail) setRelated() {
	result := d.initRelated()
	d.area.Find(".anime_detail_related_anime").Find("tr").Each(func(i int, related *goquery.Selection) {
		relatedList := []model.RelatedItem{}
		related.Find("td:nth-of-type(2)").Find("a").EachWithBreak(func(i int, data *goquery.Selection) bool {
			relatedLink, _ := data.Attr("href")
			splitLink := strings.Split(relatedLink, "/")

			if data.Text() == "" || utils.StrToNum(splitLink[2]) == 0 {
				return true
			}

			relatedList = append(relatedList, model.RelatedItem{
				ID:    utils.StrToNum(splitLink[2]),
				Title: data.Text(),
				Type:  splitLink[1],
			})
			return true
		})

		relatedType := related.Find("td:nth-of-type(1)").Text()
		relatedType = strings.Replace(relatedType, ":", "", -1)
		relatedType = strings.TrimSpace(strings.ToLower(relatedType))

		switch relatedType {
		case "sequel":
			result.Sequel = relatedList
		case "prequel":
			result.Prequel = relatedList
		case "alternative setting":
			result.AltSetting = relatedList
		case "alternative version":
			result.AltVersion = relatedList
		case "side story":
			result.SideStory = relatedList
		case "summary":
			result.Summary = relatedList
		case "full story":
			result.FullStory = relatedList
		case "parent story":
			result.ParentStory = relatedList
		case "spin-off":
			result.SpinOff = relatedList
		case "adaptation":
			result.Adaptation = relatedList
		case "character":
			result.Character = relatedList
		case "other":
			result.Other = relatedList
		}
	})

	d.data.Related = result
}

func (d *detail) initRelated() model.Related {
	return model.Related{
		Sequel:      []model.RelatedItem{},
		Prequel:     []model.RelatedItem{},
		AltSetting:  []model.RelatedItem{},
		AltVersion:  []model.RelatedItem{},
		SideStory:   []model.RelatedItem{},
		Summary:     []model.RelatedItem{},
		FullStory:   []model.RelatedItem{},
		ParentStory: []model.RelatedItem{},
		SpinOff:     []model.RelatedItem{},
		Adaptation:  []model.RelatedItem{},
		Character:   []model.RelatedItem{},
		Other:       []model.RelatedItem{},
	}
}

func (d *detail) setSong() {
	d.data.Song.Opening = d.getCleanSong("div[class*=\"theme-songs opnening\"]")
	d.data.Song.Ending = d.getCleanSong("div[class*=\"theme-songs ending\"]")
}

func (d *detail) getCleanSong(div string) []string {
	songs := []string{}
	d.area.Find(div).Find("span.theme-song").Each(func(i int, eachSong *goquery.Selection) {
		r := regexp.MustCompile(`#\d*:\s`)
		song := eachSong.Text()
		song = r.ReplaceAllString(song, " ")
		songs = append(songs, strings.TrimSpace(song))
	})
	return songs
}
