package datautils

import "strconv"

func FmtPercentFloat64(percent float64) string {
	return strconv.FormatFloat(percent, 'f', 2, 64)
}

func FmtPercentFloat32(percent float32) string {
	return strconv.FormatFloat(float64(percent), 'f', 2, 32)
}
