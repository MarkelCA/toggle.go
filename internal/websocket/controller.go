package websocket

import (
	"fmt"
	"log/slog"

	"github.com/markelca/toggles/pkg/flags"
	"github.com/markelca/toggles/pkg/storage"
)

type WSController struct {
	FlagService flags.FlagService
	CacheClient storage.CacheClient
}

func (ws WSController) Update(cmd *Command) Response {
	flag, err := flags.ParseFlag(cmd.Data)
	if err != nil {
		return Response{StatusInternalServerError, err}
	}
	err = ws.FlagService.Update(flag.Name, flag.Value)
	if err != nil {
		if err == flags.ErrFlagNotFound {
			return Response{StatusNotFound, err}
		}
		return Response{StatusInternalServerError, err}
	}

	return Response{StatusCreated, nil}
}

func (ws WSController) RunCommand(cmd *Command) Response {
	var response Response
	switch cmd.Command {
	case CommandTypeGet:
		response = ws.Get(cmd)
	case CommandTypeCreate:
		response = ws.Create(cmd)
	case CommandTypeUpdate:
		response = ws.Update(cmd)
	case CommandTypeDelete:
		response = ws.Delete(cmd)
	default:
		msg := fmt.Sprintf("Invalid command (%v)", cmd.Command)
		response = Response{StatusBadRequest, msg}
	}

	switch response.Status {
	case StatusInternalServerError:
		slog.Error(fmt.Sprint(response.Value))
	default:
		slog.Info(fmt.Sprint(cmd))
	}
	return response
}

func (ws WSController) Get(c *Command) Response {
	if c.Data == nil {
		flags, err := ws.FlagService.List()
		if err != nil {
			return Response{StatusInternalServerError, err}
		}
		return Response{StatusSuccess, flags}
	}
	key := c.Data.(string)
	value, err := ws.FlagService.Get(key)
	if err != nil {
		if err == flags.ErrFlagNotFound {
			return Response{StatusNotFound, err}
		}
		return Response{StatusInternalServerError, err}
	}
	return Response{StatusSuccess, value}
}

func (ws WSController) Create(cmd *Command) Response {
	flag, err := flags.ParseFlag(cmd.Data)
	if err != nil {
		return Response{StatusInternalServerError, err}
	}
	err = ws.FlagService.Create(*flag)
	if err != nil {
		if err == flags.ErrFlagAlreadyExists {
			return Response{StatusConflict, err}
		}
		return Response{StatusInternalServerError, err}
	}
	return Response{StatusCreated, flag}
}

func (ws WSController) Delete(cmd *Command) Response {
	key := fmt.Sprintf("%v", cmd.Data)
	err := ws.FlagService.Delete(key)
	if err != nil {
		if err == flags.ErrFlagNotFound {
			return Response{StatusNotFound, nil}
		}
		return Response{StatusInternalServerError, nil}
	}
	return Response{StatusSuccess, nil}
}
