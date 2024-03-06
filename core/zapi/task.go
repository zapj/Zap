package zapi

import "github.com/go-zoox/fetch"

func (z *Zapi) InstallApp(body map[string]any) *Response {
	resp, err := Post("/task/install/app", &fetch.Config{Body: body})
	if err != nil {
		return NewResponseError(500, err.Error())
	}
	return NewResponse(resp.Status, resp.Body)
}

func (z *Zapi) CancelTask(body map[string]any) *Response {
	resp, err := Post("/task/cancel", &fetch.Config{Body: body})
	if err != nil {
		return NewResponseError(500, err.Error())
	}
	return NewResponse(resp.Status, resp.Body)
}

func (z *Zapi) RefreshTask() *Response {
	resp, err := Post("/task/refresh")
	if err != nil {
		return NewResponseError(500, err.Error())
	}
	return NewResponse(resp.Status, resp.Body)
}
