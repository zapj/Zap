package zaputil

import "strings"

func GetMajorMinorVersion(version string) string {
	vers := strings.Split(version, ".")
	marjor, minor := "0", "0"

	if len(vers) > 0 {
		marjor = vers[0]
	}
	if len(vers) > 1 {
		minor = vers[1]
	}
	return marjor + "." + minor
}

func GetMajorVersion(version string) string {
	vers := strings.Split(version, ".")
	marjor := "0"

	if len(vers) > 0 {
		marjor = vers[0]
	}

	return marjor
}
