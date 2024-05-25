package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/markelca/toggles/pkg/flags"
)

type ActionType int

const (
	ActionTypeGet ActionType = iota
	ActionTypeUpdate
	ActionTypeCreate
	ActionTypeDelete
)

func (a ActionType) String() (string, error) {
	switch a {
	case ActionTypeGet:
		return "get", nil
	case ActionTypeUpdate:
		return "update", nil
	case ActionTypeCreate:
		return "create", nil
	case ActionTypeDelete:
		return "delete", nil
	}
	return "", fmt.Errorf("ActionType not found (%v)", a)
}

type Action struct {
	Type  ActionType `json:"action"`
	Flag  *string    `json:"flag"`
	Value *bool      `json:"value"`
}

func (a Action) toFlag() (*flags.Flag, error) {
	if a.Flag == nil {
		return nil, fmt.Errorf("Flag is required")
	}
	if a.Value == nil {
		return nil, fmt.Errorf("Value is required")
	}
	return &flags.Flag{Name: *a.Flag, Value: *a.Value}, nil
}

func (a Action) String() string {
	str, err := json.Marshal(a)
	if err != nil {
		return fmt.Sprintf("{\"type\": %v, \"flag\": %v, \"value\": %v}", a.Type, a.Flag, a.Value)
	}
	return string(str)
}
