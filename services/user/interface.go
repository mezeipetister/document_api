package user

// User interface
type User interface {
	Save() error
	Remove() error
}
