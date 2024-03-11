package zapi

type WebsiteApiClient interface {
	GetWebsite() (*Response, error)
	GetWebsiteList() (*Response, error)
	GetWebsiteListByUid() (*Response, error)
	GetWebsiteListByUidAndUsername() (*Response, error)
	StopWebsite() *Response
}

func (z *Zapi) StopWebsite() *Response {
	resp, err := Post("/website/stop")
	if err != nil {
		return NewResponseError(500, err.Error())
	}
	return NewResponse(resp.Status, resp.Body)
}

type Website struct {
	Id          int    `json:"id"`
	Uid         int    `json:"uid"`
	Username    string `json:"username"`
	Domain      string `json:"domain"`
	DomainNames string `json:"domain_names"`
}
