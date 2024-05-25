package websocket

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/markelca/toggles/pkg/flags"
	"github.com/markelca/toggles/pkg/storage"
)

type Controller interface {
	HandleMessage(message []byte, client *Client)
}

type ControllerV2 struct {
	FlagService flags.FlagService
	CacheClient storage.CacheClient
}

func (controller ControllerV2) HandleMessage(message []byte, c *Client) {
	var action Action
	c.actionMarshaller.Unmarshal(message, &action)

	r := controller.RunAction(&action)

	responseBytes, err := json.Marshal(r)

	if err != nil {
		slog.Error(err.Error())
		return
	}

	if action.Type == ActionTypeUpdate {
		c.hub.broadcast <- responseBytes
	} else {
		clientResponse := ClientResponse{c, responseBytes}
		c.hub.response <- clientResponse
	}
}

func (controller ControllerV2) RunAction(action *Action) Response {
	var response Response
	switch action.Type {
	case ActionTypeGet:
		response = controller.GetV2(action)
	case ActionTypeUpdate:
	case ActionTypeCreate:
	case ActionTypeDelete:
	default:
		msg := fmt.Sprintf("Invalid action type (%v)", action.Type)
		response = Response{StatusBadRequest, msg}
	}

	switch response.Status {
	case StatusInternalServerError:
		slog.Error(fmt.Sprintf("[Request failed] (%v): %v", action.String(), response.String()))
	default:
		slog.Info("[Request received] " + action.String())
	}

	return response
}

func (controller ControllerV2) GetV2(action *Action) Response {
	if action.Flag == nil {
		flags, err := controller.FlagService.List()
		if err != nil {
			return Response{StatusInternalServerError, err}
		}
		return Response{StatusSuccess, flags}
	}
	flag, err := controller.FlagService.Get(*action.Flag)
	if err != nil {
		if err == flags.ErrFlagNotFound {
			return Response{StatusNotFound, err}
		}
		return Response{StatusInternalServerError, err}
	}
	return Response{StatusSuccess, flag}
}
