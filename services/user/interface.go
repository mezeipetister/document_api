package user

// Interface : User Interface
type Interface interface {
	Save() error
	Remove() error
	Get() User
	Set(User)
	ResetPassword() (string, error)
}
