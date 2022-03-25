package api

func Err(msg string, status int) *HTTPErr {
	return &HTTPErr{
		Msg:    msg,
		Status: status,
	}
}
