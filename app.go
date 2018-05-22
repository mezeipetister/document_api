package main

import (
	"document_api/controllers"
	"log"
	"net/http"

	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {

	fmt.Println("start")

	router := httprouter.New()
	router.GET("/", controllers.Info)
	router.GET("/createuser/:username", controllers.CreateUser)
	router.GET("/getcookie", controllers.GetCookie)
	router.GET("/setcookie", controllers.SetCookie)
	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("End")
}
