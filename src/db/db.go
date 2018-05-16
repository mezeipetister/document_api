package db

import (
	model "Projects/document_api/src/model"
	"log"

	mgo "gopkg.in/mgo.v2"
)

// DAO struct
type DAO struct {
	Database, Collection string
}

const server = "localhost"

var db *mgo.Database

// Connect to database
func (d *DAO) Connect() {
	session, err := mgo.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(d.Database)
}

// Insert a document to database
func (d *DAO) Insert(document *model.Document) error {
	err := db.C(d.Collection).Insert(&document)
	return err
}
