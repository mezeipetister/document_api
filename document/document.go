/*
 * Created on Fri Jul 20 2018
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

package document

import (
	"context"

	"github.com/mezeipetister/document_api/pkg/common"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type query map[string]interface{}

// Document model
type Document struct {
	ID          string `bson:"_id"`         // Document ID
	Title       string `bson:"title"`       // Document name, like title
	Description string `bson:"description"` // Document short description what's this document is about
	File        string `bson:"file"`        // Attached PDF file
	settings    *dbSettings
}

type dbSettings struct {
	dbName, collectionName string
	dbClient               *mongo.Client
}

// NewDocument return a new, empty document
func NewDocument(client *mongo.Client) *Document {
	return &Document{
		ID: objectid.New().Hex(),
		settings: &dbSettings{
			dbClient:       client,
			dbName:         common.Config.DB.DBName,
			collectionName: common.Config.DB.CollectionDocument,
		},
	}
}

// FindDocument ...
func FindDocument(ctx context.Context, client *mongo.Client, keyword string) *Document {
	d := NewDocument(client)
	filter := bson.NewDocument(bson.EC.String("name", keyword))
	err := d.settings.dbClient.Database(d.settings.dbName).Collection(d.settings.collectionName).FindOne(ctx, filter).Decode(d)
	if err != nil {
		panic(err)
	}
	return d
}

// FindDocuments ...
func FindDocuments(ctx context.Context, client *mongo.Client, keyword string) []*Document {
	d := NewDocument(client)
	filter := bson.NewDocument(bson.EC.String("title", keyword))
	cursor, err := d.settings.dbClient.Database(d.settings.dbName).Collection(d.settings.collectionName).Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	var result []*Document

	for cursor.Next(context.Background()) {
		d := NewDocument(client)
		cursor.Decode(d)
		result = append(result, d)
	}
	return result
}

// GetDocumentByID ...
func GetDocumentByID(ctx context.Context, client *mongo.Client, id string) *Document {
	d := NewDocument(client)
	filter := bson.NewDocument(bson.EC.String("_id", id))
	// filter := map[string]string{"_id": id}
	err := d.settings.dbClient.Database(d.settings.dbName).Collection(d.settings.collectionName).FindOne(ctx, filter).Decode(d)
	if err != nil {
		panic(err)
	}
	return d
}

// Save document
func (d *Document) Save() {
	_, err := d.settings.dbClient.Database(d.settings.dbName).Collection(d.settings.collectionName).InsertOne(context.Background(), d)
	if err != nil {
		panic(err)
	}
}

// Update document
func (d *Document) Update(ctx context.Context) {
	q := make(query)
	q["$set"] = d
	oid, _ := objectid.FromHex(d.ID)
	_, err := d.settings.dbClient.Database(d.settings.dbName).Collection(d.settings.collectionName).UpdateOne(
		ctx,
		oid,
		q,
	)
	if err != nil {
		panic(err)
	}
}
