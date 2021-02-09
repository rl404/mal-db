package model

import "time"

// Club represents main club model.
type Club struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Image       string      `json:"image"`
	Information string      `json:"information"`
	Category    string      `json:"category"`
	Type        string      `json:"type"`
	Member      int         `json:"member"`
	Picture     int         `json:"picture"`
	CreatedDate time.Time   `json:"createdDate"`
	Admins      []ClubAdmin `json:"admins"`
}

// ClubAdmin represents club admin and officer model.
type ClubAdmin struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// ClubItem represents simple club model.
type ClubItem struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Member int    `json:"member"`
}

// ClubMember represents club's member model.
type ClubMember struct {
	Username string `json:"username"`
	Image    string `json:"image"`
}

// ClubRelated represents club related model.
type ClubRelated struct {
	Anime     []Item `json:"anime"`
	Manga     []Item `json:"manga"`
	Character []Item `json:"character"`
}
