package document

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Document model
type Document struct {
	ID          bson.ObjectId `bson:"_id" json:"_id"`
	Name        string        `bson:"name" json:"name"`
	SKU         string        `bson:"sku" json:"sku"`
	Description string        `bson:"description" json:"description"`
	File        string        `bson:"file" json:"file"`
	IsClosed    bool          `bson:"isclosed" json:"isclosed"`
	DueDate     time.Time     `bson:"due_date" json:"due_date"`
}
