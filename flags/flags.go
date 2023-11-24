package flags

type Flag struct {
    Name  string `json:"name"`
    Value bool `json:"value"`
}

type FlagExistsError struct{}

func (e *FlagExistsError) Error() string {
	return "This flag already exists"
}

var flags []Flag = make([]Flag,0)

func List()[]Flag {
    return flags
}

func Create(flag Flag) error {
    if Exists(flag) {
        return &FlagExistsError{}
    }
    flags = append(flags,flag)
    return nil
}

func Exists(flag Flag) bool {
    for _,currentFlag := range flags {
        if currentFlag.Name == flag.Name {
            return true
        }
    }
    return false
}
