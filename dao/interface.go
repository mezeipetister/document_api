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

import "gopkg.in/mgo.v2/bson"

// DAO Interface
type DAO interface {
	CloseSession()
	RemoveDB(dbName string) error
	RemoveCollection(dbName, collectionName string) error
	InsertNewDocument(dbName, collectionName string, newDocument interface{}) error
	InsertNewDocuments(dbName, collectionName string, newDocument ...interface{}) error
	RemoveDocumentByID(dbName, collection string, documentIDToRemove bson.ObjectId) error
	UpdateDocument(dbName, collection string, selector, documentToUpdate interface{}) error
	UpdateDocumentByID(dbName, collection, documentID string, documentToUpdate interface{}) error
	FindDocumentOne(dbName, collection string, searchQuery bson.M, result interface{}) error
	FindDocumentByID(dbName, colelction string, documentID bson.ObjectId, result interface{}) error
	FindDocumentAll(dbName, collection string, searchQuery bson.M, result interface{}) error
}
