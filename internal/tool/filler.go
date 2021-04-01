package tool

import (
	"fmt"

	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/logger"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pubsub"
	"gorm.io/gorm"
)

// Filler is tools to fill missing entries.
type Filler struct {
	log logger.Logger
	db  *gorm.DB
	ps  pubsub.PubSub
	ids map[string][]int
}

// NewFiller to create new filler tools.
func NewFiller(l logger.Logger, db *gorm.DB, ps pubsub.PubSub) *Filler {
	return &Filler{
		log: l,
		db:  db,
		ps:  ps,
		ids: make(map[string][]int),
	}
}

// Run to run filler tools.
func (f *Filler) Run() error {
	if err := f.getMissingEntries(); err != nil {
		return err
	}
	for _, t := range constant.MainTypes {
		t = "character"
		cnt := 0
		f.log.Info("%s %v", t, len(f.ids[t]))
		for _, id := range f.ids[t] {
			if err := f.ps.Publish(constant.PubSubTopic, pubsub.Message{
				Type: t,
				ID:   id,
			}); err != nil {
				return err
			}
			cnt++
			fmt.Println(id)
			if cnt > 1000 {
				return nil
			}
		}
	}
	return nil
}

func (f *Filler) getMaxID(t string) (id int, err error) {
	err = f.db.Table(t).Select("max(id)").Row().Scan(&id)
	return id, err
}

func (f *Filler) getMissingEntries() error {
	for _, t := range constant.MainTypes {
		maxID, err := f.getMaxID(t)
		if err != nil {
			return err
		}

		entryMap, err := f.getEntryMap(t)
		if err != nil {
			return err
		}

		emptyMap, err := f.getEmptyMap(t)
		if err != nil {
			return err
		}

		for id := 1; id < maxID; id++ {
			if entryMap[id] == 0 && emptyMap[id] == 0 {
				f.ids[t] = append(f.ids[t], id)
			}
		}
	}
	return nil
}

func (f *Filler) getEntryMap(t string) (map[int]int, error) {
	rows, err := f.db.Table(t).Select("id").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	m := make(map[int]int)
	for rows.Next() {
		var tmp struct {
			ID int
		}
		if err = f.db.ScanRows(rows, &tmp); err != nil {
			return nil, err
		}
		m[tmp.ID] = tmp.ID
	}

	return m, nil
}

func (f *Filler) getEmptyMap(t string) (map[int]int, error) {
	rows, err := f.db.Model(&raw.EmptyID{}).Where("type = ?", t).Select("id").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	m := make(map[int]int)
	for rows.Next() {
		var tmp struct {
			ID int
		}
		if err = f.db.ScanRows(rows, &tmp); err != nil {
			return nil, err
		}
		m[tmp.ID] = tmp.ID
	}

	return m, nil
}
