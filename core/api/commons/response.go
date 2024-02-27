package commons

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type SuccessResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) *Response {
	return &Response{Code: 0, Msg: "OK", Data: data}
}

func SuccessMsg(msg string) *Response {
	return &Response{Code: 0, Msg: msg, Data: nil}
}

func Error(code int, msg string, data any) *Response {
	return &Response{Code: code, Msg: msg, Data: data}
}
