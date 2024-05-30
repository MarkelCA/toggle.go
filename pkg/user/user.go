package user

import "github.com/markelca/toggles/pkg/security"

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Role      string
	Password  string
	ApiKey    string
}

func NewUser(userName, role, password string) (*User, error) {
	pwdHash, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}
	apiKey, err := security.GenerateAPIKey()
	if err != nil {
		return nil, err
	}
	return &User{
		UserName: userName,
		Role:     role,
		Password: pwdHash,
		ApiKey:   apiKey,
	}, nil
}

// //////////
// Errors
// //////////
type UserAuthError string

func (e UserAuthError) Error() string { return string(e) }

const ErrUserAuthenticationFailed = UserAuthError("toggles: User authentication failed")
const ErrApiKeyMismatch = UserAuthError("toggles: API key does not match")
