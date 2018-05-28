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
	"testing"

	"strconv"

	"gopkg.in/mgo.v2/bson"
)

const (
	testServer         = "localhost"
	testDBName         = "DEMO"
	testCollectionName = "DEMO"
)

var testDocumentID = bson.NewObjectId()

type testDocumentStruct struct {
	ID      bson.ObjectId `json:"_id" bson:"_id"`
	Name    string        `json:"Name" bson:"name"`
	Message string        `json:"message" bson:"message"`
}

var testDocumentToInsert = &testDocumentStruct{
	ID:      testDocumentID,
	Name:    "John Doe",
	Message: "Hello World",
}

func TestDAO(t *testing.T) {
	if d, err := New(testServer); err != nil {
		defer d.CloseSession()
		t.Error("Error while database initialize.")
	}
}

func TestInsertNewDocument(t *testing.T) {
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		if err := d.InsertNewDocument(testDBName, testCollectionName,
			&testDocumentToInsert); err != nil {
			t.Error("Error while inserting a new document.")
		}
	}
}

func TestInsertMultipleDocuments(t *testing.T) {
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		if err := d.InsertNewDocuments(testDBName, testCollectionName,
			&testDocumentStruct{bson.NewObjectId(), "Peti", "Multiple insert test"},
			&testDocumentStruct{bson.NewObjectId(), "Gabi", "Multiple insert test"},
			&testDocumentStruct{bson.NewObjectId(), "Kriszti", "Multiple insert test"}); err != nil {
			t.Errorf("Error while inserting a new documents. Error message: %s", err)
		}
	}
}

func TestFindByID(t *testing.T) {
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		var result testDocumentStruct
		if err := d.FindDocumentByID(testDBName, testCollectionName,
			testDocumentToInsert.ID, &result); err == nil {
			if result.Message != testDocumentToInsert.Message {
				t.Error("Found a document but not the inserted test document.")
			}
		} else {
			t.Errorf("Error while finding document by ID. Error message: %s", err)
		}
	}
}

func TestFindOne(t *testing.T) {
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		var result testDocumentStruct
		if err := d.FindDocumentOne(testDBName, testCollectionName,
			bson.M{"name": "John Doe"}, &result); err == nil {
			if result.Message != testDocumentToInsert.Message {
				t.Error("Found a document but not the inserted test document.")
			}
		} else {
			t.Errorf("Error while finding document by ID. Error message: %s", err)
		}
	}
}

func TestFindAll(t *testing.T) {
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		var result []testDocumentStruct
		if err := d.FindDocumentAll(testDBName, testCollectionName,
			bson.M{"message": "Multiple insert test"}, &result); err == nil {
			resultLen := len(result)
			if resultLen != 3 {
				t.Errorf("Found documents but not 3 but %s.", strconv.Itoa(resultLen))
			}
		} else {
			t.Errorf("Error while finding document by ID. Error message: %s", err)
		}
	}
}

// Important!
// Do not remove this test as the last one!
// This test removes demo collection at the end of the test circle.
func TestDropCollection(t *testing.T) {
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		if err := d.RemoveCollection(testDBName, testCollectionName); err != nil {
			t.Errorf("Error while finding document by ID. Error message: %s", err)
		}
	}
}
