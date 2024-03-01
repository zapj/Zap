package ngx

import (
	"bytes"
	"fmt"
	"strings"
)

type NgxConfLocation struct {
	Root    string `json:"root"`
	Expires string `json:"expires"`

	ProxyPass                string            `json:"proxy_pass"`
	ProxyRedirect            string            `json:"proxy_redirect"`
	ProxySetHeader           map[string]string `json:"proxy_set_header"`
	ProxyTempPath            string            `json:"proxy_temp_path"`
	ProxyConnectTimeout      string            `json:"proxy_connect_timeout"`
	ProxySendTimeout         string            `json:"proxy_send_timeout"`
	ProxyReadTimeout         string
	ProxyBufferSize          string
	ProxyBusyBuffersSize     string
	Charset                  string
	ClientMaxBodySize        string
	ClientBodyBufferSize     string
	ClientBodyTemp           string
	ClientBodyTempPath       string
	Rewrite                  []string `json:"rewrite"`
	Include                  string
	FastcgiPass              string `json:"fastcgi_pass"`
	FastcgiIndex             string `json:"fastcgi_index"`
	FastcgiParam             string `json:"fastcgi_param"`
	FastcgiIgnoreClientAbort string

	ValidReferers string

	Alias string
	Deny  []string
	Allow []string
}

type NgxConfServer struct {
	Listen                 []string `json:"listen"`
	ServerName             []string
	RunDirectory           string `json:"run_directory"` // public
	Root                   string `json:"root"`
	AccessLog              string `json:"access_log"`
	ErrorLog               string `json:"error_log"`
	Index                  string `json:"index"`
	Charset                string `json:"charset"`
	Locations              []NgxConfLocation
	Include                []string `json:"include"`
	Ssl                    string   `json:"ssl"`
	SslCertificate         string   `json:"ssl_certificate"`
	SslCertificateKey      string   `json:"ssl_certificate_key"`
	SslSessionTimeout      string   `json:"ssl_session_timeout"`
	SslProtocols           string   `json:"ssl_protocols"`
	SslCiphers             string   `json:"ssl_ciphers"`
	SslPreferServerCiphers string   `json:"ssl_prefer_server_ciphers"`
	SslStapling            string
	SslStaplingVerify      string
	SslTrustedCertificate  string
	Resolver               string

	Return string `json:"return"`

	ErrorPage string

	KeepaliveTimeout string

	Set []string

	If []string

	ForceHttps bool
}

func NewNgxConfServer(servername ...string) *NgxConfServer {
	return &NgxConfServer{
		ServerName:             servername,
		Index:                  "index.html index.htm index.php",
		Charset:                "utf-8",
		SslPreferServerCiphers: "on",
		SslProtocols:           "TLSv1 TLSv1.2 TLSv1.3",
		SslStapling:            "on",
		SslStaplingVerify:      "on",
		Resolver:               "1.1.1.1 8.8.8.8",
	}
}

func (n *NgxConfServer) AddServerName(servername ...string) {
	n.ServerName = append(n.ServerName, servername...)
}

func (n *NgxConfServer) GenerateToString() (string, error) {
	var serverConf bytes.Buffer
	var http2, ssl bool = false, false

	//start server
	serverConf.WriteString("server {\n")

	if len(n.Listen) > 0 {
		for _, listen := range n.Listen {
			if strings.Contains(listen, "http2") {
				http2 = true
			}
			if strings.Contains(listen, "ssl") {
				ssl = true
			}
			fmt.Fprintf(&serverConf, "\n\tlisten %s;", listen)
		}

	} else {
		serverConf.WriteString("\n\tlisten 80;")
	}

	//http2
	_ = http2
	//server name
	fmt.Fprintf(&serverConf, "\n\tserver_name %s;", strings.Join(n.ServerName, " "))
	//access log
	if n.AccessLog != "" {
		fmt.Fprintf(&serverConf, "\n\taccess_log %s;", n.AccessLog)
	} else {
		fmt.Fprintln(&serverConf, "\n\taccess_log off;")
	}
	if n.ErrorLog != "" {
		fmt.Fprintf(&serverConf, "\n\terror_log %s;", n.ErrorLog)
	} else {
		fmt.Fprintln(&serverConf, "\n\terror_log off;")
	}
	//root
	if n.RunDirectory != "" && n.Root != "" {
		fmt.Fprintf(&serverConf, "\n\tset $base %s;", n.Root)
		fmt.Fprintf(&serverConf, "\n\troot $base/%s;", n.RunDirectory)
	} else if n.Root != "" {
		fmt.Fprintf(&serverConf, "\n\troot %s;", n.Root)
	}
	//charset
	if n.Charset != "" {
		fmt.Fprintf(&serverConf, "\n\tcharset %s;", n.Charset)
	}
	//index
	if n.Index != "" {
		fmt.Fprintf(&serverConf, "\n\tindex %s;", n.Index)
	}

	//force to https
	if n.ForceHttps && ssl {
		fmt.Fprintln(&serverConf, "if ($scheme = http) {return 301 https://$host$request_uri;}")
	}

	// keepaliveTimeout
	if n.KeepaliveTimeout != "" {
		fmt.Fprintf(&serverConf, "\n\tkeepalive_timeout %s;", n.KeepaliveTimeout)
	}

	if ssl && n.SslCertificate != "" && n.SslCertificateKey != "" {
		fmt.Fprintf(&serverConf, "\n\tssl_certificate %s;", n.SslCertificate)
		fmt.Fprintf(&serverConf, "\n\tssl_certificate_key %s;", n.SslCertificateKey)

		if n.SslPreferServerCiphers != "" {
			fmt.Fprintf(&serverConf, "\n\tssl_prefer_server_ciphers %s;", n.SslPreferServerCiphers)
		}

		if n.SslProtocols != "" {
			fmt.Fprintf(&serverConf, "\n\tssl_protocols %s;", n.SslProtocols)
		}

		if n.SslSessionTimeout != "" {
			fmt.Fprintf(&serverConf, "\n\tssl_session_timeout %s;", n.SslSessionTimeout)
		}

		if n.SslStapling != "" {
			fmt.Fprintf(&serverConf, "\n\tssl_stapling %s;", n.SslStapling)
			fmt.Fprintf(&serverConf, "\n\tssl_stapling_verify %s;", n.SslStaplingVerify)
		}

		if n.Resolver != "" {
			fmt.Fprintf(&serverConf, "\n\tresolver %s;", n.Resolver)
		}
	}

	// end server
	serverConf.WriteString("\n}\n")
	return serverConf.String(), nil
}
