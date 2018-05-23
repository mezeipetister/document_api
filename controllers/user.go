package controllers

import (
	serviceDB "document_api/services/database"
	serviceUser "document_api/services/user"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

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
		Email:    r.PostFormValue("email"),
		Username: r.PostFormValue("username"),
		LName:    r.PostFormValue("lname"),
		FName:    r.PostFormValue("fname"),
	})

	u1.SetPassword(r.PostFormValue("password"))

	error := u1.Save()
	if error != nil {
		fmt.Fprint(w, error)
		return
	}
	fmt.Fprint(w, u1.Get().ID)
}

// Login ...
// TODO +ADD Authentaction middleware
// TODO Add UID + Role Claims. Role is important!
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	db, err := serviceDB.New("localhost", "DEMO", "doc1")
	defer db.CloseSession()
	defer r.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	u1, _ := serviceUser.New(&db)
	UID, err := u1.Login(r.PostFormValue("username"), r.PostFormValue("password"))
	if err == nil {
		token := createToken(UID, "Admin")
		// _, r := ValidateToken(token)
		// for k, v := range r {
		// 	fmt.Println(k + ": " + v.(string))
		// }
		fmt.Fprint(w, token)
		return
	}
	fmt.Fprint(w, err)
}
