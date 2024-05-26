package user

import "fmt"

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Role      string
}

func (u User) String() string {
	return fmt.Sprintf("{UserName: %s, FirstName: %s, LastName: %s, Role: %s}", u.UserName, u.FirstName, u.LastName, u.Role)
}
