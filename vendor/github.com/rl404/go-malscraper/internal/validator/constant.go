package validator

// Main types.
const (
	AnimeType = "anime"
	MangaType = "manga"
)

// Season list.
const (
	Winter = "winter"
	Spring = "spring"
	Summer = "summer"
	Fall   = "fall"
)

var seasons = []string{Winter, Spring, Summer, Fall}

// TopDefault is default top list.
const TopDefault = iota

// Top anime types.
const (
	TopAiring = iota + 1
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

var topAnimeTypes = []int{TopDefault, TopAiring, TopUpcoming, TopTV, TopMovie, TopOVA, TopONA, TopSpecial, TopPopularAnime, TopFavoriteAnime}
var topMangaTypes = []int{TopDefault, TopManga, TopNovel, TopOneshot, TopDoujin, TopManhwa, TopManhua, TopPopularManga, TopFavoriteManga}

// TypeDefault is default anime/manga type.
const TypeDefault = iota

// Anime types.
const (
	TypeTV = iota + 1
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

var animeTypes = []int{TypeDefault, TypeTV, TypeOVA, TypeMovie, TypeSpecial, TypeONA, TypeMusic}
var mangaTypes = []int{TypeDefault, TypeManga, TypeLightNovel, TypeOneShot, TypeDoujinshi, TypeManhwa, TypeManhua, TypeNovel}

// Anime & manga airing/publishing status.
const (
	StatusOnGoing = iota + 1
	StatusFinished
	StatusUpcoming
	StatusHiatus       // manga only
	StatusDiscontinued // manga only
)

var animeStatuses = []int{StatusDefault, StatusOnGoing, StatusFinished, StatusUpcoming}
var mangaStatuses = []int{StatusDefault, StatusOnGoing, StatusFinished, StatusUpcoming, StatusHiatus, StatusDiscontinued}

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

var statuses = []int{StatusDefault, StatusCurrent, StatusCompleted, StatusOnHold, StatusDropped, StatusPlanned, StatusAll}

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

var ratings = []int{RatingDefault, RatingG, RatingPG, RatingPG13, RatingR17, RatingR, RatingRx}

// OrderDefault is default user list order.
const OrderDefault = iota

// User anime list order.
const (
	OrderAnimeTitle = iota + 1
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

var animeOrders = []int{OrderDefault, OrderAnimeTitle, OrderAnimeFinishDate, OrderAnimeStartDate, OrderAnimeScore, OrderAnimeType, OrderAnimeRated, OrderAnimePriority, OrderAnimeProgress, OrderAnimeStorage, OrderAnimeAirStart, OrderAnimeAirEnd}
var mangaOrders = []int{OrderDefault, OrderMangaTitle, OrderMangaFinishDate, OrderMangaStartDate, OrderMangaScore, OrderMangaPriority, OrderMangaChapter, OrderMangaVolume, OrderMangaType, OrderMangaPublishStart, OrderMangaPublishEnd}

// Review types.
const (
	AnimeReview = "anime"
	MangaReview = "manga"
	BestReview  = "bestvoted"
)

// Gender list.
const (
	GenderDefault = iota
	GenderMale
	GenderFemale
	GenderNonBinary
)

var genders = []int{GenderDefault, GenderMale, GenderFemale, GenderNonBinary}

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

var categories = []int{AllCategory, AnimeCategory, ConventionCategory, ActorCategory, CharacterCategory, CompanyCategory, GameCategory, JapanCategory, CityCategory, MusicCategory, MangaCategory, SchoolCategory, OtherCategory}

// Club sorts.
const (
	SortDefault = iota
	SortName
	SortComment
	SortPost
	_
	SortMember
)

var sorts = []int{SortDefault, SortName, SortComment, SortPost, SortMember}
