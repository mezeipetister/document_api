package user

// Interface : User Interface
type Interface interface {
	Save() error
	Remove() error
	Get() (User, error)
	Set(User) error
	ResetPassword() (string, error)
}
