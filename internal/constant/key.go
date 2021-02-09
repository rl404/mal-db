package constant

import (
	"fmt"
	"strings"
)

// PubSubTopic is pubsub topic/channel name.
const PubSubTopic = "mal-db"

// List of redis key used in malkeeper.
const (
	KeyAnime            = "mal:2:anime"
	KeyAnimeCharacter   = "mal:2:anime-character"
	KeyAnimeStaff       = "mal:2:anime-staff"
	KeyManga            = "mal:2:manga"
	KeyMangaCharacter   = "mal:2:manga-character"
	KeyStats            = "mal:2:stats"
	KeyCharacter        = "mal:2:character"
	KeyCharacterOgraphy = "mal:2:character-ography"
	KeyCharacterVA      = "mal:2:character-va"
	KeyPeople           = "mal:2:people"
	KeyPeopleVA         = "mal:2:people-va"
	KeyPeopleStaff      = "mal:2:people-staff"
	KeyPeopleManga      = "mal:2:people-manga"
	KeyProducerMagazine = "mal:2:producer-magazine"
	KeyGenres           = "mal:2:genres"
	KeyTotal            = "mal:2:total"

	KeyTopAnime     = "mal:2:top-anime"
	KeyTopManga     = "mal:2:top-manga"
	KeyTopCharacter = "mal:2:top-character"
	KeyTopPeople    = "mal:2:top-people"
	KeySeasonal     = "mal:2:seasonal"
	KeyYearScore    = "mal:2:year-score"
	KeyProducer     = "mal:2:producer"
	KeyMagazine     = "mal:2:magazine"
	KeyUser         = "mal:2:user"
	KeyUserAnime    = "mal:2:user-anime"
	KeyUserManga    = "mal:2:user-manga"
	KeyUserStats    = "mal:2:user-stats"
	KeyUserScore    = "mal:2:user-score"
	KeyUserType     = "mal:2:user-type"
	KeyUserGenre    = "mal:2:user-genre"
	KeyUserStudio   = "mal:2:user-studio"
	KeyUserAuthor   = "mal:2:user-author"
	KeyUserProgress = "mal:2:user-progress"
	KeyUserYear     = "mal:2:user-year"
)

// GetKey to generate cache key.
func GetKey(key string, params ...interface{}) string {
	strParams := []string{key}
	for _, p := range params {
		strParams = append(strParams, fmt.Sprintf("%v", p))
	}
	return strings.Join(strParams, ":")
}
