/*
 * Created on Sun May 27 2018
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

package dao

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	// Error message when mgo.Dial(string) fails
	errorMongoSessionError string = "Error while mgo dial."
)

// Model struct. Storing database MGO session.
type session struct {
	session *mgo.Session
}

// New database. Returns a database instance with MGO Session.
func New(serverAddress string) (*session, error) {
	if mgosession, err := mgo.Dial(serverAddress); err == nil {
		defer mgosession.Close()
		return &session{mgosession.Copy()}, nil
	}
	return &session{}, errors.New(errorMongoSessionError)
}

// CloseSession close MGO active session.
// As we use pointer *model we close the session at the right place.
func (db *session) CloseSession() {
	db.session.Close()
}

// RemoveDB removes the required database
func (db *session) RemoveDB(dbName string) error {
	if err := db.session.DB(dbName).DropDatabase(); err != nil {
		return err
	}
	return nil
}

// RemoveCollection from the given database
func (db *session) RemoveCollection(dbName, collectionName string) error {
	if err := db.session.DB(dbName).C(collectionName).DropCollection(); err != nil {
		return err
	}
	return nil
}

// InsertNewDocument return a callection with its built-in methods
func (db *session) InsertNewDocument(dbName, collectionName string, newDocument interface{}) error {
	if err := db.session.DB(dbName).C(collectionName).Insert(newDocument); err != nil {
		return err
	}
	return nil
}

// InsertNewDocuments return a callection with its built-in methods
func (db *session) InsertNewDocuments(dbName, collectionName string, newDocument ...interface{}) error {
	if err := db.session.DB(dbName).C(collectionName).Insert(newDocument...); err != nil {
		return err
	}
	return nil
}

// RemoveDocumentById ...
func (db *session) RemoveDocumentByID(dbName, collection, documentIDToRemove string) error {
	if err := db.session.DB(dbName).C(collection).Remove(bson.M{"_id": documentIDToRemove}); err != nil {
		return err
	}
	return nil
}

// UpdateDocument ...
// TODO: Error handling!
func (db *session) UpdateDocument(dbName, collection string, selector, documentToUpdate interface{}) error {
	Selector, _ := bson.Marshal(selector)
	DocumentToUpdate, _ := bson.Marshal(documentToUpdate)
	if err := db.session.DB(dbName).C(collection).Update(Selector, DocumentToUpdate); err != nil {
		return err
	}
	return nil
}

// UpdateDocumentByID ...
// TODO: Error handling!
func (db *session) UpdateDocumentByID(dbName, collection, documentID string, documentToUpdate interface{}) error {
	DocumentToUpdate, _ := bson.Marshal(documentToUpdate)
	if err := db.session.DB(dbName).C(collection).Update(bson.M{"_id": documentID}, DocumentToUpdate); err != nil {
		return err
	}
	return nil
}

// FindDocumentOne ...
// TODO: Error handling!
func (db *session) FindDocumentOne(dbName, collection string, searchQuery map[string]string, result interface{}) error {
	var SearchQuery bson.M
	for k, v := range searchQuery {
		SearchQuery[k] = v
	}
	if err := db.session.DB(dbName).C(collection).Find(SearchQuery).One(result); err != nil {
		return err
	}
	return nil
}

// FindDocumentByID ...
func (db *session) FindDocumentByID(dbName, collection string, documentID bson.ObjectId, result interface{}) error {
	if err := db.session.DB(dbName).C(collection).Find(bson.M{"_id": documentID}).One(result); err != nil {
		return err
	}
	return nil
}

// FindDocumentAll ...
func (db *session) FindDocumentAll(dbName, collection string, searchQuery bson.M, result interface{}) error {
	if err := db.session.DB(dbName).C(collection).Find(searchQuery).All(result); err != nil {
		return err
	}
	return nil
}
