package main

import (
	model "Projects/document_api/src/model"
	service "Projects/document_api/src/service"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	db := service.DB("localhost", "DEMO")

	d1 := &model.Document{
		ID:   bson.NewObjectId(),
		Name: "Mizu?",
	}

	db.Save("demoi", d1)
}
