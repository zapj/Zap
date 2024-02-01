package task

import (
	"github.com/go-zoox/fetch"
	"github.com/zapj/zap/core/utils/pathutil"
)

var defaultConfig = fetch.Config{
	UnixDomainSocket:      pathutil.GetPath("data/task"),
	TLSInsecureSkipVerify: true,
	BaseURL:               "https://127.0.0.1",
}

func Get(url string, config ...*fetch.Config) (*fetch.Response, error) {
	return fetch.New(config...).Get(url, &defaultConfig).Execute()
}

func Post(url string, config ...*fetch.Config) (*fetch.Response, error) {
	return fetch.New(config...).Post(url, &defaultConfig).Execute()
}

func Put(url string, config ...*fetch.Config) (*fetch.Response, error) {
	return fetch.New(config...).Put(url, &defaultConfig).Execute()
}

func Delete(url string, config ...*fetch.Config) (*fetch.Response, error) {
	return fetch.New(config...).Delete(url, &defaultConfig).Execute()
}
