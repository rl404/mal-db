package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rl404/mal-db/internal/errors"
)

// List of available cache type.
const (
	NoCache  = "nocache"
	InMemory = "inmemory"
	Redis    = "redis"
	Memcache = "memcache"
)

// List of available pubsub type.
const (
	NSQ      = "nsq"
	RabbitMQ = "rabbitmq"
)

var caches = map[string]int{NoCache: 1, InMemory: 1, Redis: 1, Memcache: 1}
var pubsubs = map[string]int{Redis: 1, NSQ: 1, RabbitMQ: 1}

// Config is configuration model for whole malscraper project.
type Config struct {
	// Web server config.
	Web webConfig `envconfig:"WEB"`
	// Worker config.
	Worker workerConfig `envconfig:"WORKER"`
	// Cache config.
	Cache cacheConfig `envconfig:"CACHE"`
	// Database config.
	DB dbConfig `envconfig:"DB"`
	// Logging config.
	Log logConfig `envconfig:"LOG"`
	// Elasticsearch config.
	ES esConfig `envconfig:"ES"`
	// PubSub config.
	PubSub pubSubConfig `envconfig:"PUBSUB"`
}

type webConfig struct {
	// HTTP port.
	Port string `envconfig:"PORT" default:"8006"`
	// Read timeout (in seconds).
	ReadTimeout int `envconfig:"READ_TIMEOUT" default:"5"`
	// Write timeout (in seconds).
	WriteTimeout int `envconfig:"WRITE_TIMEOUT" default:"5"`
	// Graceful shutdown timeout (in seconds).
	GracefulTimeout int `envconfig:"GRACEFUL_TIMEOUT" default:"10"`
}

type workerConfig struct {
	// Old data age limit. Recent parsed data will not be re-parsed.
	AgeLimit int `envconfig:"AGE_LIMIT" default:"604800"`
	// Break time between parsing.
	BreakTime int `envconfig:"BREAK_TIME" default:"5"`
}

type cacheConfig struct {
	// Type of caching (string).
	Dialect string `envconfig:"DIALECT" default:"inmemory"`
	// Cache address with format `host:port`.
	Address string `envconfig:"ADDRESS"`
	// Cache password if exists.
	Password string `envconfig:"PASSWORD"`
	// Caching time duration (in seconds).
	Time int `envconfig:"TIME" default:"86400"`
}

type logConfig struct {
	// Log level.
	Level int `envconfig:"LEVEL" default:"4"`
	// Log color.
	Color bool `envconfig:"COLOR" default:"true"`
}

type dbConfig struct {
	// Database host and port.
	Address string `envconfig:"ADDRESS"`
	// Database name.
	Name string `envconfig:"NAME"`
	// Database user.
	User string `envconfig:"USER"`
	// Database password.
	Password string `envconfig:"PASSWORD"`
	// Max open connection.
	MaxConnOpen int `envconfig:"MAX_CONN_OPEN" default:"10"`
	// Max idle connection.
	MaxConnIdle int `envconfig:"MAX_CONN_IDLE" default:"10"`
	// Max connection lifetime.
	MaxConnLifetime int `envconfig:"MAX_CONN_LIFETIME" default:"60"`
}

type esConfig struct {
	// Elasticsearch addresses. Split by comma.
	Address string `envconfig:"ADDRESS"`
	// Elasticsearch username.
	User string `envconfig:"USER"`
	// Elasticsearch password.
	Password string `envconfig:"PASSWORD"`
}

type pubSubConfig struct {
	// Type of pubsub.
	Dialect string `envconfig:"DIALECT"`
	// Pubsub address.
	Address string `envconfig:"ADDRESS"`
	// Pubsub user if exists.
	User string `envconfig:"USER"`
	// Pubsub password if exists.
	Password string `envconfig:"PASSWORD"`
}

const envPath = "../../.env"
const envPrefix = "MAL"

// GetConfig to read and parse config from `.env`.
func GetConfig() (cfg Config, err error) {
	// Load .env file.
	_ = godotenv.Load(envPath)

	// Convert env to struct.
	if err = envconfig.Process(envPrefix, &cfg); err != nil {
		return cfg, err
	}

	// Override port.
	if port := os.Getenv("PORT"); port != "" {
		cfg.Web.Port = port
	}

	// Validate cache type.
	if caches[cfg.Cache.Dialect] == 0 {
		return cfg, errors.ErrInvalidCacheType
	}

	// Require database config.
	if cfg.DB.Address == "" {
		return cfg, errors.ErrRequiredDB
	}

	// Validate pubsub type.
	if pubsubs[cfg.PubSub.Dialect] == 0 {
		return cfg, errors.ErrInvalidPubsubType
	}

	// Require pubsub config.
	if cfg.PubSub.Address == "" {
		return cfg, errors.ErrRequiredPubsub
	}

	return cfg, nil
}
