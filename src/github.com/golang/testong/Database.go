package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//CallBack is
type CallBack func(v interface{}) bool

var database *gorm.DB
var err error

func openDatabase() {
	database, err = gorm.Open("postgres", "host=192.168.100.36 user=postgres dbname=TMN-GO sslmode=disable password=truemoney2016")
	if err != nil {
		panic("failed to connect database")
	}
	database.DB().SetMaxOpenConns(100)
	database.DB().SetMaxIdleConns(1)
}

func init() {
	openDatabase()
}

func db() *gorm.DB {
	if database == nil {
		openDatabase()
	}
	return database
}

func insert(v interface{}) bool {
	if !db().NewRecord(v) {
		return false
	}
	db().Create(v)
	if db().NewRecord(v) {
		return false
	}
	return true
}

func update(v interface{}, callback CallBack) bool {
	tx := db().Begin().Debug()
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Where(v).First(v).Error; err != nil {
		tx.Rollback()
		return false
	}
	if !callback(v) {
		tx.Rollback()
		return false
	}
	tx.Save(v)
	tx.Commit()
	return true
}

//Select is
func Select(v interface{}) error {

	if err := db().Where(v).First(v).Error; err != nil {
		return err
	}
	return nil
}
