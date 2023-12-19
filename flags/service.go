package flags

import (
	"strconv"
	"github.com/markelca/toggles/storage"
)


type FlagService struct {
    Cache storage.CacheClient
}

func NewFlagService(r storage.CacheClient) FlagService {
    return FlagService{r}
}

func (s FlagService) Get(key string) (bool,error) {
    cachedResult,err := s.Cache.Get(key) 
    if err == storage.Nil {
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

    err = s.Cache.Set(f.Name,f.Value,0)
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
    err = s.Cache.Set(name,value,0)
    if err != nil {
        return err
    }
    return nil
}

func (s FlagService) Exists(key string) (bool,error) {
    return s.Cache.Exists(key)
}
 
func (s FlagService) List()([]Flag, error) {
    keys,err := s.Cache.Keys()
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

