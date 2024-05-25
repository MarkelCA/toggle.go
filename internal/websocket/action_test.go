package websocket

import (
	"encoding/json"
	"testing"
)

func TestActionUnmarshal(t *testing.T) {
	var a Action = Action{}
	str := `{"action":"get","flag":"flag1"}`
	if err := json.Unmarshal([]byte(str), &a); err != nil {
		t.Errorf("Error: %v", err)
	}

	if v, _ := a.Type.String(); v != "get" {
		t.Errorf("Error: %v", a.Type)
	}

	str = `{"action":"update","flag":"flag1","value":true}`
	if err := json.Unmarshal([]byte(str), &a); err != nil {
		t.Errorf("Error: %v", err)
	}

	if v, _ := a.Type.String(); v != "update" {
		t.Errorf("Error: %v", a.Type)
	}
}
