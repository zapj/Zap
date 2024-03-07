package websites

type webSiteRequest struct {
	Domain       string `json:"domain" validate:"required"`
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description"`
	DomainNames  string `json:"domain_names"`
	WwwRoot      string `json:"www_root"` //只包含目录名 Userhome + wwwroot
	RunDirectory string `json:"run_directory"`
	AccessLog    string `json:"access_log"`
	ErrorLog     string `json:"error_log"`
	WebsiteId    int    `json:"website_id"` // models website id
	IndexFiles   string `json:"index_files"`
	applicationRequest
}

type applicationRequest struct {
	Application int    `json:"application"`
	ExposePort  string `json:"expose_port"`
	ExposeProto string `json:"expose_proto"`
}

func (w webSiteRequest) GetIndexFiles() string {
	if w.IndexFiles == "" {
		w.IndexFiles = "index.html index.htm"
	}
	return w.IndexFiles
}
