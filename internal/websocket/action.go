package websocket

import (
	"fmt"
)

type ActionType int

const (
	Get ActionType = iota
	Update
	Create
	Delete
)

func (a ActionType) String() (string, error) {
	switch a {
	case Get:
		return "get", nil
	case Update:
		return "update", nil
	case Create:
		return "create", nil
	case Delete:
		return "delete", nil
	}
	return "", fmt.Errorf("ActionType not found (%v)", a)
}

type Action struct {
	Action ActionType `json:"action"`
	Flag   *string    `json:"flag"`
	Value  *bool      `json:"value"`
}
