package zapi

func (z *Zapi) Upgrade() *Response {
	resp, err := Post("/upgrade")
	if err != nil {
		return NewResponseError(500, err.Error())
	}
	return NewResponse(resp.Status, resp.Body)
}
