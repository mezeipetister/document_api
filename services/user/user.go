package user

import database "document_api/services/database"

// New : create and return a new user
func New(session database.Database, db, collection string) (user, error) {
	settings := dbSettings{
		dbSession:  session,
		dbName:     db,
		collection: collection,
	}
	return user{db: settings}, nil
}

func (u *user) Save() {
	// err := u.db.dbSession.Session.DB(u.db.dbName).C(u.db.collection).Insert()
}

// GetByID : get a user by a given ID
func GetByID(id string) (user, error) {
	return user{}, nil
}

// SearchByName : search user by given name
func SearchByName(name string) ([]user, error) {
	return []user{}, nil
}

// SearchByEmail : search user by given email address
func SearchByEmail(email string) ([]user, error) {
	return []user{}, nil
}
