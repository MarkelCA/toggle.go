package flags

import (
	"strconv"
	"time"

	"github.com/markelca/toggles/pkg/storage"
)

const (
	DEFAULT_EXPIRATION_TIME = 5 * time.Minute
	redisKeyPrefix          = "flag:"
)

type FlagService interface {
	Get(key string) (bool, error)
	Create(f Flag) error
	Update(name string, value bool) error
	Exists(key string) (bool, error)
	List() ([]Flag, error)
	Delete(name string) error
}

type DefaultFlagService struct {
	repository  FlagRepository
	cacheClient storage.CacheClient
}

func NewFlagService(cacheClient storage.CacheClient, repository FlagRepository) FlagService {
	return DefaultFlagService{repository, cacheClient}
}

func (flagService DefaultFlagService) Get(key string) (bool, error) {
	cachedResult, err := flagService.cacheClient.Get(redisKeyPrefix + key)
	if err == nil {
		// We update the TTL on every successfull key access
		err = flagService.cacheClient.Expire(redisKeyPrefix+key, DEFAULT_EXPIRATION_TIME)
		if err != nil {
			return false, nil
		}
	} else if err == storage.Nil {
		value, err := flagService.repository.Get(key)
		if err == nil {
			flagService.cacheClient.Set(redisKeyPrefix+key, value, DEFAULT_EXPIRATION_TIME)
		}
		return value, err
	} else if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(cachedResult)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (flagService DefaultFlagService) Create(f Flag) error {
	exists, err := flagService.Exists(f.Name)
	if err != nil {
		return err
	} else if exists {
		return ErrFlagAlreadyExists
	}

	err = flagService.repository.Set(f.Name, f.Value)
	if err != nil {
		return err
	}

	return nil
}

func (flagService DefaultFlagService) Update(name string, value bool) error {
	exists, err := flagService.Exists(name)
	if err != nil {
		return err
	} else if !exists {
		return ErrFlagNotFound
	}
	err = flagService.repository.Set(name, value)
	if err != nil {
		return err
	} else {
		err = flagService.cacheClient.Delete(redisKeyPrefix + name)
		return err
	}
}

func (flagService DefaultFlagService) Exists(key string) (bool, error) {
	return flagService.repository.Exists(key)
}

func (flagService DefaultFlagService) List() ([]Flag, error) {
	flags, err := flagService.repository.List()
	if err != nil {
		return nil, err
	}
	return flags, nil
}

func (flagService DefaultFlagService) Delete(name string) error {
	if exists, err := flagService.Exists(name); err != nil {
		return err
	} else if !exists {
		return ErrFlagNotFound
	}
	if err := flagService.repository.Delete(name); err != nil {
		return err
	}
	if err := flagService.cacheClient.Delete(redisKeyPrefix + name); err != nil {
		return err
	}
	return nil
}
