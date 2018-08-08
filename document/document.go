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
	"time"

	"github.com/mezeipetister/document_api/pkg/common"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type query map[string]interface{}

// Document model
type Document struct {
	ID            objectid.ObjectID `bson:"_id"`            // Document ID
	Name          string            `bson:"name"`           // Document name, like title
	Description   string            `bson:"description"`    // Document short description what's this document is about
	File          string            `bson:"file"`           // Attached PDF file
	Folder        string            `bson:"folder"`         // Folder; logic manager
	Partners      []string          `bson:"partners"`       // Partner list, related partners
	LinkedFolders []Folders         `bson:"linked_folders"` // LinkedFolders; logical relations
	DueDate       string            `bson:"due_date"`       // DueDate; global duedate for each document; e.g. contract withdrawal time
	Tasks         []Task            `bson:"tasks"`          // Tasdks; contains related tasks
	Comments      []Comment         `bson:"comments"`       // Comments for team discussions
	Logs          []Log             `bson:"logs"`           // Logs contains changelog; auto generated
	IsRemoved     bool              `bson:"is_removed"`     // IsRemoved; boolt field for logical remove; true means deleted
	settings      *dbSettings
}

// Comment model
type Comment struct {
	ID          objectid.ObjectID // Comment ID for identification
	UID         objectid.ObjectID // UID => User ID
	Comment     string            // Comment body
	DateCreated time.Time         // DateCreated, time when the comment is created
}

// Log model
type Log struct {
	ID          objectid.ObjectID // ID for identification
	UID         objectid.ObjectID // UID => User ID
	Message     string            // Log message
	DateCreated time.Time         // DateCreated, time when the log was created
}

// Task model
type Task struct {
	ID          objectid.ObjectID
	Owner       objectid.ObjectID // Owner => USER ID
	Title       string
	Description string
	DateCreated string
	DueDate     string
	State       string
}

// Folders ...
type Folders struct {
	ID   objectid.ObjectID
	Name string
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

// Remove document
func (d *Document) Remove(ctx context.Context) {
	d.IsRemoved = true
	d.Update(ctx)
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

// SetComment insert a new comment
func (d *Document) SetComment(ctx context.Context, comment string) {
	newComment := &Comment{
		ID:          objectid.New(),
		Comment:     comment,
		DateCreated: time.Now(),
	}
	d.Comments = append(d.Comments, *newComment)
	d.Update(ctx)
}

// SetPartner insert a new partner
func (d *Document) SetPartner(ctx context.Context, partnerID string) {
	d.Partners = append(d.Partners, partnerID)
	d.Update(ctx)
}

// SetTask insert a new TASK
func (d *Document) SetTask(ctx context.Context, task Task) {
	d.Tasks = append(d.Tasks, task)
	d.Update(ctx)
}
