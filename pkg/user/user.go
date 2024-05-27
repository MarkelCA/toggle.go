package user

import "fmt"

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Role      string
	password  string
}

func (u User) String() string {
	return fmt.Sprintf("{UserName: %s, FirstName: %s, LastName: %s, Role: %s}", u.UserName, u.FirstName, u.LastName, u.Role)
}

// //////////
// Errors
// //////////
type UserError string

func (e UserError) Error() string { return string(e) }

const ErrUserAuthenticationFailed = UserError("toggles: User authentication failed")
