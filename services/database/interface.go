package database

// Database interface : basic database driver is mgo.v2
type Interface interface {
	CloseSession()
}
