package storage

import (
	"time"

	"github.com/markelca/toggles/flags"
)

type MemoryRepository struct {}

func NewMemoryRepository() flags.FlagRepository {
    return MemoryRepository{}
}

var flagsStorage []flags.Flag = make([]flags.Flag,0)

func(r MemoryRepository) Keys() ([]string, error) {
    result := make([]string, len(flagsStorage))
    for i,f := range flagsStorage {
        result[i] = f.Name
    }
    return result,nil
}

func(r MemoryRepository) Get(key string) (bool, error) {
    for _,flag := range flagsStorage {
        if flag.Name ==  key {
            return flag.Value,nil
        }
    }
    return false,nil
}


func (r MemoryRepository) Set(f flags.Flag, expiration time.Duration) error {
    for i,currentFlag := range flagsStorage {
        if currentFlag.Name == f.Name {
            flagsStorage[i].Value = f.Value
            return nil
        }
    }
    // If it doesn't find it it adds it
    flagsStorage = append(flagsStorage,f)
    return nil
}

func (r MemoryRepository) Exists(name string) (bool,error) {
    for _,currentFlag := range flagsStorage {
        if currentFlag.Name == name {
            return true,nil
        }
    }
    return false,nil
}

