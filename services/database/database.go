package database

import (
	mgo "gopkg.in/mgo.v2"
)

// Database struct. Storing database MGO session.
type Database struct {
	Session *mgo.Session
}

// New database. Returns a database instanse with MGO Session.
func New(serverAddress string) (*Database, error) {
	session, err := mgo.Dial(serverAddress)
	defer session.Close()
	if err != nil {
		return &Database{}, err
	}
	return &Database{session.Copy()}, nil
}

// CloseSession close MGO active session.
func (db Database) CloseSession() {
	db.Session.Close()
}
