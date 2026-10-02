package data

import (
	"NetGap/internal/server/data/model"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// db 单例模式
var (
	once sync.Once
	db   *gorm.DB
)

func InitServerDb(dbFile string) error {
	var err error
	db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Client{})
	return err
}

func GetServerDb() *gorm.DB {
	once.Do(func() {
		if db != nil {
			return
		}
		_ = InitServerDb("data.db")
	})
	return db
}
