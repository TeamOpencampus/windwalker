package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//SetupDatabase connects to database and migrates schemas.
func SetupDatabase() (*gorm.DB, error) {
	uri, ok := os.LookupEnv("DATABASE_URL")
	var dialect gorm.Dialector
	if ok {
		dialect = postgres.Open(uri)
	} else {
		dialect = sqlite.Open("windwalker.sqlite")
	}
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&User{},
		&Profile{},
		&AcademicProfile{},
		&WorkProfile{},
		&CollegeProfile{},
		&Department{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
