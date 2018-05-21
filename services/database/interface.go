package database

import "gopkg.in/mgo.v2/bson"

// Interface : basic database driver is mgo.v2
type Interface interface {
	CloseSession() error
	CollectionInsert(i interface{}) error
	RemoveDocumentById(documentId bson.ObjectId)
}
