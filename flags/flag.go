package flags

type Flag struct {
    Name  string `json:"name"`
    Value bool `json:"value"`
}

type FlagExistsError struct{}

func (e *FlagExistsError) Error() string {
	return "This flag already exists"
}

type FlagRepository interface {
    List() []Flag
    Create(f Flag) error
    Exists(name string) bool
}
