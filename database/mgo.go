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
	"errors"

	mgo "gopkg.in/mgo.v2"
)

const (
	// Error message when mgo.Dial(string) fails
	errorMongoSessionError string = "Error while mgo dial."
)

// Interface : basic database driver is mgo.v2
type Interface interface {

	// Close active database session
	CloseSession()

	// Deprecated! Do not use it!
	// It will be removed in the near future.
	// TODO: Remove it from the next release!
	GetSession() *mgo.Session
}

// Model struct. Storing database MGO session.
type model struct {
	session *mgo.Session
}

// New database. Returns a database instanse with MGO Session.
func New(serverAddress string) (Interface, error) {
	if session, err := mgo.Dial(serverAddress); err == nil {
		defer session.Close()
		return &model{session.Copy()}, nil
	}
	return &model{}, errors.New(errorMongoSessionError)
}

// GetSession ...
// IMPORTANT: deprecated method. Use just built in methods.
// TODO: Remove it during the next iteration!
func (db *model) GetSession() *mgo.Session {
	return db.session
}

// CloseSession close MGO active session.
// As we use pointer *model we close the session at the right place.
func (db *model) CloseSession() {
	db.session.Close()
}
