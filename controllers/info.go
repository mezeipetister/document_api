package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"

	serviceDB "document_api/services/database"
	serviceUser "document_api/services/user"
)

// Info controller
func Info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello!\n")
}

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := serviceDB.New("localhost", "DEMO", "doc1")
	defer db.CloseSession()
	defer r.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	u1, _ := serviceUser.New(&db)

	u1.Set(&serviceUser.User{
		ID:       bson.NewObjectId(),
		Email:    "mezeipetister@gmail.com",
		Username: p.ByName("username"),
		LName:    "Mezei",
		FName:    "Péter",
	})

	error := u1.Save()
	if error != nil {
		fmt.Fprint(w, error)
		return
	}
	fmt.Fprint(w, u1.Get().ID)
}

// SetCookie ...
func SetCookie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set Cookie
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "Name", Value: "Peti", Expires: expiration}
	http.SetCookie(w, &cookie)
	// Get Cookie
	fmt.Fprint(w, "Cookie Set\n")
}

// GetCookie ...
func GetCookie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, error := r.Cookie("Name")
	if error != nil {
		fmt.Fprint(w, "No cookie found")
	}
	// Get Cookie
	fmt.Fprint(w, cookie.Value)
}
