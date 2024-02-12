package defines

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
	return z[key].(string)
}

func (z ZapMap) GetInt(key string) int {
	return z[key].(int)
}

func (z ZapMap) GetInt64(key string) int64 {
	return z[key].(int64)
}

func (z ZapMap) GetFloat64(key string) float64 {
	return z[key].(float64)
}

func (z ZapMap) GetBool(key string) bool {
	return z[key].(bool)
}

func (z ZapMap) GetZapMap(key string) ZapMap {
	return z[key].(ZapMap)
}

func (z ZapMap) Len() int {
	return len(z)
}
