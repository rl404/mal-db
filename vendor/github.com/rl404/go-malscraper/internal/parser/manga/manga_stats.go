package manga

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

type stats struct {
	area *goquery.Selection
	data model.Stats
}

// GetStats to get manga stats.
func (p *parser) GetStats(a *goquery.Selection) *model.Stats {
	v := stats{area: a}
	v.setSummary()
	v.setScore()
	return &v.data
}

func (s *stats) setSummary() {
	area := s.area.Find("h2").First().Next()
	for goquery.NodeName(area) == "div" {

		summaryType := area.Find("span").Text()
		summaryType = strings.Replace(summaryType, ":", "", -1)
		summaryType = strings.ToLower(summaryType)

		area.Find("span").Remove()
		summaryValue := area.Text()

		switch strings.TrimSpace(summaryType) {
		case "reading":
			s.data.Summary.Current = utils.StrToNum(summaryValue)
		case "completed":
			s.data.Summary.Completed = utils.StrToNum(summaryValue)
		case "on-hold":
			s.data.Summary.OnHold = utils.StrToNum(summaryValue)
		case "dropped":
			s.data.Summary.Dropped = utils.StrToNum(summaryValue)
		case "plan to read":
			s.data.Summary.Planned = utils.StrToNum(summaryValue)
		case "total":
			s.data.Summary.Total = utils.StrToNum(summaryValue)
		}

		area = area.Next()
	}
}

func (s *stats) setScore() {
	s.area.Find("h2").First().Remove()
	area := s.area.Find("h2").First().Next()
	if goquery.NodeName(area) == "table" {
		area.Find("tr").Each(func(i int, eachScore *goquery.Selection) {
			switch s.getScoreType(eachScore) {
			case 1:
				s.data.Score.Score1.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score1.Percent = s.getScorePercent(eachScore)
			case 2:
				s.data.Score.Score2.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score2.Percent = s.getScorePercent(eachScore)
			case 3:
				s.data.Score.Score3.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score3.Percent = s.getScorePercent(eachScore)
			case 4:
				s.data.Score.Score4.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score4.Percent = s.getScorePercent(eachScore)
			case 5:
				s.data.Score.Score5.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score5.Percent = s.getScorePercent(eachScore)
			case 6:
				s.data.Score.Score6.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score6.Percent = s.getScorePercent(eachScore)
			case 7:
				s.data.Score.Score7.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score7.Percent = s.getScorePercent(eachScore)
			case 8:
				s.data.Score.Score8.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score8.Percent = s.getScorePercent(eachScore)
			case 9:
				s.data.Score.Score9.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score9.Percent = s.getScorePercent(eachScore)
			case 10:
				s.data.Score.Score10.Vote = s.getScoreVote(eachScore)
				s.data.Score.Score10.Percent = s.getScorePercent(eachScore)
			}
		})
	}
}

func (s *stats) getScoreType(eachScore *goquery.Selection) int {
	return utils.StrToNum(eachScore.Find("td").First().Text())
}

func (s *stats) getScoreVote(eachScore *goquery.Selection) int {
	vote := eachScore.Find("td:nth-of-type(2) span small").Text()
	vote = strings.Replace(vote, " votes", "", -1)
	return utils.StrToNum(vote[1 : len(vote)-1])
}

func (s *stats) getScorePercent(eachScore *goquery.Selection) float64 {
	eachScore.Find("td:nth-of-type(2) span small").Remove()
	percent := eachScore.Find("td:nth-of-type(2) span").Text()
	return utils.StrToFloat(strings.Replace(percent, "%", "", -1))
}
