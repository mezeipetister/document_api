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

package model

import (
	"context"
	"time"

	"github.com/mezeipetister/document_api/pkg/common"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type query map[string]interface{}

// Document model
type Document struct {
	ID          objectid.ObjectID `bson:"_id"`
	Name        string            `bson:"name"`
	Description string            `bson:"description"`
	File        string            `bson:"file"`
	Folder      string            `bson:"folder"`
	Partners    []Partner
	DueDate     string `bson:"due_date"`
	Tasks       []Task
	Comments    []Comment
	Logs        []Log
	IsRemoved   bool `bson:"is_removed"`
	settings    *dbSettings
}

type dbSettings struct {
	dbName, collectionName string
	dbClient               *mongo.Client
}

// NewDocument return a new, empty document
func NewDocument(client *mongo.Client) *Document {
	return &Document{
		ID: objectid.New(),
		settings: &dbSettings{
			dbClient:       client,
			dbName:         common.Config.DB.DBName,
			collectionName: common.Config.DB.CollectionDocument,
		},
	}
}

// Remove document
func (d *Document) Remove(ctx context.Context) {
	d.IsRemoved = true
	d.Update(ctx)
}

// Save document
func (d *Document) Save() {
	d.settings.dbClient.Database(d.settings.dbName).Collection(d.settings.collectionName).InsertOne(context.Background(), d)
}

// Update document
func (d *Document) Update(ctx context.Context) {
	q := make(query)
	q["$set"] = d
	_, err := d.settings.dbClient.Database(d.settings.dbName).Collection(d.settings.collectionName).UpdateOne(
		ctx,
		map[string]objectid.ObjectID{"_id": d.ID},
		q,
	)
	if err != nil {
		panic(err)
	}
}

// SetLog insert a new log message
func (d *Document) SetLog(ctx context.Context, msg string) {
	newLog := &Log{
		ID:          objectid.New(),
		Message:     msg,
		DateCreated: time.Now(),
	}
	d.Logs = append(d.Logs, *newLog)
	d.Update(ctx)
}
