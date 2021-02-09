// Package malscraper provides methods to parse MyAnimeList web page.
//
// Malscraper will get HTML body response, parse necessary information from the elements,
// clean them, and convert to golang model (struct) that can be used easily
// by other project/package. Malscraper is using "github.com/PuerkitoBio/goquery" package
// for parsing the HTML body.
//
//  // Init malscraper.
//  m, err := malscraper.NewDefault()
//  if err != nil {
//   	// handle error
//  }
//
//  // Don't forget to close.
//  defer m.Close()
//
//  // Parse anime data.
//  anime, err := m.GetAnime(1)
//  if err != nil {
//  	// handle error
//  }
//
//  // Use anime data.
//  fmt.Println(anime.Title)
//
// Caching
//
// Malscraper provides caching system. As default, malscraper is using in-memory
// but it is recommended to use persistent cache such as redis. That's why malscraper allows
// you to create your own cacher which implements this interface. Or you can
// choose from `https://github.com/rl404/mal-plugin/tree/master/cache`.
//
//  type Cacher interface {
//  	Get(key string, data interface{}) error
//  	Set(key string, data interface{}) error
//  	Delete(key string) error
//  	Close() error
//  }
//
// And use it when initiating malscraper.
//
//  m, err := malscraper.New(malscraper.Config{
//  	Cacher: yourCacher,
//  })
//
// Logging
//
// Logging is also interface so you can use your own logging.
//
//  type Logger interface {
//  	Trace(format string, args ...interface{})
//  	Debug(format string, args ...interface{})
//  	Info(format string, args ...interface{})
//  	Warn(format string, args ...interface{})
//  	Error(format string, args ...interface{})
//  	Fatal(format string, args ...interface{})
//  }
//
// And use it when initiating malscraper.
//
//  m, err := malscraper.New(malscraper.Config{
//  	Logger: yourLogger,
//  })
//
// Params
//
// Some methods require specific value for the parameter. So, it is recommended to
// fill the parameter with provided constant to prevent unwanted errors.
//
//  m, err := malscraper.New(malscraper.Config{
//  	CacheTime:     24 * time.Hour,
//  	CleanImageURL: true,
//  	CleanVideoURL: true,
//  	LogLevel:      malscraper.LevelTrace,
//  	LogColor:      true,
//  })
//
//  m.GetRecommendation(malscraper.AnimeType, 1, 6)
//  m.GetSeason(malscraper.Winter, 2019)
//  m.GetTopAnime(malscraper.TopDefault, 2)
//
// Error
//
// Errors returned by methods are guaranteed from malscraper errors package. You should
// be able to handle the errors easier. The original errors (from the dependency package)
// can still be viewed through printed log (if you turn on error log level). All methods
// also return HTTP response code to help you distinguish the error.
//
// Request Limit
//
// All methods are requesting and accessing MyAnimeList web page so use them
// responsibly and don't spam too much or your ip will get banned or blacklisted.
// Recommended 1 request per 3 seconds and use caching. Or for more safety, 1 request per 5 seconds.
package malscraper
