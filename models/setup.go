package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//SetupDatabase connects to database and migrates schemas.
func SetupDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
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
