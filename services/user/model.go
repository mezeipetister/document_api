package user

import dbService "document_api/services/database"

// Model represents a user document + datastore
type Model struct {
	document  uUser
	datastore dbService.Interface
}

type User struct {
	ID       string `bson:"_id" json:"_id"`
	FName    string `bson:"fname" json:"fname"`
	LName    string `bson:"lname" json:"lname"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}
