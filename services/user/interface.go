package user

// User interface
type Interface interface {
	Save() error
	Remove() error
	Get() (user, error)
	Set(u user) error
	ResetPassword() (string, error)
}
