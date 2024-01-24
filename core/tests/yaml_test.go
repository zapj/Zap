package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/zapj/zap/core/workflows"
	"gopkg.in/yaml.v3"
)

func Test_YAMLPackage(t *testing.T) {
	ng := "/home/zap/works/go/Zap/data/appstore/zap/build.yaml"
	conf, _ := os.ReadFile(ng)
	out := &workflows.Workflows{}
	yaml.Unmarshal(conf, out)
	t.Log(out.Steps)
	// t.Logf("%v", out)
}

func Test_YAMLREAD(t *testing.T) {
	ng := "/home/zap/works/go/Zap/data/appstore/zap/build.yaml"
	conf, _ := os.ReadFile(ng)
	rs := make(map[any]any)
	err := yaml.Unmarshal(conf, &rs)
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
