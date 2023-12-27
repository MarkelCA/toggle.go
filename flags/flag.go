package flags

import (
    "encoding/json"
)

type Flag struct {
    Name  string `json:"name"`
    Value bool `json:"value"`
}

func ParseFlag(data interface{}) (*Flag,error) {
    jsonBody,err := json.Marshal(data)
    if err != nil {
        return nil,err
    }
    var flag Flag
    if err = json.Unmarshal(jsonBody, &flag); err != nil {
        return nil,err
    }
    return &flag,nil
}


////////////
// Errors
////////////
type FlagError string
func (e FlagError) Error() string { return string(e) }

const ErrFlagAlreadyExists = FlagError("toggles: Flag already exists") // nolint:errname
const ErrFlagNotFound = FlagError("toggles: Flag not found") // nolint:errname
