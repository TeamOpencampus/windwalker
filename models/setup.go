package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//SetupDatabase connects to database and migrates schemas.
func SetupDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&User{}, &Profile{}, &AcademicProfile{}, &WorkProfile{}); err != nil {
		return nil, err
	}

	return db, nil
}
