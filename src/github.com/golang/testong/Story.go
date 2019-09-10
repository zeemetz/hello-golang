package main

import "github.com/jinzhu/gorm"

//Story entity
type Story struct {
	gorm.Model
	Title    string
	AuthorID int
	Author   *User
}

func init() {
	db().AutoMigrate(&Story{})
}
