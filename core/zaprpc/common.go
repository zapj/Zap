package zaprpc

import (
	"encoding/json"
	"fmt"

	"github.com/zapj/zap/core/global"
)

type Response struct {
	Code int
	Msg  string
	Data any
}

func (r *Response) HasError() bool {
	return r.Code != 0
}

func (r *Response) ErrorCode() int {
	return r.Code
}

func (r *Response) Message() string {
	return r.Msg
}

func (r *Response) Get(args ...any) any {
	if r.Data == nil && len(args) == 1 {
		return args[0]
	}
	return r.Data
}

func (r *Response) GetString(args ...any) string {
	if r.Code == 0 && r.Data != nil {
		return fmt.Sprint(r.Data)
	}
	if len(args) == 1 {
		return fmt.Sprint(args[0])
	}
	return ""
}

func (r *Response) GetInt(args ...any) int {
	if r.Code == 0 && r.Data != nil {
		return r.Data.(int)
	}
	if len(args) == 1 {
		return args[0].(int)
	}
	return 0
}

func (r *Response) GetInt64(args ...any) int64 {
	if r.Code == 0 && r.Data != nil {
		return r.Data.(int64)
	}
	if len(args) == 1 {
		return args[0].(int64)
	}
	return 0
}

func (r *Response) GetBool(args ...any) bool {
	if r.Code == 0 && r.Data != nil {
		return r.Data.(bool)
	}
	if len(args) == 1 {
		return args[0].(bool)
	}
	return false
}

func (r *Response) GetJSON(v any, args ...any) error {
	if r.Code == 0 && r.Data != nil {
		return json.Unmarshal([]byte(r.Data.(string)), &v)

	}
	if len(args) == 1 {
		v = args[0]
		return nil
	}
	return fmt.Errorf("parsing JSON format failed")
}

func (r *Response) SetData(v any) {
	r.Data = v
}

func (r *Response) SetDataJson(v any) {
	var err error
	data, err := json.Marshal(v)
	if err != nil {
		r.Data = nil
	}
	r.Data = string(data)
}

type Args map[string]any

func (a Args) Reset() {
	for k := range a {
		delete(a, k)
	}
}

func (a Args) Get(key string, args ...any) any {
	if v, ok := a[key]; ok {
		return v
	}
	if len(args) == 1 {
		return args[0]
	}
	return nil
}

func (a Args) GetString(key string, args ...any) string {
	if v, ok := a[key]; ok {
		return v.(string)
	}
	if len(args) == 1 {
		return args[0].(string)
	}
	return ""
}

func (a Args) GetInt(key string, args ...any) int {
	if v, ok := a[key]; ok {
		return v.(int)
	}
	if len(args) == 1 {
		return args[0].(int)
	}
	return 0
}

func (a Args) GetInt64(key string, args ...any) int64 {
	if v, ok := a[key]; ok {
		return v.(int64)
	}
	if len(args) == 1 {
		return args[0].(int64)
	}
	return 0
}

func (a Args) GetBool(key string, args ...any) bool {
	if v, ok := a[key]; ok {
		return v.(bool)
	}
	if len(args) == 1 {
		return args[0].(bool)
	}
	return false
}

func (a Args) GetStringArray(key string, args ...string) []string {
	if v, ok := a[key]; ok {
		global.LOG.Debugf("vvv %v", v)
		switch val := v.(type) {
		case []string:
			return val
		default:
			return nil
		}
	}
	if len(args) == 1 {
		return args
	}
	return nil
}
