package websocket

import (
	"fmt"
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
	switch cmd.Command {
	case CommandTypeGet:
		return ws.Get(cmd)
	case CommandTypeCreate:
		return ws.Create(cmd)
	case CommandTypeUpdate:
		return ws.Update(cmd)
	case CommandTypeDelete:
		return ws.Delete(cmd)
	default:
		msg := fmt.Sprintf("Invalid command (%v)", cmd.Command)
		return Response{StatusBadRequest, msg}
	}
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
