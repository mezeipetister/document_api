package main

import (
	"document_api/controllers"
	"log"
	"net/http"
	"time"

	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {

	fmt.Println("start")

	router := httprouter.New()
	router.GET("/", controllers.Info)
	router.POST("/user/create", controllers.CreateUser)
	router.POST("/user/login", controllers.Login)
	router.GET("/getcookie", controllers.GetCookie)
	router.GET("/setcookie", controllers.SetCookie)

	log.Fatal(http.ListenAndServe(":8080", logger(router)))

	fmt.Println("End")
}

func demo() string {
	time.Sleep(3 * time.Second)
	return "OK"
}

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (r.RequestURI != "/user/login") && r.RequestURI != ("/user/create") {
			if len(r.Header.Get("x-user")) > 80 {
				userToken := r.Header.Get("x-user")
				valid := false
				valid, _ = controllers.ValidateToken(userToken)
				if !valid {
					http.Error(w, "Missing or wrong user key", http.StatusUnauthorized)
					return
				} else {
					h.ServeHTTP(w, r)
				}
			} else {
				http.Error(w, "Missing or wrong user key", http.StatusUnauthorized)
			}
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
