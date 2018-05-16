package model

// User model
type User struct {
	ID       string `bson:"_id" json:"_id"`
	FName    string `bson:"fname" json:"fname"`
	LName    string `bson:"lname" json:"lname"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}
