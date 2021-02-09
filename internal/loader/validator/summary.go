package validator

import "github.com/rl404/mal-db/internal/model"

// GetEntryCount to get all entry count.
func (v *Validator) GetEntryCount() (*model.Total, map[string]interface{}, int, error) {
	return v.api.GetEntryCount()
}
