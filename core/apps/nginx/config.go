package nginx

import (
	"bytes"
	"path"
	"regexp"

	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/cmdutil"
)

type ConfigError struct {
	ErrMsg   string
	ErrLevel string
	ErrFile  string
	ErrNum   string
}

func nginxConfigTest() *ConfigError {
	//default config file
	ngxSbin := path.Join(global.APPS_DIR, "nginx/sbin/nginx")
	ngxConf := path.Join(global.APPS_DIR, "nginx/conf/nginx.conf")
	output, err := cmdutil.ExecCmd(ngxSbin, "-t", "-c", ngxConf)
	if err == nil {
		return nil
	}

	if bytes.Contains(output, []byte("nginx.conf test is successful")) {
		return nil
	}

	var re = regexp.MustCompile(`nginx:\s\[(?P<errlevel>.*?)\]\s(?P<errmsg>.*?)in\s(?P<errfile>\/.*)\:(?P<errnum>\d+)`)
	result := re.FindAllStringSubmatch(string(output), -1)
	if len(result) > 0 {
		match := result[0]
		if re.SubexpIndex("errfile") == -1 || re.SubexpIndex("errlevel") == -1 || re.SubexpIndex("errmsg") == -1 {
			return nil
		}
		return &ConfigError{
			ErrMsg:   match[re.SubexpIndex("errmsg")],
			ErrLevel: match[re.SubexpIndex("errlevel")],
			ErrFile:  match[re.SubexpIndex("errfile")],
			ErrNum:   match[re.SubexpIndex("errnum")],
		}
	}
	return nil
}
