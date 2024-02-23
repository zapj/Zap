package zaputil

import "strings"

func ParseKeyPair(query string) map[string]any {
	keyPairs := make(map[string]any)
	for _, pair := range strings.Split(query, "&") {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			keyPairs[kv[0]] = kv[1]
		}
	}
	return keyPairs
}
