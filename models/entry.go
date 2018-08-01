package models

import "gopkg.in/mgo.v2/bson"

type Entry struct {
	ID      bson.ObjectId   `bson:"_id" json:"_id"`
	EntryID int             `bson:"id" json:"id"`
	UserID  bson.ObjectId   `bson:"user" json:"user"`
	TopicID bson.ObjectId   `bson:"topic" json:"topic"`
	Text    string          `bson:"text" json:"text"`
	Up      []bson.ObjectId `bson:"up" json:"up"`
	Down    []bson.ObjectId `bson:"down" json:"down"`
}
