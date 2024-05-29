package user

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Role      string
	Password  string
}

// //////////
// Errors
// //////////
type UserError string

func (e UserError) Error() string { return string(e) }

const ErrUserAuthenticationFailed = UserError("toggles: User authentication failed")
