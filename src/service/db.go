package service

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

// Database interface
type Database interface {
	Save(collection string, i interface{}) error
	Delete()
	Find()
	FindAll()
}

type db struct {
	db *mgo.Database
}

func (s db) Save(collection string, i interface{}) error {
	err := s.db.C(collection).Insert(i)
	return err
}

func (s db) Delete() {
	// Do something
}

func (s db) Find() {
	// Do something
}

func (s db) FindAll() {
	// Do something
}

// DB return new database layer
func DB(serverAddr, dbName string) Database {
	session, err := mgo.Dial(serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	return &db{db: session.DB(dbName)}
}
