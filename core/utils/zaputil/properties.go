package zaputil

import "github.com/zapj/go-properties"

func NewProperties() *properties.Properties {
	p := properties.NewProperties()
	p.WriteSeparator = "="
	return p
}
