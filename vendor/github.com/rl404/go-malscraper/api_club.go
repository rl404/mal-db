package malscraper

import (
	"github.com/rl404/go-malscraper/model"
)

// GetClubs to get club list.
//
// Example: https://myanimelist.net/clubs.php.
func (m *Malscraper) GetClubs(page ...int) ([]model.ClubSearch, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetClubs(p)
}

// GetClub to get club detail information.
//
// Example: https://myanimelist.net/clubs.php?cid=1.
func (m *Malscraper) GetClub(id int) (*model.Club, int, error) {
	return m.api.GetClub(id)
}

// GetClubMember to get club member list.
//
// Example: https://myanimelist.net/clubs.php?action=view&t=members&id=1.
func (m *Malscraper) GetClubMember(id int, page ...int) ([]model.ClubMember, int, error) {
	p := 1
	if len(page) > 0 {
		p = page[0]
	}
	return m.api.GetClubMember(id, p)
}

// GetClubPicture to get club picture list.
//
// Example: https://myanimelist.net/clubs.php?action=view&t=pictures&id=1.
func (m *Malscraper) GetClubPicture(id int) ([]string, int, error) {
	return m.api.GetClubPicture(id)
}

// GetClubRelated to get club related list.
//
// Example: https://myanimelist.net/clubs.php?cid=1.
func (m *Malscraper) GetClubRelated(id int) (*model.ClubRelated, int, error) {
	return m.api.GetClubRelated(id)
}
