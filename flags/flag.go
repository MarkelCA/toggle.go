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

const FlagAlreadyExistsError = FlagError("toggles: Flag already exists") // nolint:errname
const FlagNotFoundError = FlagError("toggles: Flag not found") // nolint:errname
