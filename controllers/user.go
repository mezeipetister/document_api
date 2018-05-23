package controllers

import (
	serviceDB "document_api/services/database"
	serviceUser "document_api/services/user"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

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
		Email:    "mezeipetister@gmail.com",
		Username: p.ByName("username"),
		LName:    "Mezei",
		FName:    "Péter",
	})

	u1.SetPassword("HelloBello")

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
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"UID": UID,
		})
		tokenString, errJWT := token.SignedString([]byte("HelloBello"))
		if errJWT != nil {
			fmt.Println(errJWT)
		}
		fmt.Fprint(w, tokenString)
		return
	}
	fmt.Fprint(w, err)
}
