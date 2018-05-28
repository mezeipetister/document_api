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

type dao interface {
	CloseSession()
	RemoveDB(dbName string) error
	RemoveCollection(dbName, collectionName string) error
	IntertNewDocument(dbName, collectionName string, newDocument *interface{}) error
	RemoveDocument(dbName, collection string, documentToRemove *interface{}) error
	UpdateDocument(dbName, collection string, selector, documentToUpdate *interface{}) error
	UpdateDocumentById(dbName, collection, documentID string, documentToUpdate *interface{}) error
	FindDocumentOne(dbName, collection string, searchQuery, result *interface{}) error
	FindDocumentByID(dbName, colelction, documentID string, result *interface{}) error
	FindDocumentAll(dbName, collection string, searchQuery *interface{}, result []*interface{}) error
}