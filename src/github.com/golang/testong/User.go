package main

import "github.com/jinzhu/gorm"

//User entity
type User struct {
	gorm.Model
	Name  string
	Email string
}

func init() {
	db().AutoMigrate(&User{})
}
