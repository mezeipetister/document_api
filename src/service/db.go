package service

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// Document model
type Document struct {
	Name        string    `bson:"name" json:"name"`
	SKU         string    `bson:"sku" json:"sku"`
	Description string    `bson:"description" json:"description"`
	File        string    `bson:"file" json:"file"`
	IsClosed    bool      `bson:"isclosed" json:"isclosed"`
	DueDate     time.Time `bson:"due_date" json:"due_date"`
}

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
