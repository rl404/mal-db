// Package errors contains all errors returned by malscraper.
//
// All errors from dependencies have converted and are guaranteed
// to be one of these. You should be able to handle the error easier.
//
// The original errors are printed by logger (if you turn on the error level).
package errors

import "errors"

var (
	// ErrInitCache if failed initiating cache.
	ErrInitCache = errors.New("failed initiating cache")
	// ErrPrepareRequest if failed creating new HTTP request.
	ErrPrepareRequest = errors.New("failed preparing request to MyAnimeList")
	// ErrHTTPRequest if failed HTTP request to MAL web.
	ErrHTTPRequest = errors.New("failed HTTP request to MyAnimeList")
	// ErrNot200 if MAL doesn't return 200.
	ErrNot200 = errors.New("MyAnimeList not return 200")
	// ErrParseBody if goquery failed parsing MAL body.
	ErrParseBody = errors.New("failed parsing request body")
	// ErrDecodeJSON if failed unmarshaling JSON.
	ErrDecodeJSON = errors.New("failed decoding JSON")
	// ErrInvalidID if id is invalid (must positive and not zero).
	ErrInvalidID = errors.New("invalid ID")
	// Err3LettersSearch if search query string is less than 3 letters.
	Err3LettersSearch = errors.New("search query needs at least 3 letters")
	// ErrInvalidType if not a valid type.
	ErrInvalidType = errors.New("invalid type")
	// ErrInvalidSeason if value is not a valid season name.
	ErrInvalidSeason = errors.New("invalid season name")
	// ErrInvalidYear if year is negative.
	ErrInvalidYear = errors.New("invalid year")
	// ErrInvalidPage if page number is invalid (must positive and not zero).
	ErrInvalidPage = errors.New("invalid page")
	// ErrInvalidScore if score is less than 0 or higher than 10.
	ErrInvalidScore = errors.New("invalid score")
	// ErrInvalidTag if tag is invalid.
	ErrInvalidTag = errors.New("invalid tag")
	// ErrInvalidSortType if sort type is invalid.
	ErrInvalidSortType = errors.New("invalid sort type")
	// ErrInvalidClubCategory if club catergory is invalid.
	ErrInvalidClubCategory = errors.New("invalid club category")
	// ErrInvalidStatus if anime/manga list status is invalid.
	ErrInvalidStatus = errors.New("invalid status")
	// ErrInvalidRating if anime rating is invalid.
	ErrInvalidRating = errors.New("invalid rating")
	// ErrInvalidProducer if producer is invalid.
	ErrInvalidProducer = errors.New("invalid producer")
	// ErrInvalidGenre if genre is invalid.
	ErrInvalidGenre = errors.New("invalid genre")
	// ErrInvalidFirstLetter if letter is more than 1 letter.
	ErrInvalidFirstLetter = errors.New("must be only 1 letter")
	// ErrInvalidMagazine if magazine is invalid.
	ErrInvalidMagazine = errors.New("invalid magazine")
	// ErrInvalidOrder if anime/manga list status is invalid.
	ErrInvalidOrder = errors.New("invalid order")
	// ErrInvalidUsername if username is invalid.
	ErrInvalidUsername = errors.New("invalid username")
	// ErrInvalidAge if age is negative.
	ErrInvalidAge = errors.New("invalid age")
	// ErrInvalidGender if gender is invalid.
	ErrInvalidGender = errors.New("invalid gender")
)
