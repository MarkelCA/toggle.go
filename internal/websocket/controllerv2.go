package websocket

import (
	"fmt"
	"log/slog"

	"github.com/markelca/toggles/pkg/flags"
)

func (ws WSController) RunAction(action *Action) Response {
	var response Response
	switch action.Action {
	case Get:
		response = ws.GetV2(action)
	case Update:
	case Create:
	case Delete:
	default:
		msg := fmt.Sprintf("Invalid action (%v)", action.Action)
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

func (ws WSController) GetV2(action *Action) Response {
	if action.Flag == nil {
		flags, err := ws.FlagService.List()
		if err != nil {
			return Response{StatusInternalServerError, err}
		}
		return Response{StatusSuccess, flags}
	}
	flag, err := ws.FlagService.Get(*action.Flag)
	if err != nil {
		if err == flags.ErrFlagNotFound {
			return Response{StatusNotFound, err}
		}
		return Response{StatusInternalServerError, err}
	}
	return Response{StatusSuccess, flag}
}
