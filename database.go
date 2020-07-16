package goarticles

import "github.com/jinzhu/gorm"

var db *gorm.DB

func Connect(username, password, host, dbname string) {
	var d, err = gorm.Open("mysql", username+":"+password+"@("+host+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}