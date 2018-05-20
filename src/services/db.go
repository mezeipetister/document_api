package services

import (
	interfaces "Projects/document_api/src/interfaces"
	"log"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"
)

type db struct {
	db *mgo.Database
}

// Save object to DB
func (d db) Save(collection string, i interface{}) error {
	err := d.db.C(collection).Insert(i)
	return err
}

// Remove object from DB (Collection) by ID
func (d db) RemoveById(collection string, ID bson.ObjectId) error {
	// Delete record
	err := d.db.C(collection).Remove(bson.M{"_id": ID})
	if err != nil {
		panic(err)
	}
	return err
}

func (d db) Find(collection string, i bson.M) interface{} {
	var result interface{}
	err := d.db.C(collection).Find(i).One(result)
	if err != nil {
		panic(err)
	}
	return result
}

func (d db) FindAll() {
	// Do something
}

func (d db) Close() {
	d.db.Session.Close()
}

// NewDB return new database layer
func NewDB(serverAddr, dbName string) interfaces.Database {
	session, err := mgo.Dial(serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	// Error check on every access
	session.SetSafe(&mgo.Safe{})
	return &db{db: session.DB(dbName)}
}
