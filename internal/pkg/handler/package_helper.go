package handler

type response struct {
	Message string `json:"message"`
}

func newMsgResponse(msg string) response {
	return response{
		Message: msg,
	}
}
