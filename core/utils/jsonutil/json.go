package jsonutil

import (
	"encoding/json"

	"github.com/zapj/zap/core/defines"
)

func EncodeToString(v any) string {
	r, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(r)
}

func Encode(v any) []byte {
	r, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return r
}

func DecodeToMap(r string) map[string]any {
	v := make(map[string]any)
	json.Unmarshal([]byte(r), &v)
	return v
}

func DecodeToZapMap(r string) defines.ZapMap {
	v := make(defines.ZapMap)
	json.Unmarshal([]byte(r), &v)
	return v
}
