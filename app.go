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
	DB       Store
}

func (m Model) Save() {
	m.DB.Save(m.Document)
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
		DB: NewDataStore(),
	}
	d1.Save()
}

// package main

// import (
// 	db "Projects/document_api/src/db"
// 	model "Projects/document_api/src/model"
// )

// func main() {
// 	document := &model.Document{
// 		Name:        "Demo document",
// 		Description: "Kriszti vagyok, szép vagyok.",
// 	}

// 	document.Name = "Demo document"

// 	db := new(db.DAO)
// 	db.Database = "DEMO"
// 	db.Collection = "documentstore"
// 	db.Connect()
// 	db.Insert(document)
// }

// // package dao

// // import (
// // 	"log"

// // 	. "github.com/mlabouardy/movies-restapi/models"
// // 	mgo "gopkg.in/mgo.v2"
// // 	"gopkg.in/mgo.v2/bson"
// // )

// // type MoviesDAO struct {
// // 	Server   string
// // 	Database string
// // }

// // var db *mgo.Database

// // const (
// // 	COLLECTION = "movies"
// // )

// // // Establish a connection to database
// // func (m *MoviesDAO) Connect() {
// // 	session, err := mgo.Dial(m.Server)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	db = session.DB(m.Database)
// // }

// // // Find list of movies
// // func (m *MoviesDAO) FindAll() ([]Movie, error) {
// // 	var movies []Movie
// // 	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
// // 	return movies, err
// // }

// // // Find a movie by its id
// // func (m *MoviesDAO) FindById(id string) (Movie, error) {
// // 	var movie Movie
// // 	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
// // 	return movie, err
// // }

// // // Insert a movie into database
// // func (m *MoviesDAO) Insert(movie Movie) error {
// // 	err := db.C(COLLECTION).Insert(&movie)
// // 	return err
// // }

// // // Delete an existing movie
// // func (m *MoviesDAO) Delete(movie Movie) error {
// // 	err := db.C(COLLECTION).Remove(&movie)
// // 	return err
// // }

// // // Update an existing movie
// // func (m *MoviesDAO) Update(movie Movie) error {
// // 	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
// // 	return err
// // }
