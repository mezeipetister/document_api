package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Interface : basic database driver is mgo.v2
type Interface interface {
	CloseSession() error
	CollectionInsert(interface{}) error
	RemoveDocumentById(documentID bson.ObjectId)
	FindOne(i bson.M, resObj interface{}) error
	GetSession() *mgo.Session
}
