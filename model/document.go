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

	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

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
	Changelog   []Log
	Status      bool `bson:"status"`
	dbClient    *mongo.Client
}

// Remove document
func (d *Document) Remove() {
	d.dbClient.Database("DEMO").Collection("A").DeleteOne(context.Background(), d)
}

// Save document
func (d *Document) Save() {
	d.dbClient.Database("DEMO").Collection("A").InsertOne(context.Background(), d)
}

// NewDocument return a new, empty document
func NewDocument(client *mongo.Client) *Document {
	return &Document{ID: objectid.New(), dbClient: client}
}
