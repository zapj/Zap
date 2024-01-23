package time_utils

import "time"

var DATE_TIME_FORMAT = "2006-01-02 15:04:05"
var DATE_FORMAT = "2006-01-02 15:04:05"
var TIME_FORMAT = "15:04:05"

func FormatUnixToDateTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DATE_TIME_FORMAT)
}
func FormatUnixMilliToDateTime(timestamp int64) string {
	return time.UnixMilli(timestamp).Format(DATE_TIME_FORMAT)
}

func FormatUnixMicroToDateTime(timestamp int64) string {
	return time.UnixMicro(timestamp).Format(DATE_TIME_FORMAT)
}

func FormatUnixToDate(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(DATE_FORMAT)
}

func FormatUnixToTime(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TIME_FORMAT)
}
