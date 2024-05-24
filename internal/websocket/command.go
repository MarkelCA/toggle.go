package websocket

import (
	"encoding/json"
	"fmt"
)

type CommandType string

const (
	CommandTypeGet    = "get"
	CommandTypeUpdate = "update"
	CommandTypeCreate = "create"
	CommandTypeDelete = "delete"
)

type Command struct {
	Command   string      `json:"command"`
	Data      interface{} `json:"data"`
	broadcast bool
	emmiter   *Client
}

func (c Command) String() string {
	str, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("{\"command\": %v, \"data\": %v}", c.Command, c.Data)
	}
	return string(str)
}
