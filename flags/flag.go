package flags

type Flag struct {
    Name  string `json:"name"`
    Value bool `json:"value"`
}


////////////
// Errors
////////////
type FlagError string
func (e FlagError) Error() string { return string(e) }

const ErrFlagAlreadyExists = FlagError("toggles: Flag already exists") // nolint:errname
const ErrFlagNotFound = FlagError("toggles: Flag not found") // nolint:errname
