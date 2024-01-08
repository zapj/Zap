package zaprpc

type errorsMapping map[int]string

func (e errorsMapping) GetString(code int, args ...any) string {
	if v, ok := e[code]; ok {
		return v
	}
	if len(args) == 1 {
		return args[0].(string)
	}
	return ""
}

var ErrorCodeMapping = errorsMapping{
	0:   "成功",
	500: "服务器内部错误",
}
