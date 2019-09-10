package main

import "github.com/jinzhu/gorm"

//Product is
type Product struct {
	gorm.Model
	Code  string
	Price int
}

func init() {
	db().AutoMigrate(&Product{})
}
