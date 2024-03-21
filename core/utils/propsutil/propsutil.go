package propsutil

import (
	"bytes"
	"io"
	"os"

	"github.com/magiconair/properties"
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

func NewPropertiesFromString(v string) *properties.Properties {
	p := properties.NewProperties()
	p.WriteSeparator = "="
	p.Load([]byte(v), properties.UTF8)
	return p
}

func NewPropertiesFromBytes(v []byte) *properties.Properties {
	p := properties.NewProperties()
	p.WriteSeparator = "="
	p.Load(v, properties.UTF8)
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

func FmtString(v string) string {
	p := NewPropertiesFromString(v)
	var buf bytes.Buffer
	// b := bufio.NewWriter(&buf)
	p.Write(io.Writer(&buf), properties.UTF8)
	// b.Flush()
	return buf.String()
}
