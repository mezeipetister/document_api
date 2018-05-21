package database

import (
	mgo "gopkg.in/mgo.v2"
)

// New database. Returns a database instanse with MGO Session.
func New(serverAddress, dbName, collection string) (Interface, error) {
	session, err := mgo.Dial(serverAddress)
	defer session.Close()
	if err != nil {
		return &Model{}, err
	}
	return &Model{
		session.Copy(),
		dbName,
		collection}, nil
}

// CloseSession close MGO active session.
func (db *Model) CloseSession() {
	db.Session.Close()
}
