package cache

import (
	"github.com/zapj/zap/core/global"
)

func Get(k string) (interface{}, bool) {
	return global.CACHE.Get(k)
}

func GetValue(k string) (any, bool) {
	return global.CACHE.Get(k)
}

type CacheFunc func() any

func GetFromFunc[T any](k string, f CacheFunc) T {

	if v, ok := global.CACHE.Get(k); ok {
		return v.(T)
	}
	v := f()
	global.CACHE.SetDefault(k, v)
	return v.(T)
}

func GetString(k string) (string, bool) {
	if v, ok := global.CACHE.Get(k); ok {
		return v.(string), true
	}
	return "", false
}

func MustGetString(k string) string {
	if v, ok := global.CACHE.Get(k); ok {
		return v.(string)
	}
	return ""
}

func GetInt(k string) (int, bool) {
	if v, ok := global.CACHE.Get(k); ok {
		return v.(int), true
	}
	return 0, false
}

func MustGetInt(k string) int {
	if v, ok := global.CACHE.Get(k); ok {
		return v.(int)
	}
	return 0
}
