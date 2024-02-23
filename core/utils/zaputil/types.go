package zaputil

import "strconv"

func MustConvertStringToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return r
}

func MustConvertStringToUint(s string) uint {
	r, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return uint(r)
}
