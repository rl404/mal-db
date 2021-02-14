package validator

import "github.com/rl404/mal-db/internal/model"

// GetEntryCount to get all entry count.
func (v *Validator) GetEntryCount() (*model.Total, map[string]interface{}, int, error) {
	return v.api.GetEntryCount()
}

// GetYearSummary to get yearly anime & manga summary.
func (v *Validator) GetYearSummary() ([]model.YearSummary, map[string]interface{}, int, error) {
	return v.api.GetYearSummary()
}
