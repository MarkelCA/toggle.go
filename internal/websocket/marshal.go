package websocket

import (
	"encoding/json"
	"fmt"
)

type ActionMarshaller interface {
	Marshal(action *Action) ([]byte, error)
	Unmarshal(bytes []byte, action *Action) error
}

type JsonMarshaller struct{}

func (m JsonMarshaller) Marshal(action *Action) ([]byte, error) {
	return json.Marshal(action)
}

func (m JsonMarshaller) Unmarshal(bytes []byte, action *Action) error {
	return json.Unmarshal(bytes, action)
}

type BinaryMarshaller struct{}

func (m BinaryMarshaller) Marshal(action *Action) ([]byte, error) {
	return nil, fmt.Errorf("Not implemented yet")
}

func (m BinaryMarshaller) Unmarshal(bytes []byte, action *Action) error {
	return fmt.Errorf("Not implemented yet")
}
