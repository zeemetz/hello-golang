package main

//Version is
type Version struct {
	// gorm.Model
	ID             uint `gorm:"primary_key"`
	Detailversions []Detailversion
	VersionName    string
	ChangeLog      string
}

func init() {
	db().AutoMigrate(&Version{})
}
