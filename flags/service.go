package flags

import (
	"strconv"
	"time"
	"github.com/markelca/toggles/storage"
)

const DEFAULT_EXPIRATION_TIME = 5 * time.Minute

type FlagService struct {
    repository FlagRepository
    cacheClient storage.CacheClient
}

func NewFlagService(cacheClient storage.CacheClient, repository FlagRepository) FlagService {
    return FlagService{repository, cacheClient}
}

func (flagService FlagService) Get(key string) (bool,error) {
    cachedResult,err := flagService.cacheClient.Get(key) 
    if err == nil {
        // We update the TTL on every successfull key access
        err = flagService.cacheClient.Expire(key,DEFAULT_EXPIRATION_TIME)
        if err != nil {
            return false,nil
        }
    } else if err == storage.Nil {
        value,err := flagService.repository.Get(key)
        if err == nil {
            flagService.cacheClient.Set(key,value,DEFAULT_EXPIRATION_TIME)
        }
        return value,err
    } else if err != nil{
        return false,err
    }

    result,err := strconv.ParseBool(cachedResult)
    if err != nil {
        return false,err
    }
    return result,nil
}

func (flagService FlagService) Create(f Flag) error {
    exists,err := flagService.Exists(f.Name)
    if err != nil {
        return err
    } else if exists {
        return ErrFlagAlreadyExists
    } 

    err = flagService.repository.Set(f.Name,f.Value)
    if err != nil {
        return err
    }

    return nil
}

func (flagService FlagService) Update(name string, value bool) error {
    exists,err := flagService.Exists(name)
    if err != nil {
        return err
    } else if !exists {
        return ErrFlagNotFound
    }
    err = flagService.repository.Set(name,value)
    if err != nil {
        return err
    } else {
        err = flagService.cacheClient.Delete(name)
        return err
    }
}

func (flagService FlagService) Exists(key string) (bool,error) {
    return flagService.repository.Exists(key)
}
 
func (flagService FlagService) List()([]Flag, error) {
    flags,err := flagService.repository.List()
    if err != nil {
        return nil,err
    }
    return flags,nil
}

