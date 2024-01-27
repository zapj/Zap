package tests

import (
	"strings"
	"testing"

	"github.com/zapj/zap/core/workflows"
	"gopkg.in/yaml.v3"
)

var BUILD_YAML = `
# nginx 安装包
name : Nginx
title : Nginx
version : 1.0.0
icon: "icon.png"
author: ZAP
OrganizationName: ZAP Inc
actions: 
  install: 安装
  build: 编译安装
steps:
  - name : "初始化环境"
    run : | 
          time -a "asdfasdf"
          date
          ls -la
  - name : "Build"
    run : | 
          time
          date
          ls -la
  - name: "v1.0.2"
    file: v1-0-2.yaml
          

`

func Test_YAMLPackage(t *testing.T) {

	out := &workflows.Workflows{}
	yaml.Unmarshal([]byte(BUILD_YAML), out)
	t.Log(out.Steps)
	// t.Logf("%v", out)
}

func Test_YAMLREAD(t *testing.T) {

	rs := make(map[any]any)
	err := yaml.Unmarshal([]byte(BUILD_YAML), &rs)
	t.Log(err)
	runs := rs["steps"].([]any)
	for _, v := range runs {
		for kk, val := range v.(map[string]any) {
			t.Logf("%s = %v", kk, val)
			if kk == "run" {
				for _, vv := range strings.Split(val.(string), "\n") {
					t.Log(vv)
				}

			}
		}

	}

}
