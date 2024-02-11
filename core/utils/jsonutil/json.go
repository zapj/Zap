package jsonutil

import "encoding/json"

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
