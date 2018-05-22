package user

import (
	dbService "document_api/services/database"
)

// New : create and return a new user
func New(db *dbService.Interface) (Model, error) {
	return Model{datastore: *db}, nil
}

// Login ...
func (u *Model) Login(username, password string) (bool, error) {
	// TODO change it!!! Just for testing purpose!!!!!!

	// res := User{}
	// u.datastore.GetSession().DB("DEMO").C("doc1").Find(bson.M{"username": "mezeipetister"}).One(&res)

	// u.datastore.FindOne(bson.M{"username": "mezeipetister"}, &res)
	// fmt.Println(res.Username)

	// TODO: Its too slow!
	// result := checkPasswordHash(password, res.Password)
	// result := false
	// start := time.Now()
	// hash, _ := hashPassword(password)
	// finish := time.Now()
	// result = hash == res.Password
	// fmt.Println(finish.Sub(start))
	return true, nil
}

// Save the current user document to database
func (u *Model) Save() error {
	return u.datastore.CollectionInsert(u.document)
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

// SetFName ...
func (u *Model) SetFName(fname string) {
	u.document.FName = fname
}

// SetLName ...
func (u *Model) SetLName(lname string) {
	u.document.LName = lname
}

// SetEmail ...
func (u *Model) SetEmail(email string) {
	u.document.Email = email
}

// SetPassword ...
func (u *Model) SetPassword(password string) error {
	hash, error := hashPassword(password)
	if error == nil {
		u.document.Password = hash
		return nil
	}
	return error
}

// ResetPassword ...
func (u *Model) ResetPassword() string {
	// TODO
	return ""
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
