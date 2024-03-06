package zapi

import "github.com/go-zoox/fetch"

func (z *Zapi) UninstallApp(body map[string]any) *Response {
	resp, err := Post("/app/uninstall", &fetch.Config{Body: body})
	if err != nil {
		return NewResponseError(500, err.Error())
	}
	return NewResponse(resp.Status, resp.Body)
}
