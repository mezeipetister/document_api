package database

import (
	mgo "gopkg.in/mgo.v2"
)

// database struct. Storing database MGO session.
type database struct {
	session *mgo.Session
}

// New database. Returns a database instanse with MGO Session.
func New(serverAddress, databaseName string) (Database, error) {
	session, err := mgo.Dial(serverAddress)
	defer session.Close()
	if err != nil {
		return &database{}, err
	}
	return &database{session.Copy()}, nil
}

// CloseSession close MGO active session.
func (db database) CloseSession() {
	db.session.Close()
}
