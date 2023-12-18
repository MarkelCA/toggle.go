package storage

import "time"

type CacheClient interface {
    Get(key string) (bool, error)
    Keys() ([]string, error)
    Exists(name string) (bool, error)
    Set(key string, value interface{}, expiration time.Duration) error
}
