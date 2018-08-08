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

package user

import (
	"context"
	"errors"

	"github.com/mezeipetister/document_api/pkg/common"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type query map[string]interface{}

// User model
type User struct {
	ID           objectid.ObjectID `bson:"_id"`           // Document ID
	Username     string            `bson:"username"`      // Document name, like title
	FirstName    string            `bson:"first_name"`    // First name of the user
	LastName     string            `bson:"last_name"`     // Last name of the user
	Email        string            `bson:"email"`         // User email address
	PasswordHash string            `bson:"password_hash"` // Password hash
	IsRemoved    bool              `bson:"is_removed"`    // Logical value
	settings     *dbSettings
}

type dbSettings struct {
	dbName, collectionName string
	dbClient               *mongo.Client
}

// NewUser returns a new user object
func NewUser(client *mongo.Client) *User {
	return &User{
		ID: objectid.New(),
		settings: &dbSettings{
			dbClient:       client,
			dbName:         common.Config.DB.DBName,
			collectionName: common.Config.DB.CollectionUser,
		},
	}
}

// Login user, return User object when true
func Login(client *mongo.Client, email, password string) (*User, error) {
	d := NewUser(client)
	filter := bson.NewDocument(bson.EC.String("email", email))
	err := d.settings.dbClient.Database(d.settings.dbName).Collection(d.settings.collectionName).FindOne(context.Background(), filter).Decode(d)
	if err != nil {
		panic(err)
	}
	if checkPasswordHash(password, d.PasswordHash) {
		return d, nil
	}
	return nil, errors.New("Password is incorrect")
}

// Remove document
func (d *User) Remove(ctx context.Context) {
	d.IsRemoved = true
	d.Update(ctx)
}

// Save document
func (d *User) Save() {
	_, err := d.settings.dbClient.Database(d.settings.dbName).Collection(d.settings.collectionName).InsertOne(context.Background(), d)
	if err != nil {
		panic(err)
	}
}

// Update document
func (d *User) Update(ctx context.Context) {
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

// SetPassword insert a new TASK
func (d *User) SetPassword(ctx context.Context, password string) {
	d.PasswordHash, _ = hashPassword(password)
	d.Update(ctx)
}

// ValidatePassword compares the stored and the sent password
func (d *User) ValidatePassword(password string) bool {
	return checkPasswordHash(password, d.PasswordHash)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
