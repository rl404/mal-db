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
	KeyYearSummary      = "mal:2:year-summary"
)

// GetKey to generate cache key.
func GetKey(key string, params ...interface{}) string {
	strParams := []string{key}
	for _, p := range params {
		strParams = append(strParams, fmt.Sprintf("%v", p))
	}
	return strings.Join(strParams, ":")
}
