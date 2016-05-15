package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Sqlite ...
func Sqlite() *gorm.DB {
	db, err := gorm.Open("sqlite3", DbDatabase)
	if err != nil {
		panic(err)
	}
	return db
}

// Postgres ...
func Postgres() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		DbHost, DbUsername, DbDatabase, DbSSL, DbPassword))
	if err != nil {
		panic(err)
	}
	return db
}

// Mysql ...
func Mysql() *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", DbUsername, DbPassword, DbDatabase))
	if err != nil {
		panic(err)
	}
	return db
}