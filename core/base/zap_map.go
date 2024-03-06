package base

import "strconv"

type ZapMap map[string]any

func (z ZapMap) Set(key string, value any) {
	z[key] = value
}

func (z ZapMap) GetAny(key string) any {
	return z[key]
}

func (z ZapMap) Get(key string) (any, bool) {
	v, ok := z[key]
	return v, ok
}

func (z ZapMap) GetString(key string) string {
	if _, ok := z[key]; !ok {
		return ""
	}
	switch z[key].(type) {
	case string:
		return z[key].(string)
	case int:
		return strconv.Itoa(z[key].(int))
	case int64:
		return strconv.FormatInt(z[key].(int64), 10)
	case float64:
		return strconv.FormatFloat(z[key].(float64), 'f', -1, 64)
	case nil:
		return ""
	}
	return z[key].(string)
}

func (z ZapMap) GetInt(key string) int {
	if _, ok := z[key]; !ok {
		return 0
	}
	switch z[key].(type) {
	case string:
		v, err := strconv.Atoi(z[key].(string))
		if err != nil {
			return 0
		}
		return v
	case int:
		return z[key].(int)
	case int64:
		return int(z[key].(int64))
	case float64:
		return int(z[key].(float64))
	}
	return 0
}

func (z ZapMap) GetInt64(key string) int64 {
	if _, ok := z[key]; !ok {
		return 0
	}
	switch z[key].(type) {
	case string:
		v, err := strconv.ParseInt(z[key].(string), 10, 64)
		if err != nil {
			return 0
		}
		return v
	case int:
		return int64(z[key].(int))
	case int64:
		return z[key].(int64)
	case float64:
	}
	return z[key].(int64)
}

func (z ZapMap) GetFloat64(key string) float64 {
	if _, ok := z[key]; !ok {
		return 0
	}
	switch z[key].(type) {
	case string:
		v, err := strconv.ParseFloat(z[key].(string), 64)
		if err != nil {
			return 0
		}
		return v
	case int:
		return float64(z[key].(int))
	case int64:
		return float64(z[key].(int64))
	case float64:
	}
	return z[key].(float64)
}

func (z ZapMap) GetBool(key string) bool {
	if _, ok := z[key]; !ok {
		return false
	}
	return z[key].(bool)
}

func (z ZapMap) GetZapMap(key string) ZapMap {
	if _, ok := z[key]; !ok {
		return make(ZapMap)
	}
	switch zmap := z[key].(type) {
	case ZapMap:
		return zmap
	default:
		return make(ZapMap)
	}
}

func (z ZapMap) Len() int {
	return len(z)
}
