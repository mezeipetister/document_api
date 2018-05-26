package database

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// New database. Returns a database instanse with MGO Session.
func New(serverAddress, dbName, collection string) (Interface, error) {
	session, err := mgo.Dial(serverAddress)
	defer session.Close()
	if err != nil {
		return Model{}, err
	}
	return Model{
		session.Copy(),
		dbName,
		collection}, nil
}

// GetSession ...
func (db Model) GetSession() *mgo.Session {
	return db.session
}

// CloseSession close MGO active session.
func (db Model) CloseSession() error {
	db.session.Close()
	return nil
}

// FindOne ...
func (db Model) FindOne(q bson.M, resObj interface{}) error {
	db.session.DB(db.dbName).C(db.collection).Find(q).One(resObj)
	return nil
}

// CollectionInsert ...
func (db Model) CollectionInsert(i interface{}) error {
	err := db.session.DB(db.dbName).C(db.collection).Insert(i)
	return err
}

// RemoveDocumentById ...
func (db Model) RemoveDocumentById(documentId bson.ObjectId) {
	db.session.DB(db.dbName).C(db.collection).Remove(bson.M{"_id": documentId})
}
