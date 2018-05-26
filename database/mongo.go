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

package database

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Model struct. Storing database MGO session.
type Model struct {
	session    *mgo.Session
	dbName     string
	collection string
}

// Interface : basic database driver is mgo.v2
type Interface interface {
	CloseSession() error
	CollectionInsert(interface{}) error
	RemoveDocumentByID(documentID bson.ObjectId)
	FindOne(i bson.M, resObj interface{}) error
	GetSession() *mgo.Session
}

// New database. Returns a database instanse with MGO Session.
func New(serverAddress, dbName, collection string) (Interface, error) {
	session, err := mgo.Dial(serverAddress)
	defer session.Close()
	if err != nil {
		return &Model{}, err
	}
	return &Model{
		session.Copy(),
		dbName,
		collection}, nil
}

// GetSession ...
func (db *Model) GetSession() *mgo.Session {
	return db.session
}

// CloseSession close MGO active session.
func (db *Model) CloseSession() error {
	db.session.Close()
	return nil
}

// FindOne ...
func (db *Model) FindOne(q bson.M, resObj interface{}) error {
	db.session.DB(db.dbName).C(db.collection).Find(q).One(resObj)
	return nil
}

// CollectionInsert ...
func (db *Model) CollectionInsert(i interface{}) error {
	err := db.session.DB(db.dbName).C(db.collection).Insert(i)
	return err
}

// RemoveDocumentByID ...
func (db *Model) RemoveDocumentByID(documentID bson.ObjectId) {
	db.session.DB(db.dbName).C(db.collection).Remove(bson.M{"_id": documentID})
}
