package flags

import "time"

type Flag struct {
    Name  string `json:"name"`
    Value bool `json:"value"`
}

type FlagRepository interface {
    Get(key string) (bool, error)
    Keys() ([]string, error)
    Exists(name string) (bool, error)
    Set(f Flag, expiration time.Duration) error
}

////////////
// Errors
////////////
type FlagError string
func (e FlagError) Error() string { return string(e) }

const FlagAlreadyExistsError = FlagError("toggles: Flag already exists") // nolint:errname
const FlagNotFoundError = FlagError("toggles: Flag not found") // nolint:errname
