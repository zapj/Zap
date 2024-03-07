package nginx_test

import (
	"testing"
)

func TestNginxConfig(t *testing.T) {
	t.Log("TestNginxConfig")
	// ngxSbin := path.Join("/usr/local/apps/nginx/sbin/nginx")
	// ngxConf := path.Join("/usr/local/apps/nginx/conf/nginx.conf")
	// output, err := cmdutil.ExecShCmd(fmt.Sprintf("%s -t -c %s", ngxSbin, ngxConf))
	// t.Log(output, err)
	// if bytes.Contains(output, []byte("nginx.conf test is successful")) {
	// 	t.Log("nginx.conf test is successful")
	// } else {
	// 	t.Error("nginx.conf test is failed")
	// }
	// var re = regexp.MustCompile(`nginx:\s\[(?P<errlevel>.*?)\]\s(?P<errmsg>.*?)in\s(?P<errfile>\/.*)\:\d+`)
	// s := re.FindAllStringSubmatch(string(output), -1)
	// for _, v := range s {
	// 	t.Logf("%s", v[0])
	// 	t.Logf("%s", v[re.SubexpIndex("errlevel")])
	// 	t.Logf("%s", v[re.SubexpIndex("errmsg")])
	// 	t.Logf("%s", v[re.SubexpIndex("errfile")])
	// }

}
