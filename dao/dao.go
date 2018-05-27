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

// RemoveDocument ...
func (db *session) RemoveDocument(dbName, collection string, documentToRemove *interface{}) error {
	if err := db.session.DB(dbName).C(collection).Remove(documentToRemove); err != nil {
		return err
	}
	return nil
}

// UpdateDocument ...
func (db *session) UpdateDocument(dbName, collection string, selector, documentToUpdate *interface{}) error {
	if err := db.session.DB(dbName).C(collection).Update(selector, documentToUpdate); err != nil {
		return err
	}
	return nil
}

// UpdateDocumentByID ...
func (db *session) UpdateDocumentByID(dbName, collection, documentID string, documentToUpdate *interface{}) error {
	if err := db.session.DB(dbName).C(collection).Update(documentID, documentToUpdate); err != nil {
		return err
	}
	return nil
}

// FindDocumentOne ...
func (db *session) FindDocumentOne(dbName, collection string, searchQuery, result *interface{}) error {
	if err := db.session.DB(dbName).C(collection).Find(searchQuery).One(&result); err != nil {
		return err
	}
	return nil
}

// FindDocumentByID ...
func (db *session) FindDocumentByID(dbName, collection, documentID string, result *interface{}) error {
	if err := db.session.DB(dbName).C(collection).FindId(documentID).One(&result); err != nil {
		return err
	}
	return nil
}

// FindDocumentAll ...
func (db *session) FindDocumentAll(dbName, collection string, searchQuery *interface{}, result []*interface{}) error {
	if err := db.session.DB(dbName).C(collection).Find(searchQuery).All(result); err != nil {
		return err
	}
	return nil
}
