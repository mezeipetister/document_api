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

package user

import (
	"errors"

	dbService "github.com/mezeipetister/document_api/database"

	"gopkg.in/mgo.v2/bson"
)

// Model represents a user document + datastore
type Model struct {
	document  User
	datastore dbService.Interface
}

// User model
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"_id"`
	Username string        `bson:"username" json:"username"`
	FName    string        `bson:"fname" json:"fname"`
	LName    string        `bson:"lname" json:"lname"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"password" json:"password"`
}

// Interface : User Interface
type Interface interface {
	Save() error
	Remove() error
	Get() User
	Set(User)
	SetFName(string)
	SetLName(string)
	SetEmail(string)
	SetPassword(string) error
	ResetPassword() string
	Login(string, string) (string, error)
}

// New : create and return a new user
func New(db *dbService.Interface) (Model, error) {
	return Model{datastore: *db}, nil
}

// Login ...
func (u *Model) Login(username, password string) (string, error) {
	// TODO change it!!! Just for testing purpose!!!!!!
	res := User{}
	u.datastore.GetSession().DB("DEMO").C("doc1").Find(bson.M{"username": "mezeipetister"}).One(&res)

	u.datastore.FindOne(bson.M{"username": username}, &res)
	if checkPasswordHash(password, res.Password) {
		return res.ID.Hex(), nil
	}
	return "", errors.New("Authentication error")
}

// Save the current user document to database
func (u *Model) Save() error {
	return u.datastore.CollectionInsert(u.document)
}

// Remove the current user document from database
func (u *Model) Remove() error {
	u.datastore.RemoveDocumentByID(u.Get().ID)
	return nil
}

// Get back the current user document
func (u *Model) Get() User {
	// err := u.db.dbSession.Session.DB(u.db.dbName).C(u.db.collection).Insert()
	return u.document
}

// Set a new user document to the current user object
func (u *Model) Set(userDocument *User) error {
	u.document = *userDocument
	return nil
}

// SetFName ...
func (u *Model) SetFName(fname string) {
	u.document.FName = fname
}

// SetLName ...
func (u *Model) SetLName(lname string) {
	u.document.LName = lname
}

// SetEmail ...
func (u *Model) SetEmail(email string) {
	u.document.Email = email
}

// SetPassword ...
func (u *Model) SetPassword(password string) error {
	hash, error := hashPassword(password)
	if error == nil {
		u.document.Password = hash
		return nil
	}
	return error
}

// ResetPassword ...
func (u *Model) ResetPassword() string {
	// TODO
	return ""
}
