package websocket

type CommandType string
const (
    CommandTypeGet    = "get"
    CommandTypeUpdate = "update"
    CommandTypeCreate = "create"
    CommandTypeDelete = "delete"
)

type Command struct {
    Command string `json:"command"`
    Data interface{} `json:"data"`
    broadcast bool
    emmiter *Client
}

