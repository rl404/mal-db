package malscraper

// Log level list.
// Used for initiating logger.
const (
	LevelZero  = iota // no log
	LevelError        // error, fatal
	LevelInfo         // info, error, fatal
	LevelDebug        // debug, info, warning, error, fatal
	LevelTrace        // trace, debug, info, warning, error, fatal

	// Default level.
	LevelDefault = LevelError
)

// Main types.
const (
	AllType = iota
	AnimeType
	MangaType
)

var mainTypes = []string{"", "anime", "manga"}

// Season list.
const (
	Winter = "winter"
	Spring = "spring"
	Summer = "summer"
	Fall   = "fall"
)

// Top anime types.
const (
	TopDefault = iota
	TopAiring
	TopUpcoming
	TopTV
	TopMovie
	TopOVA
	TopONA
	TopSpecial
	TopPopularAnime
	TopFavoriteAnime
)

// Top manga types.
const (
	TopManga = iota + 1
	TopNovel
	TopOneshot
	TopDoujin
	TopManhwa
	TopManhua
	TopPopularManga
	TopFavoriteManga
)

// Anime types.
const (
	TypeDefault = iota
	TypeTV
	TypeOVA
	TypeMovie
	TypeSpecial
	TypeONA
	TypeMusic
)

// Manga types.
const (
	TypeManga = iota + 1
	TypeLightNovel
	TypeOneShot
	TypeDoujinshi
	TypeManhwa
	TypeManhua
	_
	TypeNovel
)

// Anime & manga airing/publishing status.
const (
	StatusOnGoing = iota + 1
	StatusFinished
	StatusUpcoming
	StatusHiatus       // manga only
	StatusDiscontinued // manga only
)

// Anime ratings.
const (
	RatingDefault = iota
	RatingG       // all ages
	RatingPG      // children
	RatingPG13    // teens 13 or older
	RatingR17     // 17+ (violence & profanity)
	RatingR       // mild nudity
	RatingRx      // hentai
)

// User list status.
const (
	StatusDefault = iota
	StatusCurrent
	StatusCompleted
	StatusOnHold
	StatusDropped
	_
	StatusPlanned
	StatusAll
)

// User anime list order.
const (
	OrderDefault = iota
	OrderAnimeTitle
	OrderAnimeFinishDate
	OrderAnimeStartDate
	OrderAnimeScore
	_
	OrderAnimeType
	_
	OrderAnimeRated
	_
	_
	OrderAnimePriority
	OrderAnimeProgress
	OrderAnimeStorage
	OrderAnimeAirStart
	OrderAnimeAirEnd
)

// User manga list order.
const (
	OrderMangaTitle = iota + 1
	OrderMangaFinishDate
	OrderMangaStartDate
	OrderMangaScore
	_
	_
	OrderMangaPriority
	OrderMangaChapter
	OrderMangaVolume
	OrderMangaType
	OrderMangaPublishStart
	OrderMangaPublishEnd
)

// Review types.
const (
	AnimeReview = iota
	MangaReview
	BestReview
)

var reviewStr = []string{"anime", "manga", "bestvoted"}

// Gender list.
const (
	GenderDefault = iota
	GenderMale
	GenderFemale
	GenderNonBinary
)

// Club categories.
const (
	AllCategory = iota
	AnimeCategory
	ConventionCategory
	ActorCategory
	CharacterCategory
	CompanyCategory
	GameCategory
	JapanCategory
	CityCategory
	MusicCategory
	MangaCategory
	SchoolCategory
	OtherCategory
)

// Club sorts.
const (
	SortDefault = iota
	SortName
	SortComment
	SortPost
	_
	SortMember
)
