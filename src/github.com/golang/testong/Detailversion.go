package main

//Detailversion is
type Detailversion struct {
	// gorm.Model
	// Version   Version
	ID        uint `gorm:"primary_key"`
	VersionID int
	From      string
	To        string
	Filesize  int
}

func init() {
	db().AutoMigrate(&Detailversion{})
}
