package websocket

import (
	"encoding/json"
	"fmt"
	"log"
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
	log.Println(message)
	var action Action
	c.actionMarshaller.Unmarshal(message, &action)
	// action.Type =

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
		response = controller.Update(action)
	case ActionTypeCreate:
		response = controller.Create(action)
	case ActionTypeDelete:
		response = controller.Delete(action)
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

func (controller ControllerV2) Delete(action *Action) Response {
	key := fmt.Sprintf("%v", action.Flag)
	err := controller.FlagService.Delete(key)
	if err != nil {
		if err == flags.ErrFlagNotFound {
			return Response{StatusNotFound, nil}
		}
		return Response{StatusInternalServerError, nil}
	}
	return Response{StatusSuccess, nil}
}

func (controller ControllerV2) Create(action *Action) Response {
	flag, err := action.toFlag()
	if err != nil {
		return Response{StatusBadRequest, err}
	}
	err = controller.FlagService.Create(*flag)
	if err != nil {
		if err == flags.ErrFlagAlreadyExists {
			return Response{StatusConflict, err}
		}
		return Response{StatusInternalServerError, err}
	}
	return Response{StatusCreated, flag}
}

func (controller ControllerV2) Update(action *Action) Response {
	flag, err := action.toFlag()
	if err != nil {
		return Response{StatusBadRequest, err}
	}
	err = controller.FlagService.Update(flag.Name, flag.Value)
	if err != nil {
		if err == flags.ErrFlagNotFound {
			return Response{StatusNotFound, err}
		}
		return Response{StatusInternalServerError, err}
	}

	return Response{StatusCreated, nil}
}
