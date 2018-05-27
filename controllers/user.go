/*
 * Created on Sat May 26 2018
 * Copyright (c) 2018 Peter Mezei
 *
 * License AGPL v3.0
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>
 *
 * For more information please contact me
 * via github.com
 */

package controllers

// import (
// 	"fmt"
// 	"net/http"

// 	serviceDB "github.com/mezeipetister/document_api/database"
// 	serviceUser "github.com/mezeipetister/document_api/user"

// 	"github.com/julienschmidt/httprouter"
// 	"gopkg.in/mgo.v2/bson"
// )

// // CreateUser ...
// func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	db, err := serviceDB.New("localhost", "DEMO", "doc1")
// 	defer db.CloseSession()
// 	defer r.Body.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	u1, _ := serviceUser.New(&db)

// 	u1.Set(&serviceUser.User{
// 		ID:       bson.NewObjectId(),
// 		Email:    r.PostFormValue("email"),
// 		Username: r.PostFormValue("username"),
// 		LName:    r.PostFormValue("lname"),
// 		FName:    r.PostFormValue("fname"),
// 	})

// 	u1.SetPassword(r.PostFormValue("password"))

// 	error := u1.Save()
// 	if error != nil {
// 		fmt.Fprint(w, error)
// 		return
// 	}
// 	fmt.Fprint(w, u1.Get().ID)
// }

// // Login ...
// // TODO +ADD Authentaction middleware
// // TODO Add UID + Role Claims. Role is important!
// func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	db, err := serviceDB.New("localhost", "DEMO", "doc1")
// 	defer db.CloseSession()
// 	defer r.Body.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	u1, _ := serviceUser.New(&db)
// 	UID, err := u1.Login(r.PostFormValue("username"), r.PostFormValue("password"))
// 	if err == nil {
// 		token := createToken(UID, "Admin")
// 		// _, r := ValidateToken(token)
// 		// for k, v := range r {
// 		// 	fmt.Println(k + ": " + v.(string))
// 		// }
// 		fmt.Fprint(w, token)
// 		return
// 	}
// 	fmt.Fprint(w, err)
// }
