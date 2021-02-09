# mal-db

[MyAnimeList](https://myanimelist.net/) database dump and REST API.

Powered by my [go-malscraper](https://github.com/rl404/go-malscraper).

## Features

- Save anime detail, character list, staff list, and stats.
- Save manga detail, character list, and stats.
- Save character detail, animeography list, mangaography list, and voice actor list.
- Save people detail, anime character role list, anime staff role list, and published manga list.
- Database ([postgresql](https://www.postgresql.org/))
- Caching (in-memory, [redis](https://redis.io/), [memcache](https://memcached.org/))
- Logging ([elasticsearch](https://www.elastic.co/))
- PubSub ([redis](https://redis.io/), [nsq](https://nsq.io/), [rabbitmq](https://www.rabbitmq.com/))
- [Swagger](https://mal-rest-api.herokuapp.com/swagger/index.html)
- [Docker](https://www.docker.com/)

*More will be coming soon...*

## Requirement

- Database ([postgresql](https://www.postgresql.org/))
- PubSub ([redis](https://redis.io/) or [nsq](https://nsq.io/) or [rabbitmq](https://www.rabbitmq.com/))
- (optional) Caching ([redis](https://redis.io/) or [memcache](https://memcached.org/))

## Installation

### With [Go](https://golang.org/)

1. Clone the repository.
```
git clone github.com/rl404/mal-db
```
2. Update `.env` file.
3. Prepare the database.
```
make install
```
4. Build and run worker.
```
make worker
```
5. In another console/terminal, build and run server.
```
make server
```

### With [Docker](https://www.docker.com/) & [Docker compose](https://docs.docker.com/compose/)

1. Clone the repository.
```
git clone github.com/rl404/mal-db
```
2. Update `.env` file.
3. Prepare the database.
```
make docker-install
```
4. Build and run server & worker.
```
make docker-up
```
To stop containers.
```
make docker-stop
```

## Config

> Env are optional except the ones with `*`.

Env | Default | Description
--- | :---: | ---
`MAL_WEB_PORT` | `8006` | HTTP port
`MAL_WEB_READ_TIMEOUT` | `5` | HTTP read timeout (in seconds)
`MAL_WEB_WRITE_TIMEOUT` | `5` | HTTP write timeout (in seconds)
`MAL_WEB_GRACEFUL_TIMEOUT` | `10` | HTTP server shutdown timeout (in seconds)
`MAL_WORKER_AGE_LIMIT` | `604800` | Entry age that needs to be re-parsed (in seconds)
`MAL_WORKER_BREAK_TIME` | `5` | Break time between parsing (in seconds)
`MAL_CACHE_DIALECT` | `inmemory` | Cache type (`nocache`, `inmemory`, `redis`, `memcache`)
`MAL_CACHE_ADDRESS` |  | Cache address
`MAL_CACHE_PASSWORD` |  | Cache password
`MAL_CACHE_TIME` | `86400` | Cache time (in seconds)
`MAL_LOG_LEVEL` | `4` | Log all
`MAL_LOG_COLOR` | `true` | Log color
`MAL_DB_ADDRESS*` | | Postgresql host and port
`MAL_DB_NAME` | | Database name
`MAL_DB_USER` | | Database username
`MAL_DB_PASSWORD` | | Database password
`MAL_DB_MAX_CONN_OPEN` | `10` | Max database open connection
`MAL_DB_MAX_CONN_IDLE` | `10` | Max database idle connection
`MAL_DB_MAX_CONN_LIFETIME` | `60` | Max database connection lifetime
`MAL_ES_ADDRESS` | | Elasticsearch host and port
`MAL_ES_USER` | | Elasticsearch user
`MAL_ES_PASSWORD` | | Elasticsearch password
`MAL_PUBSUB_DIALECT*` | | PubSub type (`redis`, `nsq`, `rabbitmq`)
`MAL_PUBSUB_ADDRESS*` | | PubSub address
`MAL_PUBSUB_USER` | | PubSub user
`MAL_PUBSUB_PASSWORD` | | PubSub password

**Required*

### Log Level

Level | Trace | Debug | Info | Warn | Error | Fatal
:---: | :---: | :---: | :---: | :---: | :---: | :---: |
`0` | :x: | :x: | :x: | :x: | :x: | :x:
`1` | :x: | :x: | :x: | :x: | :heavy_check_mark: | :heavy_check_mark:
`2` | :x: | :x: | :heavy_check_mark: | :x: | :heavy_check_mark: | :heavy_check_mark:
`3` | :x: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:
`4` | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:

## Disclamer

_mal-db_ is meant for educational purpose and personal usage only. Although there is no limit in using the API, do remember that every scraper method is accessing MyAnimeList page so use it responsibly according to MyAnimeList's [Terms Of Service](https://myanimelist.net/about/terms_of_use).

All data (including anime, manga, people, etc) belong to their respective copyrights owners. mal-db does not have any affiliation with content providers.

## License

MIT License

Copyright (c) 2021 Axel