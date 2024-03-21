package propsutil

import (
	"os"

	"github.com/zapj/go-properties"
)

func MapToString(data map[string]any) string {
	p := properties.NewProperties()
	p.WriteSeparator = "="
	for k, v := range data {
		p.SetValue(k, v)
	}
	return p.String()
}

func NewProperties() *properties.Properties {
	p := properties.NewProperties()
	p.WriteSeparator = "="
	return p
}

func NewPropertiesFromMap(kv any) *properties.Properties {
	p := properties.NewProperties()
	p.WriteSeparator = "="
	switch kv := kv.(type) {
	case map[string]any:
		for k, v := range kv {
			p.SetValue(k, v)
		}
	}
	return p
}

func WriteMapToFile(kv any, path string) error {
	p := NewPropertiesFromMap(kv)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = p.Write(f, properties.UTF8)
	return err
}

func WritePropertiesToFile(p *properties.Properties, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = p.Write(f, properties.UTF8)
	return err
}
