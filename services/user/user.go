package user

import dbService "document_api/services/database"

// New : create and return a new user
func New(db dbService.Interface) (*Model, error) {
	return &Model{datastore: db}, nil
}

// Save the current user document to database
func (u *Model) Save() error {
	u.datastore.CollectionInsert(u.document)
	return nil
}

// Remove the current user document from database
func (u *Model) Remove() error {
	u.datastore.RemoveDocumentById(u.Get().ID)
	return nil
}

// Get back the current user document
func (u *Model) Get() User {
	// err := u.db.dbSession.Session.DB(u.db.dbName).C(u.db.collection).Insert()
	return u.document
}

// Set a new user document to the current user object
func (u *Model) Set(userDocument *User) error {
	u.document = *userDocument
	return nil
}

// // GetByID : get a user by a given ID
// func GetByID(id string) (User, error) {
// 	return make(User), nil
// }

// // SearchByName : search user by given name
// func SearchByName(name string) ([]User, error) {
// 	return []User{}, nil
// }

// // SearchByEmail : search user by given email address
// func SearchByEmail(email string) ([]User, error) {
// 	return []User{}, nil
// }

// Helo Bello
