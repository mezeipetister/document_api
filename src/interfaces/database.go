package interfaces

import "gopkg.in/mgo.v2/bson"

// Database interface
type Database interface {
	Save(collection string, i interface{}) error
	RemoveById(collection string, ID bson.ObjectId) error
	Find(collection string, i bson.M) interface{}
	FindAll()
	Close()
}
