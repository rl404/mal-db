package database

import (
	"reflect"
	"strings"

	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-db/internal/constant"
	"github.com/rl404/mal-db/internal/model/raw"
	"github.com/rl404/mal-db/internal/pkg/utils"
	"gorm.io/gorm"
)

// Migrate to prepare database and table schema and data.
func Migrate(db *gorm.DB, mal *malscraper.Malscraper) (err error) {
	// Migrate tables.
	err = db.AutoMigrate(
		// Main tables.
		&raw.Anime{},
		&raw.Manga{},
		&raw.Character{},
		&raw.People{},
		&raw.EmptyID{},

		// Relation tables.
		&raw.AnimeCharacter{},
		&raw.AnimeProducer{},
		&raw.AnimeStaff{},
		&raw.MangaCharacter{},
		&raw.MangaMagazine{},
		&raw.MediaGenre{},
		&raw.MediaRelated{},
		&raw.PeopleManga{},
		&raw.Song{},
		&raw.Stats{},
		&raw.StatsHistory{},

		// Master tables.
		&raw.Genre{},
		&raw.Language{},
		&raw.Position{},
		&raw.ProducerMagazine{},
		&raw.Rating{},
		&raw.Related{},
		&raw.Source{},
		&raw.Status{},
		&raw.Type{},
	)

	if err != nil {
		return err
	}

	// Begin transaction.
	tx := db.Begin()
	if err = tx.Error; err != nil {
		return err
	}
	defer tx.Rollback()

	// Fill master tables.
	if err = fillMaster(tx); err != nil {
		return err
	}

	// Fill producer magazine table.
	if err = fillProducerMagazine(tx, mal); err != nil {
		return err
	}

	// Fill genre table.
	if err = fillGenre(tx, mal); err != nil {
		return err
	}

	// Commit transaction.
	return tx.Commit().Error
}

func fillMaster(db *gorm.DB) (err error) {
	masterList := []struct {
		name  string
		model interface{}
		data  map[int]string
	}{
		{raw.Source{}.TableName(), raw.Source{}, constant.Sources},
		{raw.Rating{}.TableName(), raw.Rating{}, constant.Ratings},
		{raw.Language{}.TableName(), raw.Language{}, constant.Languages},
		{raw.Position{}.TableName(), raw.Position{}, constant.Positions},
		{raw.Related{}.TableName(), raw.Related{}, constant.RelatedTypes},
	}

	for _, t := range masterList {
		if err = fillTable(db, t.name, t.model, t.data); err != nil {
			return err
		}
	}

	masterList2 := []struct {
		name  string
		model interface{}
		data  map[string]map[int]string
	}{
		{raw.Type{}.TableName(), raw.Type{}, constant.Types},
		{raw.Status{}.TableName(), raw.Status{}, constant.Statuses},
	}

	for _, t := range masterList2 {
		if err = fillTableWithMap(db, t.name, t.model, t.data); err != nil {
			return err
		}
	}

	return nil
}

func fillTable(db *gorm.DB, name string, model interface{}, datas map[int]string) (err error) {
	// Clean existing data.
	if err = db.Delete(model).Error; err != nil {
		return err
	}

	// Check if table anime_rating.
	colCount := 2
	isRating := reflect.TypeOf(model) == reflect.TypeOf(raw.Rating{})
	if isRating {
		colCount = 3
	}

	// Create raw batch insert query.
	query := utils.BatchInsertQuery(name, len(datas), colCount)
	var args []interface{}
	for k, d := range datas {
		if isRating {
			split := strings.Split(d, " - ")
			args = append(args, k)
			args = append(args, split[0])
			args = append(args, split[1])
		} else {
			args = append(args, k)
			args = append(args, d)
		}
	}

	// Create new relation.
	return db.Exec(query, args...).Error
}

func fillTableWithMap(db *gorm.DB, name string, model interface{}, datas map[string]map[int]string) (err error) {
	// Clean existing data.
	if err = db.Delete(model).Error; err != nil {
		return err
	}

	// Prepare raw batch insert query.
	var args []interface{}
	count := 0
	for t, data := range datas {
		for k, d := range data {
			args = append(args, k)
			args = append(args, t)
			args = append(args, d)
			count++
		}
	}

	// Create raw batch insert query.
	query := utils.BatchInsertQuery(name, count, 3)

	// Create new relation.
	return db.Exec(query, args...).Error
}

func fillProducerMagazine(db *gorm.DB, mal *malscraper.Malscraper) error {
	// Get producers.
	producers, _, err := mal.GetProducers()
	if err != nil {
		return err
	}

	// Get magazines.
	magazines, _, err := mal.GetMagazines()
	if err != nil {
		return err
	}

	// Clean existing data.
	if err = db.Delete(raw.ProducerMagazine{}).Error; err != nil {
		return err
	}

	// Create raw batch insert query.
	query := utils.BatchInsertQuery(raw.ProducerMagazine{}.TableName(), len(producers)+len(magazines), 3)
	var args []interface{}

	for _, p := range producers {
		args = append(args, p.ID)
		args = append(args, constant.AnimeType)
		args = append(args, p.Name)
	}

	for _, m := range magazines {
		args = append(args, m.ID)
		args = append(args, constant.MangaType)
		args = append(args, m.Name)
	}

	// Insert data to db.
	return db.Exec(query, args...).Error
}

func fillGenre(db *gorm.DB, mal *malscraper.Malscraper) error {
	// Get anime genres.
	animeGenres, _, err := mal.GetAnimeGenres()
	if err != nil {
		return err
	}

	// Get manga genres.
	mangaGenres, _, err := mal.GetMangaGenres()
	if err != nil {
		return err
	}

	// Clean existing data.
	if err = db.Delete(raw.Genre{}).Error; err != nil {
		return err
	}

	// Create raw batch insert query.
	query := utils.BatchInsertQuery(raw.Genre{}.TableName(), len(animeGenres)+len(mangaGenres), 3)
	var args []interface{}

	for _, g := range animeGenres {
		args = append(args, g.ID)
		args = append(args, constant.AnimeType)
		args = append(args, g.Name)
	}

	for _, g := range mangaGenres {
		args = append(args, g.ID)
		args = append(args, constant.MangaType)
		args = append(args, g.Name)
	}

	// Insert data to db.
	return db.Exec(query, args...).Error
}
