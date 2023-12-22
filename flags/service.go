package flags

import (
	"strconv"
	"time"

	"github.com/markelca/toggles/storage"
)


type FlagService struct {
    Repository FlagRepository
    CacheClient storage.CacheClient
}

func NewFlagService(cacheClient storage.CacheClient, repository FlagRepository) FlagService {
    return FlagService{repository, cacheClient}
}

func (s FlagService) Get(key string) (bool,error) {
    // x,err := db.Get("new-login-pagee")
    // if err == flags.FlagNotFoundError {
    //     log.Printf("la puta vida")
    // }
    // log.Printf("looog: %v -> %v",x,err)

    cachedResult,err := s.CacheClient.Get(key) 
    if err == nil {
        // We update the TTL on every successfull key access
        expiration := time.Minute * 5
        err = s.CacheClient.Expire(key,expiration)
        if err != nil {
            return false,nil
        }
    } else if err == storage.Nil {
        return false,FlagNotFoundError
    } else if err != nil{
        return false,err
    }

    result,err := strconv.ParseBool(cachedResult)
    if err != nil {
        return false,err
    }
    return result,nil
}

func (s FlagService) Create(f Flag) error {
    exists,err := s.Exists(f.Name)
    if err != nil {
        return err
    } else if exists {
        return FlagAlreadyExistsError
    } 

    expiration := time.Minute * 5
    err = s.CacheClient.Set(f.Name,f.Value,expiration)
    if err != nil {
        return err
    }

    return nil
}

func (s FlagService) Update(name string, value bool) error {
    exists,err := s.Exists(name)
    if err != nil {
        return err
    } else if !exists {
        return FlagNotFoundError
    }
    expiration := time.Minute * 5
    err = s.CacheClient.Set(name,value,expiration)
    if err != nil {
        return err
    }
    return nil
}

func (s FlagService) Exists(key string) (bool,error) {
    return s.CacheClient.Exists(key)
}
 
func (s FlagService) List()([]Flag, error) {
    keys,err := s.CacheClient.Keys()
    if err != nil {
        return nil,err
    }

    result := make([]Flag,len(keys))

    for i,key := range keys {
        val,err := s.Get(key)
        if err != nil {
            return nil,err
        }
        result[i] = Flag{
            Name: key,
            Value: val,
        }
    }
    return result,nil
}

