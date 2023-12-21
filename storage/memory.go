package storage

import (
	"fmt"
	"log"
	"time"
)


type MemoryRepository struct {}

func NewMemoryRepository() KeyValueStore {
    return MemoryRepository{}
}

var flagsStorage map[string]string = make(map[string]string)


func(r MemoryRepository) Keys() ([]string, error) {
    result := make([]string, 0)
    for i,_ := range flagsStorage {
        result = append(result, i)
    }
    return result,nil
}

func(r MemoryRepository) Get(key string) (string, error) {
    for k,v := range flagsStorage {
        if k == key {
            return v,nil
        }
    }
    return "",Nil
}

func (r MemoryRepository) Set(key string, value interface{}, expiration time.Duration) error {
    for k,v := range flagsStorage {
        if k == key {
            flagsStorage[key] = v
            return nil
        }
    }
    // If it doesn't find it it adds it
    log.Printf("fmt: %v",value)
    flagsStorage[key] = fmt.Sprint("%v",value)
    return nil
}

func (r MemoryRepository) Exists(name string) (bool,error) {
    if _, ok := flagsStorage[name]; ok {
        return true,nil
    }
    return false,nil
}

