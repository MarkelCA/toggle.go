package storage

import "time"

const DEFAULT_EXPIRATION_TIME = 5 * time.Minute

type KeyValueStore interface {
	Get(key string) (string, error)
	Delete(key string) error
	Keys() ([]string, error)
	Exists(name string) (bool, error)
}

type CacheClient interface {
	KeyValueStore
	Set(key string, value any, expiration time.Duration) error
	Expire(key string, expiration time.Duration) error
	GetList(key string) ([]string, error)
	AppendToList(key string, expiration time.Duration, values ...any) error
	RemoveFromList(key string, values ...any) error
}

type KeyValueDBClient interface {
	KeyValueStore
	Set(key string, value any) error
}

// //////////
// Errors
// //////////
type StorageError string

func (e StorageError) Error() string { return string(e) }

const Nil = StorageError("toggles: Flag not found") // nolint:errname
