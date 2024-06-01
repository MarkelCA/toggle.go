package user

import "github.com/markelca/toggles/pkg/storage"

type UserService struct {
	repository  UserRepository
	cacheClient storage.CacheClient
}
