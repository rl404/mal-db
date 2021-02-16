package loader

import (
	"time"

	"github.com/rl404/mal-db/internal/cacher"
	"github.com/rl404/mal-db/internal/loader/api"
	_cacher "github.com/rl404/mal-db/internal/loader/cacher"
	"github.com/rl404/mal-db/internal/loader/database"
	"github.com/rl404/mal-db/internal/loader/validator"
	"github.com/rl404/mal-db/internal/logger"
	"github.com/rl404/mal-db/internal/pubsub"
	"gorm.io/gorm"
)

// New to create new loader.
func New(l logger.Logger, db *gorm.DB, c cacher.Cacher, ps pubsub.PubSub, ageLimit time.Duration) (api api.API) {
	api = database.New(l, db, ps, ageLimit)
	api = _cacher.New(l, api, c)
	api = validator.New(api)
	return api
}
