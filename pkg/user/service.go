package user

import (
	"github.com/markelca/toggles/pkg/security"
	"github.com/markelca/toggles/pkg/storage"
)

type UserService interface {
	FindAll() ([]*User, error)
	FindByUserName(userName string) (*User, error)
	Create(user User) error
	Update(user *User) error
	Upsert(user User) error
	Authenticate(userName, password, apiKey string) (*User, error)
	HasPermission(userName, permission string) bool
	GetPermissions(userName string) ([]string, error)
	AddPermission(userName, permission string) error
	RemovePermission(userName, permission string) error
}

type DefaultUserService struct {
	repository  UserRepository
	cacheClient storage.CacheClient
}

func NewUserService(repository UserRepository, cacheClient storage.CacheClient) UserService {
	return DefaultUserService{repository, cacheClient}
}

func (service DefaultUserService) FindAll() ([]*User, error) {
	return service.repository.FindAll()
}

func (service DefaultUserService) FindByUserName(userName string) (*User, error) {
	return service.repository.FindByUserName(userName)
}

func (service DefaultUserService) Create(user User) error {
	return service.repository.Create(user)
}

func (service DefaultUserService) Update(user *User) error {
	return service.repository.Update(user)
}

func (service DefaultUserService) Upsert(user User) error {
	return service.repository.Upsert(user)
}

func (service DefaultUserService) GetPermissions(userName string) ([]string, error) {
	return service.repository.GetPermissions(userName)
}

// NOTE: This should be cached
func (service DefaultUserService) Authenticate(userName, password, apiKey string) (*User, error) {
	user, err := service.repository.FindByUserName(userName)
	if err != nil {
		return nil, err
	}
	if !security.CheckPasswordHash(password, user.Password) {
		return nil, ErrUserAuthenticationFailed
	}

	if user.ApiKey != apiKey {
		return nil, ErrApiKeyMismatch
	}

	return user, nil
}

// NOTE: This query should be cached
func (service DefaultUserService) HasPermission(userName, permission string) bool {
	user, err := service.repository.FindByUserName(userName)
	if err != nil {
		return false
	}
	for _, p := range user.Permissions {
		if p == permission {
			return true
		}
	}
	return false
}

// NOTE: This should update permission caches
func (service DefaultUserService) AddPermission(userName, permission string) error {
	return service.repository.AddPermission(userName, permission)
}

// NOTE: This should update permission caches
func (service DefaultUserService) RemovePermission(userName, permission string) error {
	return service.repository.RemovePermission(userName, permission)
}
