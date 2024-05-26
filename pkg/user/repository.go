package user

type UserRepository interface {
	FindAll() ([]*User, error)
	FindByUserName(userName string) (*User, error)
	Create(user *User) error
	Update(user *User) error
}
