package main

import (
	services "Projects/document_api/src/services"

	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	db := services.NewDB("localhost", "DEMO")

	// d1 := &models.Document{
	// 	ID:   bson.NewObjectId(),
	// 	Name: "Mizu? :)",
	// }

	// db.Save("demoi", &d1)

	// db.RemoveById("demoi", bson.ObjectIdHex("5afca055d57dd37b39830a0f"))

	result := db.Find("demoi", bson.M{"name": "Mizu?"})
	fmt.Println(result)
}
