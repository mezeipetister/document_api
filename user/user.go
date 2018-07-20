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
	"github.com/mezeipetister/document_api/dao"
	"gopkg.in/mgo.v2/bson"
)

// Model represents a user document + datastore
type Model struct {
	document       User
	datastore      dao.DAO
	dbName         string
	collectionName string
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
func New(db dao.DAO, dbName, collectionName string) (Model, error) {
	return Model{
		datastore:      db,
		dbName:         dbName,
		collectionName: collectionName,
		document:       User{ID: bson.NewObjectId()}}, nil
}

// FindUserByID find a user by ID
func (u *Model) FindUserByID(userID string) error {
	var user User
	err := u.datastore.FindDocumentByID(u.dbName, u.collectionName, bson.ObjectId(userID), user)
	return err
}

// Remove the current user document from database
func (u *Model) Remove() error {
	return u.datastore.RemoveDocumentByID(u.dbName, u.collectionName, u.document.ID)
}

// // Get back the current user document
// func (u *Model) Get() User {
// 	// err := u.db.dbSession.Session.DB(u.db.dbName).C(u.db.collection).Insert()
// 	return u.document
// }

// Save the current user document from database
func (u *Model) Save() error {
	// d, _ := bson.Marshal(u.document)
	return u.datastore.UpdateDocumentByID(u.dbName, u.collectionName, string(u.document.ID), u.document)
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
	_, error := hashPassword(password)
	if error == nil {
		u.document.Password = "H"
		return nil
	}
	return error
}

// // Login ...
// func (u *Model) Login(username, password string) (string, error) {
// 	res := User{}
// 	u.datastore.GetSession().DB("DEMO").C("doc1").Find(bson.M{"username": "mezeipetister"}).One(&res)

// 	u.datastore.FindOne(bson.M{"username": username}, &res)
// 	if checkPasswordHash(password, res.Password) {
// 		return res.ID.Hex(), nil
// 	}
// 	return "", errors.New("Authentication error")
// }
