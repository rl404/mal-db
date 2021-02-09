package model

import "time"

// User represents main user model.
type User struct {
	Username       string    `json:"username"`
	Image          string    `json:"image"`
	LastOnline     time.Time `json:"lastOnline"`
	Gender         string    `json:"gender"`
	Birthday       Date      `json:"birthday"`
	Location       string    `json:"location"`
	JoinedDate     time.Time `json:"joinedDate"`
	ForumPost      int       `json:"forumPost"`
	Review         int       `json:"review"`
	Recommendation int       `json:"recommendation"`
	BlogPost       int       `json:"blogPost"`
	Club           int       `json:"club"`
	Friend         int       `json:"friend"`
	Sns            []string  `json:"sns"`
	About          string    `json:"about"`
}

// UserStats represents anime & manga user stats model.
type UserStats struct {
	Anime UserAnimeStats `json:"anime"`
	Manga UserMangaStats `json:"manga"`
}

// UserAnimeStats represents anime user stats model.
type UserAnimeStats struct {
	Days      float64 `json:"days"`
	MeanScore float64 `json:"meanScore"`
	Current   int     `json:"current"`
	Completed int     `json:"completed"`
	OnHold    int     `json:"onHold"`
	Dropped   int     `json:"dropped"`
	Planned   int     `json:"planned"`
	Total     int     `json:"total"`
	Rewatched int     `json:"rewatched"`
	Episode   int     `json:"episode"`
}

// UserMangaStats represents manga user stats model.
type UserMangaStats struct {
	Days      float64 `json:"days"`
	MeanScore float64 `json:"meanScore"`
	Current   int     `json:"current"`
	Completed int     `json:"completed"`
	OnHold    int     `json:"onHold"`
	Dropped   int     `json:"dropped"`
	Planned   int     `json:"planned"`
	Total     int     `json:"total"`
	Reread    int     `json:"reread"`
	Chapter   int     `json:"chapter"`
	Volume    int     `json:"volume"`
}

// UserFavorite represents user favorite model.
type UserFavorite struct {
	Anime     []UserFavoriteItem `json:"anime"`
	Manga     []UserFavoriteItem `json:"manga"`
	Character []UserFavoriteItem `json:"character"`
	People    []UserFavoriteItem `json:"people"`
}

// UserFavoriteItem represents each user favorite entry.
type UserFavoriteItem struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// UserFriend represents user friend model.
type UserFriend struct {
	Username    string    `json:"username"`
	Image       string    `json:"image"`
	LastOnline  time.Time `json:"lastOnline"`
	FriendSince time.Time `json:"friendSince"`
}

// UserHistory represents user anime/manga history model.
type UserHistory struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Type     string    `json:"type"`
	Progress int       `json:"progress"`
	Date     time.Time `json:"date"`
}

// UserListQuery represents query model for user anime & manga list.
type UserListQuery struct {
	Username string
	Page     int
	Status   int
	Order    int
	Tag      string
}

// UserRawAnime represents MyAnimeList raw JSON anime list response.
type UserRawAnime struct {
	Status            int         `json:"status"`
	Score             int         `json:"score"`
	Tags              interface{} `json:"tags"`
	IsRewatching      interface{} `json:"is_rewatching"`
	WatchedEpisode    int         `json:"num_watched_episodes"`
	AnimeTitle        interface{} `json:"anime_title"`
	AnimeEpisode      int         `json:"anime_num_episodes"`
	AnimeAiringStatus int         `json:"anime_airing_status"`
	AnimeID           int         `json:"anime_id"`
	AnimeStudio       []Item      `json:"anime_studios"`
	AnimeLicensor     []Item      `json:"anime_licensors"`
	AnimeSeason       Season      `json:"anime_season"`
	HasEpisodeVideo   bool        `json:"has_episode_video"`
	HasPromotionVideo bool        `json:"has_promotion_video"`
	HasVideo          bool        `json:"has_video"`
	VideoURL          string      `json:"video_url"`
	AnimeURL          string      `json:"anime_url"`
	AnimeImage        string      `json:"anime_image_path"`
	AnimeType         string      `json:"anime_media_type_string"`
	AnimeRating       string      `json:"anime_mpaa_rating_string"`
	AnimeStartDate    string      `json:"anime_start_date_string"`
	AnimeEndDate      string      `json:"anime_end_date_string"`
	IsAddedToList     bool        `json:"is_added_to_list"`
	StartDate         string      `json:"start_date_string"`
	FinishDate        string      `json:"finish_date_string"`
	Days              int         `json:"days_string"`
	Storage           string      `json:"storage_string"`
	Priority          string      `json:"priority_string"`
}

// Season represents anime season model.
type Season struct {
	Year   int    `json:"year"`
	Season string `json:"season"`
}

// UserAnime represents user anime model.
type UserAnime struct {
	ID           int        `json:"id"`
	Title        string     `json:"title"`
	Image        string     `json:"image"`
	Score        int        `json:"score"`
	Status       int        `json:"status"`
	Type         string     `json:"type"`
	Progress     int        `json:"progress"`
	Episode      int        `json:"episode"`
	Tag          string     `json:"tag"`
	Rating       string     `json:"rating"`
	AiringStatus int        `json:"airingStatus"`
	AiringStart  Date       `json:"airingStart"`
	AiringEnd    Date       `json:"airingEnd"`
	WatchStart   *time.Time `json:"watchStart"`
	WatchEnd     *time.Time `json:"watchEnd"`
	IsRewatching bool       `json:"isRewatching"`
	Days         int        `json:"days"`
	Storage      string     `json:"storage"`
	Priority     string     `json:"priority"`
}

// UserRawManga represents MyAnimeList raw JSON manga list response.
type UserRawManga struct {
	Status         int         `json:"status"`
	Score          int         `json:"score"`
	Tags           interface{} `json:"tags"`
	IsRereading    interface{} `json:"is_rereading"`
	ReadChapter    int         `json:"num_read_chapters"`
	ReadVolume     int         `json:"num_read_volumes"`
	MangaTitle     interface{} `json:"manga_title"`
	MangaChapter   int         `json:"manga_num_chapters"`
	MangaVolume    int         `json:"manga_num_volumes"`
	MangaStatus    int         `json:"manga_publishing_status"`
	MangaID        int         `json:"manga_id"`
	MangaMagazine  []Item      `json:"manga_magazines"`
	MangaURL       string      `json:"manga_url"`
	MangaImage     string      `json:"manga_image_path"`
	IsAddedToList  bool        `json:"is_added_to_list"`
	MangaType      string      `json:"manga_media_type_string"`
	MangaStartDate string      `json:"manga_start_date_string"`
	MangaEndDate   string      `json:"manga_end_date_string"`
	Retail         string      `json:"retail_string"`
	Days           int         `json:"days_string"`
	StartDate      string      `json:"start_date_string"`
	FinishDate     string      `json:"finish_date_string"`
	Priority       string      `json:"priority_string"`
}

// UserManga represents user manga model.
type UserManga struct {
	ID               int        `json:"id"`
	Title            string     `json:"title"`
	Image            string     `json:"image"`
	Score            int        `json:"score"`
	Status           int        `json:"status"`
	Type             string     `json:"type"`
	ChapterProgress  int        `json:"chapterProgress"`
	VolumeProgress   int        `json:"volumeProgress"`
	Chapter          int        `json:"chapter"`
	Volume           int        `json:"volume"`
	Tag              string     `json:"tag"`
	PublishingStatus int        `json:"publishingStatus"`
	PublishingStart  Date       `json:"publishingStart"`
	PublishingEnd    Date       `json:"publishingEnd"`
	ReadStart        *time.Time `json:"readStart"`
	ReadEnd          *time.Time `json:"readEnd"`
	IsRereading      bool       `json:"isRereading"`
	Days             int        `json:"days"`
	Priority         string     `json:"priority"`
}
