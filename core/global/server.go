package global

type ServerConf struct {
	Host     string
	Port     int
	IPv6     bool
	CertFile string
	KeyFile  string
}
