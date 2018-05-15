package db

import (
	model "Projects/document_api/src/model"
	"encoding/json"
	"fmt"
	"log"

	mgo "gopkg.in/mgo.v2"
)

// DAO struct
type DAO struct {
	Database, Collection string
}

func init() {
	println("Hello")
}

const server = "localhost"

var db *mgo.Database

// Connect to database
func (d *DAO) Connect() {
	session, err := mgo.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	db = session.DB(d.Database)
}

// Insert a document to database
func (d *DAO) Insert(document *model.Document) error {
	b, _ := json.Marshal(document)
	println(string(b))
	err := db.C(d.Collection).Insert(&document)
	return err
}
