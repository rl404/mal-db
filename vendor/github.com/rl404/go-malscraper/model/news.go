package model

import "time"

// News represents main news model.
type News struct {
	ID       int         `json:"id"`
	Title    string      `json:"title"`
	Content  string      `json:"content"`
	Date     time.Time   `json:"date"`
	Username string      `json:"username"`
	ForumID  int         `json:"forum_id"`
	Comment  int         `json:"comment"`
	Tags     []string    `json:"tags"`
	Related  NewsRelated `json:"related"`
}

// NewsRelated represents news relation.
type NewsRelated struct {
	Anime  []Item `json:"anime"`
	Manga  []Item `json:"manga"`
	People []Item `json:"people"`
}

// NewsItem represents simple news model.
type NewsItem struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Image    string    `json:"image"`
	Content  string    `json:"content"`
	Date     time.Time `json:"date"`
	Username string    `json:"username"`
	ForumID  int       `json:"forum_id"`
	Comment  int       `json:"comment"`
}

// NewsTag represents news tag categories.
type NewsTag struct {
	Anime    []NewsTagItem `json:"anime"`
	Manga    []NewsTagItem `json:"manga"`
	People   []NewsTagItem `json:"people"`
	Music    []NewsTagItem `json:"music"`
	Event    []NewsTagItem `json:"event"`
	Industry []NewsTagItem `json:"industry"`
}

// NewsTagItem represents news tag detail.
type NewsTagItem struct {
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	Description string `json:"description"`
}
