# Go-Malscraper

![Github Test](https://github.com/rl404/go-malscraper/workflows/cron/badge.svg)
[![Coverage](https://coveralls.io/repos/github/rl404/go-malscraper/badge.svg)](https://coveralls.io/github/rl404/go-malscraper)
[![Go Report Card](https://goreportcard.com/badge/github.com/rl404/go-malscraper)](https://goreportcard.com/report/github.com/rl404/go-malscraper)
![License: MIT](https://img.shields.io/github/license/rl404/go-malscraper.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/rl404/go-malscraper.svg)](https://pkg.go.dev/github.com/rl404/go-malscraper)

_go-malscraper_ is just another unofficial API which scraping/parsing [MyAnimeList](https://myanimelist.net/) website to a useful and easy-to-use data by using [Go](https://golang.org/).

Well, it is created to help people get MyAnimeList data without relying on MyAnimeList since they limited/disabled/closed their API. It's working as long as the web is up and the UI design stays the same so it can get the page sources and parse them.

_go-malscraper_ is using [PuerkitoBio's](https://github.com/PuerkitoBio/goquery) HTML DOM parser and inspired by [Jikan's](https://github.com/jikan-me/jikan) API library and my PHP [Mal-Scraper](https://github.com/rl404/MAL-Scraper) library.

Looking for REST API one? come [here](https://github.com/rl404/mal-api).

## Features

* Get anime information (details, characters, episodes, pictures, etc)
* Get manga information (details, characters, pictures, recommendations, etc)
* Get character information (details, pictures, etc)
* Get people information (details, pictures, etc)
* Get list of all anime/manga's genres
* Get list of all anime/manga's producers/studios/licensors/magazines/serializations
* Get anime/manga's recommendations
* Get anime/manga's reviews
* Search anime, manga, character and people
* Get seasonal anime list
* Get anime, manga, character and people top list
* Get user information (profile, friends, histories, recommendations, reviews, etc)
* Get news list and details
* Get featured article list and details
* Get club list and details
* Caching

_More will be coming soon..._

## Installation

```
go get github.com/rl404/go-malscraper
```

## Quick Start

```go
package main

import (
    "fmt"
    mal "github.com/rl404/go-malscraper"
)

func main() {
    // Init with default config.
    m, err := mal.NewDefault()
    if err != nil {
        // handle error
    }

    // Don't forget to close.
    defer m.Close()

    // Parse anime ID 1.
    d, _, err := m.GetAnime(1)
    if err != nil {
        // handle error
    }

    // Use.
    fmt.Println(d.Title)
}
```

### With Configuration

```go
m, err := mal.New(mal.Config{
    CacheTime:     24 * time.Hour,
    CleanImageURL: true,
    CleanVideoURL: true,
    LogLevel:      mal.LevelTrace,
    LogColor:      true,
})
```

*For more detail config and usage, please go to the [documentation](https://pkg.go.dev/github.com/rl404/go-malscraper).*

## Disclamer

_go-malscraper_ is meant for educational purpose and personal usage only. Although there is no limit in using the API, do remember that every scraper method is accessing MyAnimeList page so use it responsibly according to MyAnimeList's [Terms Of Service](https://myanimelist.net/about/terms_of_use).

All data (including anime, manga, people, etc) and MyAnimeList logos belong to their respective copyrights owners. go-malscraper does not have any affiliation with content providers.

## License

MIT License

Copyright (c) 2021 Axel
