package validator

import "github.com/rl404/mal-db/internal/loader/api"

// Validator implements API interface.
type Validator struct {
	api api.API
}

// New to create new validator.
func New(api api.API) api.API {
	return &Validator{
		api: api,
	}
}
