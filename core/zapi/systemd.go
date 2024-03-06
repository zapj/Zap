package zapi

import (
	"github.com/go-zoox/fetch"
)

func (z *Zapi) Systemctl(action string, name string) *Response {
	resp, err := Post("/systemctl", &fetch.Config{
		Body: map[string]interface{}{
			"action": action,
			"name":   name,
		},
	})
	if err != nil {
		return NewResponseError(500, err.Error())
	}
	return NewResponse(resp.Status, resp.Body)
}
