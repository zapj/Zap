package base

type ServerConf struct {
	SigningKey string
	EncryptKey string
	Host       string
	Port       int
	IPv6       bool
	CertFile   string
	KeyFile    string
}
