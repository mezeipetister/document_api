package database

import mgo "gopkg.in/mgo.v2"

// Database struct. Storing database MGO session.
type Model struct {
	Session    *mgo.Session
	DBName     string
	Collection string
}
