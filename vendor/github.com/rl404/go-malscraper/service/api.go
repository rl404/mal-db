package service

import "github.com/rl404/go-malscraper/model"

// API contains all malscraper api function.
// All function return HTTP response code for helping to distinguish
// the error when building REST API.
type API interface {
	// Anime.
	GetAnime(id int) (*model.Anime, int, error)
	GetAnimeCharacter(id int) ([]model.CharacterItem, int, error)
	GetAnimeStaff(id int) ([]model.Role, int, error)
	GetAnimeVideo(id int, page int) (*model.Video, int, error)
	GetAnimeEpisode(id int, page int) ([]model.Episode, int, error)
	GetAnimeStats(id int) (*model.Stats, int, error)
	GetAnimeReview(id int, page int) ([]model.Review, int, error)
	GetAnimeRecommendation(id int) ([]model.Recommendation, int, error)
	GetAnimeNews(id int) ([]model.NewsItem, int, error)
	GetAnimeArticle(id int) ([]model.ArticleItem, int, error)
	GetAnimeClub(id int) ([]model.ClubItem, int, error)
	GetAnimePicture(id int) ([]string, int, error)
	GetAnimeMoreInfo(id int) (string, int, error)

	// Manga.
	GetManga(id int) (*model.Manga, int, error)
	GetMangaReview(id int, page int) ([]model.Review, int, error)
	GetMangaRecommendation(id int) ([]model.Recommendation, int, error)
	GetMangaStats(id int) (*model.Stats, int, error)
	GetMangaCharacter(id int) ([]model.Role, int, error)
	GetMangaNews(id int) ([]model.NewsItem, int, error)
	GetMangaArticle(id int) ([]model.ArticleItem, int, error)
	GetMangaClub(id int) ([]model.ClubItem, int, error)
	GetMangaPicture(id int) ([]string, int, error)
	GetMangaMoreInfo(id int) (string, int, error)

	// Character.
	GetCharacter(id int) (*model.Character, int, error)
	GetCharacterOgraphy(_type string, id int) ([]model.Role, int, error)
	GetCharacterArticle(id int) ([]model.ArticleItem, int, error)
	GetCharacterPicture(id int) ([]string, int, error)
	GetCharacterClub(id int) ([]model.ClubItem, int, error)
	GetCharacterVA(id int) ([]model.Role, int, error)

	// People.
	GetPeople(id int) (*model.People, int, error)
	GetPeopleCharacter(id int) ([]model.PeopleCharacter, int, error)
	GetPeopleStaff(id int) ([]model.Role, int, error)
	GetPeopleManga(id int) ([]model.Role, int, error)
	GetPeopleNews(id int) ([]model.NewsItem, int, error)
	GetPeopleArticle(id int) ([]model.ArticleItem, int, error)
	GetPeoplePicture(id int) ([]string, int, error)

	// Producer & magazine.
	GetProducers() ([]model.ItemCount, int, error)
	GetProducer(id int, page int) ([]model.AnimeItem, int, error)
	GetMagazines() ([]model.ItemCount, int, error)
	GetMagazine(id int, page int) ([]model.MangaItem, int, error)

	// Genre.
	GetGenres(_type string) ([]model.ItemCount, int, error)
	GetAnimeWithGenre(id int, page int) ([]model.AnimeItem, int, error)
	GetMangaWithGenre(id int, page int) ([]model.MangaItem, int, error)

	// Review.
	GetReview(id int) (*model.Review, int, error)
	GetReviews(_type string, page int) ([]model.Review, int, error)

	// Recommendation.
	GetRecommendation(_type string, id1, id2 int) (*model.Recommendation, int, error)
	GetRecommendations(_type string, page int) ([]model.Recommendation, int, error)

	// Seasonal.
	GetSeason(season string, year int) ([]model.AnimeItem, int, error)

	// News.
	GetNews(id int) (*model.News, int, error)
	GetNewsList(page int, tag string) ([]model.NewsItem, int, error)
	GetNewsTag() (*model.NewsTag, int, error)

	// Article article.
	GetArticle(id int) (*model.Article, int, error)
	GetArticles(page int, tag string) ([]model.ArticleItem, int, error)
	GetArticleTag() ([]model.ArticleTagItem, int, error)

	// Club.
	GetClubs(page int) ([]model.ClubSearch, int, error)
	GetClub(id int) (*model.Club, int, error)
	GetClubMember(id int, page int) ([]model.ClubMember, int, error)
	GetClubPicture(id int) ([]string, int, error)
	GetClubRelated(id int) (*model.ClubRelated, int, error)

	// Top list.
	GetTopAnime(_type int, page int) ([]model.TopAnime, int, error)
	GetTopManga(_type int, page int) ([]model.TopManga, int, error)
	GetTopCharacter(page int) ([]model.TopCharacter, int, error)
	GetTopPeople(page int) ([]model.TopPeople, int, error)

	// User.
	GetUser(username string) (*model.User, int, error)
	GetUserStats(username string) (*model.UserStats, int, error)
	GetUserFavorite(username string) (*model.UserFavorite, int, error)
	GetUserFriend(username string, page int) ([]model.UserFriend, int, error)
	GetUserHistory(username string, _type string) ([]model.UserHistory, int, error)
	GetUserReview(username string, page int) ([]model.Review, int, error)
	GetUserRecommendation(username string, page int) ([]model.Recommendation, int, error)
	GetUserClub(username string) ([]model.Item, int, error)
	GetUserAnime(query model.UserListQuery) ([]model.UserAnime, int, error)
	GetUserManga(query model.UserListQuery) ([]model.UserManga, int, error)

	// Search.
	SearchAnime(query model.Query) ([]model.AnimeSearch, int, error)
	SearchManga(query model.Query) ([]model.MangaSearch, int, error)
	SearchCharacter(query string, page int) ([]model.CharacterSearch, int, error)
	SearchPeople(query string, page int) ([]model.PeopleSearch, int, error)
	SearchClub(query model.ClubQuery) ([]model.ClubSearch, int, error)
	SearchUser(query model.UserQuery) ([]model.UserSearch, int, error)
}
