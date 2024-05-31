package user

import (
	"encoding/json"

	"github.com/markelca/toggles/pkg/security"
)

type User struct {
	UserName    string
	FirstName   string
	LastName    string
	Role        string
	Password    string
	ApiKey      string
	Permissions []string
}

func NewUser(userName, role, password string, permissions []string) (*User, error) {
	pwdHash, err := security.HashPassword(password)
	if err != nil {
		return nil, err
	}
	apiKey, err := security.GenerateAPIKey()
	if err != nil {
		return nil, err
	}
	return &User{
		UserName:    userName,
		Role:        role,
		Password:    pwdHash,
		ApiKey:      apiKey,
		Permissions: permissions,
	}, nil
}

func (u User) String() string {
	jsonBody, error := json.Marshal(u)
	if error != nil {
		return "{UserName: " + u.UserName + ", Role: " + u.Role + " Password: " + u.Password + " ApiKey: " + u.ApiKey + "}"
	}
	return string(jsonBody)
}

func (u User) ToPrettyStr() string {
	jsonBody, error := json.MarshalIndent(u, "", "\t")
	if error != nil {
		return "{\n\tUserName: " + u.UserName + ",\n\t Role: " + u.Role + ",\n\t Password: " + u.Password + ",\n\t ApiKey: " + u.ApiKey + "\n\t}"
	}
	return string(jsonBody)
}

// //////////
// Errors
// //////////
type UserAuthError string

func (e UserAuthError) Error() string { return string(e) }

const ErrUserAuthenticationFailed = UserAuthError("toggles: User authentication failed")
const ErrApiKeyMismatch = UserAuthError("toggles: API key does not match")
