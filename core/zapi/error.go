package zapi

import "fmt"

type ZapiError struct {
	Code int
	Msg  string
}

func NewZapiError(code int, msg string) *ZapiError {
	return &ZapiError{
		Code: code,
		Msg:  msg,
	}
}

func (e *ZapiError) Error() string {
	return fmt.Sprintf("code: %d , msg: %s", e.Code, e.Msg)
}
