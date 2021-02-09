package saver

import (
	"github.com/rl404/mal-db/internal/model/raw"
)

func (a *API) insertEmptyID(t string, id int) error {
	return a.db.Create(&raw.EmptyID{
		Type: t,
		ID:   id,
	}).Error
}

func (a *API) deleteEmptyID(t string, id int) error {
	return a.db.Delete(&raw.EmptyID{
		Type: t,
		ID:   id,
	}).Error
}
