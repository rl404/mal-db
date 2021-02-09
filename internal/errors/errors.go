package errors

import "errors"

// Error list.
var (
	ErrInvalidCacheType  = errors.New("invalid cache type (nocache|inmemory|redis|memcache)")
	ErrRequiredDB        = errors.New("require database")
	ErrInvalidDBFormat   = errors.New("invalid database address format (host:port)")
	ErrRequiredPubsub    = errors.New("required pubsub")
	ErrInvalidPubsubType = errors.New("invalid pubsub type (redis|nsq|rabbitmq)")
	ErrInvalidType       = errors.New("invalid type")
	ErrInvalidID         = errors.New("invalid id")
	ErrInvalidPage       = errors.New("invalid page")
	ErrInvalidLimit      = errors.New("invalid limit")
	ErrQueueEntry        = errors.New("entry will be queued")
	ErrInvalidSeason     = errors.New("invalid season")
	Err3LettersSearch    = errors.New("search query needs at least 3 letters")
	ErrInvalidScore      = errors.New("invalid score")
	ErrInvalidStatus     = errors.New("invalid status")
	ErrInvalidRating     = errors.New("invalid rating")
	ErrInvalidSource     = errors.New("invalid source")
	ErrInvalidYear       = errors.New("invalid year")
	ErrInvalidOrder      = errors.New("invalid order")
	ErrInvalidEpisode    = errors.New("invalid episode")
	ErrInvalidChapter    = errors.New("invalid chapter")
	ErrInvalidVolume     = errors.New("invalid volume")
	ErrInvalidDuration   = errors.New("invalid duration")
	ErrInvalidProducer   = errors.New("invalid producer")
	ErrInvalidMagazine   = errors.New("invalid magazine")
	ErrNewData           = errors.New("data is still new")
)
