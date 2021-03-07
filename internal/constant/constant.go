package constant

// Main types list.
const (
	AnimeType     = "anime"
	MangaType     = "manga"
	CharacterType = "character"
	PeopleType    = "people"
)

// MainTypes is list of available entry types.
var MainTypes = []string{AnimeType, MangaType, CharacterType, PeopleType}

// Types is list of anime & manga type options.
var Types = map[string]map[int]string{
	AnimeType: {
		1: "TV",
		2: "OVA",
		3: "Movie",
		4: "Special",
		5: "ONA",
		6: "Music",
	},
	MangaType: {
		1: "Manga",
		2: "Light Novel",
		3: "One-shot",
		4: "Doujinshi",
		5: "Manhwa",
		6: "Manhua",
		7: "OEL",
		8: "Novel",
	},
}

// Statuses is list of anime & manga status options.
var Statuses = map[string]map[int]string{
	AnimeType: {
		1: "Currently Airing",
		2: "Finished Airing",
		3: "Not yet aired",
	},
	MangaType: {
		1: "Publishing",
		2: "Finished",
		3: "Not yet published",
		4: "On Hiatus",
		5: "Discontinued",
	},
}

// Sources is list of anime source options.
var Sources = map[int]string{
	1:  "Original",
	2:  "Manga",
	3:  "4-koma Manga",
	4:  "Web Manga",
	5:  "Digital Manga",
	6:  "Novel",
	7:  "Light Novel",
	8:  "Visual Novel",
	9:  "Game",
	10: "Card Game",
	11: "Book",
	12: "Picture Book",
	13: "Radio",
	14: "Music",
}

// Ratings is list of anime rating options.
var Ratings = map[int]string{
	1: "G - All Ages",
	2: "PG - Children",
	3: "PG-13 - Teens 13 or Older",
	4: "R - 17+ (violence & profanity)",
	5: "R+ - Mild Nudity",
	6: "Rx - Hentai",
}

// RelatedTypes is list of anime & mnaga related type options.
var RelatedTypes = map[int]string{
	1:  "sequel",
	2:  "prequel",
	3:  "alternative setting",
	4:  "alternative version",
	6:  "side story",
	7:  "summary",
	8:  "full story",
	9:  "parent story",
	10: "spin-off",
	11: "adaptation",
	12: "character",
	13: "other",
}

// Anime song types.
const (
	OpeningSong = iota + 1
	EndingSong
)

// Languages is list of voice actor language options.
var Languages = map[int]string{
	1:  "Japanese",
	2:  "English",
	3:  "Korean",
	4:  "Spanish",
	5:  "German",
	6:  "French",
	7:  "Brazilian",
	8:  "Italian",
	9:  "Hungarian",
	10: "Hebrew",
	11: "Mandarin",
}

// Positions is list of anime staff position options.
var Positions = map[int]string{
	2:  "Director",
	3:  "Script",
	4:  "Storyboard",
	5:  "Episode Director",
	6:  "Music",
	7:  "Original Creator",
	8:  "Original Character Design",
	9:  "Creator",
	10: "Character Design",
	11: "Art Director",
	12: "Chief Animation Director",
	13: "Animation Director",
	14: "Mechanical Design",
	15: "Director of Photography",
	16: "Executive Producer",
	17: "2nd Key Animation",
	18: "Animation Check",
	19: "Assistant Director",
	20: "Associate Producer",
	21: "Background Art",
	22: "Color Design",
	23: "Color Setting",
	24: "Digital Paint",
	25: "Editing",
	26: "In-Between Animation",
	27: "Key Animation",
	28: "Online Editing Supervision",
	29: "Online Editor",
	30: "Planning",
	31: "Planning Producer",
	32: "Production Manager",
	33: "Publicity",
	34: "Recording",
	35: "Recording Assistant",
	36: "Series Production Director",
	37: "Setting",
	38: "Setting Manager",
	39: "Sound Director",
	40: "Sound Effects",
	41: "Sound Manager",
	42: "Special Effects",
	43: "Theme Song Arrangement",
	44: "Theme Song Composition",
	45: "Theme Song Lyrics",
	46: "Theme Song Performance",
	47: "ADR Director",
	48: "Co-Director",
	50: "Assistant Producer",
	51: "Producer",
	52: "Assistant Engineer",
	53: "Assistant Production Coordinat",
	54: "Associate Casting Director",
	55: "Casting Director",
	56: "Chief Producer",
	57: "Co-Producer",
	58: "Dialogue Editing",
	59: "Inserted Song Performance",
	60: "Post-Production Assistant",
	61: "Production Assistant",
	62: "Production Coordination",
	64: "Re-Recording Mixing",
	65: "Recording Engineer",
	66: "Sound Supervisor",
	67: "Spotting",
	68: "Assistant Animation Director",
	69: "Principle Drawing",
	70: "Layout",
	71: "Screenplay",
	72: "Series Composition",
	// For manga.
	101: "Story & Art",
	102: "Story",
	103: "Art",
}

// Season list.
const (
	Winter = "winter"
	Spring = "spring"
	Summer = "summer"
	Fall   = "fall"
)

// Seasons contains all season name.
var Seasons = []string{Winter, Spring, Summer, Fall}

// Order list.
const (
	OrderMemberA   = "member"
	OrderMemberD   = "-member"
	OrderTitleA    = "title"
	OrderTitleD    = "-title"
	OrderScoreA    = "score"
	OrderScoreD    = "-score"
	OrderNameA     = "name"
	OrderNameD     = "-name"
	OrderFavoriteA = "favorite"
	OrderFavoriteD = "-favorite"
)

// Orders is anime & manga search order.
var Orders = []string{OrderMemberA, OrderMemberD, OrderTitleA, OrderTitleD, OrderScoreA, OrderScoreD}

// Orders2 is character & people search order.
var Orders2 = []string{OrderNameA, OrderNameD, OrderFavoriteA, OrderFavoriteD}

// Orders3 is score comparison order.
var Orders3 = []string{OrderTitleA, OrderTitleD, OrderScoreA, OrderScoreD}
