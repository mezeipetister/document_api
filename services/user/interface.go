package user

// Interface : User Interface
type Interface interface {
	Save() error
	Remove() error
	Get() User
	Set(User)
	SetFName(string)
	SetLName(string)
	SetEmail(string)
	SetPassword(string) error
	ResetPassword() string
	Login(string, string) (bool, error)
}
