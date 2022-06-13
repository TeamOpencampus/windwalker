package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//SetupDatabase connects to database and migrates schemas.
func SetupDatabase(dsn string, mode string) (*gorm.DB, error) {
	var dialect gorm.Dialector
	if mode == "DEV" {
		dialect = sqlite.Open("development.sqlite")
	} else {
		dialect = postgres.Open(dsn)
	}
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&User{}, &Profile{}, &AcademicProfile{}, &WorkProfile{}); err != nil {
		return nil, err
	}

	return db, nil
}
