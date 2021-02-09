package malscraper

import (
	"time"

	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/go-malscraper/pkg/utils"
)

// GetSeason to get seasonal anime list.
//
// Season should be one of these constants.
//
//  Winter
//  Spring
//  Summer
//  Fall
//
// Example: https://myanimelist.net/anime/season.
func (m *Malscraper) GetSeason(seasonYear ...interface{}) ([]model.AnimeItem, int, error) {
	season, year := utils.GetCurrentSeason(), time.Now().Year()
	for i, param := range seasonYear {
		switch i {
		case 0:
			if v, ok := param.(string); ok {
				season = v
			}
		case 1:
			if v, ok := param.(int); ok {
				year = v
			}
		}
	}
	return m.api.GetSeason(season, year)
}
