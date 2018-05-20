package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/julienschmidt/httprouter"

	dbService "Projects/document_api/services/database"

	controllers "Projects/document_api/controllers"
)

func main() {

	db, err := dbService.New("localhost", "DEMO")
	defer db.CloseSession()

	if err != nil {
		fmt.Println(err)
	}

	router := httprouter.New()
	router.GET("/", controllers.Info)
	log.Fatal(http.ListenAndServe(":8080", router))
}
