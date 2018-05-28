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

var TD1 = &testDocumentStruct{bson.NewObjectId(), "A1", "Multiple insert test 1"}
var TD2 = &testDocumentStruct{bson.NewObjectId(), "A2", "Multiple insert test 2"}
var TD3 = &testDocumentStruct{bson.NewObjectId(), "A3", "Multiple insert test 3"}

func TestInsertMultipleRandomDocuments(t *testing.T) {
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		if err := d.InsertNewDocuments(testDBName, testCollectionName,
			TD1,
			TD2,
			TD3,
		); err != nil {
			t.Errorf("Error while inserting a new documents. Error message: %s", err)
		}
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

// We expect to be at least one document with the field "name":"John Doe"
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

// We expect to insert a multiple document insert test before.
// We expect to be exactly 3 documents in the collection,
// with the field "message":"Multiple insert test".
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

// For testing updateDocumentById we use the TD1 document, and rename it.
func TestUpdateDocumentByID(t *testing.T) {
	const newDocumentName = "A1.V2"
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		var dV1 testDocumentStruct
		var dV2 testDocumentStruct
		if err := d.FindDocumentByID(testDBName, testCollectionName, TD1.ID, &dV1); err != nil {
			t.Errorf("Error while finding document by ID. Error message: %s", err)
		}
		if err := d.UpdateDocumentByID(testDBName, testCollectionName, TD1.ID,
			bson.M{"name": newDocumentName}); err != nil {
			t.Errorf("Error while updating document by ID. Error message: %s", err)
		}
		if err := d.FindDocumentByID(testDBName, testCollectionName, TD1.ID, &dV2); err != nil {
			t.Errorf("Error while finding document by ID. Error message: %s", err)
		}
		if dV2.Name != newDocumentName {
			t.Errorf("Document updated, but after checking it again, the modification not found. Before: %s, After: %s", dV1.Name, dV2.Name)
		}
	}
}

// For testing updateDocument we use the TD2 document, and rename it.
func TestUpdateDocument(t *testing.T) {
	const newDocumentName = "A2.V2"
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		var dV1 testDocumentStruct
		var dV2 testDocumentStruct
		if err := d.FindDocumentByID(testDBName, testCollectionName, TD2.ID, &dV1); err != nil {
			t.Errorf("Error while finding document by ID. Error message: %s", err)
		}
		if err := d.UpdateDocument(testDBName, testCollectionName,
			bson.M{"name": "A2"},
			bson.M{"name": newDocumentName}); err != nil {
			t.Errorf("Error while updating document by ID. Error message: %s", err)
		}
		if err := d.FindDocumentByID(testDBName, testCollectionName, TD2.ID, &dV2); err != nil {
			t.Errorf("Error while finding document by ID. Error message: %s", err)
		}
		if dV2.Name != newDocumentName {
			t.Errorf("Document updated, but after checking it again, the modification not found. Before: %s, After: %s", dV1.Name, dV2.Name)
		}
	}
}

func TestRemoveDocumentByID(t *testing.T) {
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		if err := d.RemoveDocumentByID(testDBName, testCollectionName, testDocumentToInsert.ID); err != nil {
			t.Errorf("Error during removing a document. Error message: %s", err)
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
			t.Errorf("Error while dropping collection. Error message: %s", err)
		}
	}
}

func TestDropDB(t *testing.T) {
	if d, err := New(testServer); err == nil {
		defer d.CloseSession()
		if err := d.RemoveDB(testDBName); err != nil {
			t.Errorf("Error while dropping DB. Error message: %s", err)
		}
	}
}
