package flags

type Flag struct {
    Name  string `json:"name"`
    Value bool `json:"value"`
}

type FlagRepository interface {
    Get(key string) (bool, error)
    List() ([]Flag, error)
    Create(f Flag) error
    Exists(name string) (bool, error)
    Update(name string, value bool) error
}

////////////
// Errors
////////////
type FlagError string
func (e FlagError) Error() string { return string(e) }

const FlagAlreadyExistsError = FlagError("toggles: Flag already exists") // nolint:errname
const FlagNotFoundError = FlagError("toggles: Flag not found") // nolint:errname
