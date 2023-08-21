package repositories

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() *Database {
	dsn := "data.db"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &Database{DB: db}
}
