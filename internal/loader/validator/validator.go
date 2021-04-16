package validator

import (
	"github.com/rl404/mal-db/internal/loader"
)

// Validator implements API interface.
type Validator struct {
	api loader.API
}

// New to create new validator.
func New(api loader.API) *Validator {
	return &Validator{
		api: api,
	}
}
