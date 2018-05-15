package model

import (
	"time"
)

// Document model
type Document struct {
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	File        string    `bson:"file" json:"file"`
	IsClosed    bool      `bson:"isclosed" json:"isclosed"`
	DueDate     time.Time `bson:"due_date" json:"due_date"`
}
