package websocket

import (
	"encoding/json"
	"fmt"
)

type Status int

const (
	StatusSuccess Status = 200
	StatusCreated Status = 201

	StatusInternalServerError Status = 500

	StatusBadRequest Status = 400
	StatusNotFound   Status = 404
	StatusConflict   Status = 409
)

type Response struct {
	Status Status `json:"status"`
	Value  any    `json:"value"`
}

func (r Response) String() string {
	str, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("{\"status\": %v, \"value\": %v}", r.Status, r.Value)
	}
	return string(str)
}
