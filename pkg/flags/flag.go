package flags

import (
	"encoding/json"
	"strconv"
)

type Flag struct {
	Name  string `json:"name"`
	Value bool   `json:"value"`
}

func ParseFlag(data any) (*Flag, error) {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var flag Flag
	if err = json.Unmarshal(jsonBody, &flag); err != nil {
		return nil, err
	}
	return &flag, nil
}

func (f Flag) String() string {
	jsonBody, err := json.Marshal(f)
	if err != nil {
		return "{Name: " + f.Name + ", Value: " + strconv.FormatBool(f.Value) + "}"
	}
	return string(jsonBody)
}

// //////////
// Errors
// //////////
type FlagError string

func (e FlagError) Error() string { return string(e) }

const ErrFlagAlreadyExists = FlagError("toggles: Flag already exists") // nolint:errname
const ErrFlagNotFound = FlagError("toggles: Flag not found")           // nolint:errname
