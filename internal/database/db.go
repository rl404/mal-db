package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/rl404/mal-db/internal/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// New to create new db connection.
func New(address, name, user, password string, maxOpen, maxIdle, maxLife int) (*gorm.DB, error) {
	// Split host and port.
	split := strings.Split(address, ":")
	if len(split) != 2 {
		return nil, errors.ErrInvalidDBFormat
	}

	// Prepare dsn and open connection.
	dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable", split[0], split[1], name, user, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: true,
		Logger:            logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	tmp, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set basic config.
	tmp.SetMaxIdleConns(maxIdle)
	tmp.SetMaxOpenConns(maxOpen)
	tmp.SetConnMaxLifetime(time.Duration(maxLife) * time.Second)

	return db, nil
}
