package helpers

import (
	"gopkg.in/mgo.v2"
	"log"
)

type SozlukDB struct {
	Server   string
	Database string
}

var db *mgo.Database

func (m *SozlukDB) InitMongo() *mgo.Database {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
	return db
}

func GetDatabase() *mgo.Database {
	return db
}
