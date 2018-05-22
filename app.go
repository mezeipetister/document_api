package main

import (
	serviceDB "document_api/services/database"
	serviceUser "document_api/services/user"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	db, err := serviceDB.New("localhost", "DEMO", "doc1")
	defer db.CloseSession()

	if err != nil {
		fmt.Println(err)
	}

	u1, err := serviceUser.New(db)
	u1.Set(&serviceUser.User{
		ID:    bson.NewObjectId(),
		Email: "mezeipetister@gmail.com",
	})
	u1.Save()
	// u1.Remove()

	// router := httprouter.New()
	// router.GET("/", controllers.Info)
	// router.GET("/getcookie", controllers.GetCookie)
	// router.GET("/setcookie", controllers.SetCookie)
	// log.Fatal(http.ListenAndServe(":8080", router))
}
