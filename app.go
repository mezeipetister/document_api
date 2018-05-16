package main

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

type Model struct {
	Document Document
	db       Store
}

func (m Model) Save() {
	m.db.Save(m.Document)
}

type Store interface {
	Save(i interface{}) error
	Delete()
}

type store struct {
	db *mgo.Database
}

func (s store) Save(i interface{}) error {
	err := s.db.C("demoi").Insert(i)
	return err
}

func (s store) Delete() {
	// Do something
}

func NewDataStore() Store {
	var Server, Database string
	Server = "localhost"
	Database = "DEMO"

	session, err := mgo.Dial(Server)
	if err != nil {
		log.Fatal(err)
	}
	return &store{db: session.DB(Database)}
}

func main() {
	var d1 = &Model{
		Document: Document{
			Name: "HelloBello",
		},
		db: NewDataStore(),
	}
	d1.Save()
}
