package propsutil

import "github.com/zapj/go-properties"

func MapToString(data map[string]any) string {
	p := properties.NewProperties()
	p.WriteSeparator = "="
	for k, v := range data {
		p.SetValue(k, v)
	}
	return p.String()
}
