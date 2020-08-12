package goarticles

import "github.com/jinzhu/gorm"

type Database struct {
	Connection *gorm.DB
}

func NewDatabase(username, password, host, dbname string) *Database {
	db := &Database{}
	db.Connect(username, password, host, dbname)
	return db
}

func (db *Database) Connect(username, password, host, dbname string) {
	connection, err := gorm.Open("mysql", username+":"+password+"@("+host+")/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.Connection = connection
}
