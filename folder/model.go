package folder

// Folder model
type Folder struct {
	ID   string `bson:"_id" json:"_id"`
	Name string `bson:"name" json:"name"`
}
