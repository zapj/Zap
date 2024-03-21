package zaputil

import "github.com/magiconair/properties"

func NewProperties() *properties.Properties {
	p := properties.NewProperties()
	p.WriteSeparator = "="
	return p
}
