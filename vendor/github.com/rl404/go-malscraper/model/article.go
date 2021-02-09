package model

import "time"

// Article represents main featured article model.
type Article struct {
	ID            int            `json:"id"`
	Title         string         `json:"title"`
	Summary       string         `json:"summary"`
	Content       string         `json:"content"`
	Date          time.Time      `json:"date"`
	Username      string         `json:"username"`
	View          int            `json:"view"`
	IsSpoiler     bool           `json:"is_spoiler"`
	IsAdvertorial bool           `json:"is_advertorial"`
	Tags          []string       `json:"tags"`
	Related       ArticleRelated `json:"related"`
}

// ArticleRelated represents featured article relation model.
type ArticleRelated struct {
	Anime     []Item `json:"anime"`
	Manga     []Item `json:"manga"`
	People    []Item `json:"people"`
	Character []Item `json:"character"`
}

// ArticleItem represents simple featured article model.
type ArticleItem struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Image         string   `json:"image"`
	Summary       string   `json:"summary"`
	Username      string   `json:"username"`
	View          int      `json:"view"`
	IsSpoiler     bool     `json:"is_spoiler"`
	IsAdvertorial bool     `json:"is_advertorial"`
	Tags          []string `json:"tags"`
}

// ArticleTagItem represents featured article tag.
type ArticleTagItem struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}
