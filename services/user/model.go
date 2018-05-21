package user

import (
	dbService "document_api/services/database"

	"gopkg.in/mgo.v2/bson"
)

// Model represents a user document + datastore
type Model struct {
	Document  User
	Datastore dbService.Model
}

// User model
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"_id"`
	FName    string        `bson:"fname" json:"fname"`
	LName    string        `bson:"lname" json:"lname"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
}
