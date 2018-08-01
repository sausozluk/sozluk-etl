package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Topic struct {
	ID        bson.ObjectId   `bson:"_id" json:"_id"`
	TopicID   int             `bson:"id" json:"id"`
	Title     string          `bson:"title" json:"title"`
	Slug      string          `bson:"slug" json:"slug"`
	Locked    bool            `bson:"locked" json:"locked"`
	UpdatedAt time.Time       `bson:"updatedAt" json:"updatedAt"`
	CreatedAt time.Time       `bson:"createdAt" json:"createdAt"`
	Entries   []bson.ObjectId `bson:"entries" json:"entries"`
}
