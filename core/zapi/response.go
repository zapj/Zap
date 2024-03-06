package zapi

import (
	"encoding/json"

	"github.com/zapj/zap/core/base"
	"github.com/zapj/zap/core/utils/jsonutil"
	"github.com/zapj/zap/core/utils/zaputil"
)

type Response struct {
	Status int    //http status
	Data   []byte // json bytes
	Code   int
	ZMap   base.ZapMap
}

func NewResponse(status int, data []byte) *Response {
	zmap := jsonutil.DecodeToZapMap(string(data))
	code := 0
	if _, ok := zmap["code"]; !ok {
		code = -1
	} else {
		code = zaputil.MustConvertStringToInt(zmap.GetString("code"))
	}
	return &Response{
		Status: status,
		Data:   data,
		Code:   code,
		ZMap:   zmap,
	}
}

func NewResponseError(status int, msg string) *Response {
	return &Response{
		Status: status,
		Data: jsonutil.Encode(map[string]any{
			"code": status,
			"msg":  msg,
		}),
	}
}

func (r *Response) String() string {
	return string(r.Data)
}

func (r *Response) JSON() map[string]any {
	v := make(map[string]any)
	json.Unmarshal(r.Data, &v)
	return v
}

func (r *Response) ZapMap() base.ZapMap {
	v := make(base.ZapMap)
	json.Unmarshal(r.Data, &v)
	return v
}

func (r *Response) HasError() bool {
	return r.Status != 200
}

func (r *Response) ErrorMessage() string {
	if r.Status != 200 {
		return string(r.Data)
	} else if r.ZMap.GetString("code") != "0" {
		return r.ZMap.GetString("msg")
	}
	return ""
}
