package websocket

func (ws WSController) RunAction(action Action) {
	switch action.Action {
	case Get:
	case Update:
	case Create:
	case Delete:
	default:

	}
}

func (ws WSController) GetV2(action *Action) Response {
	flags, _ := ws.FlagService.List()
	return Response{Status: StatusSuccess, Value: flags}
}
