package user

import database "document_api/services/database"

type user struct {
	ID       string `bson:"_id" json:"_id"`
	FName    string `bson:"fname" json:"fname"`
	LName    string `bson:"lname" json:"lname"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
	db       dbSettings
}

type dbSettings struct {
	dbSession  database.Database
	dbName     string
	collection string
}
