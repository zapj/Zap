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

func AppVersionToArray(version string) []string {
	appVers := strings.Split(version, ".")

	if len(appVers) < 1 {
		appVers = []string{"0", "0", "0"}
	}
	if len(appVers) < 2 {
		appVers = append(appVers, "0")
	}
	if len(appVers) < 3 {
		appVers = append(appVers, "0", "0")
	}
	return appVers
}

func GetMajorVersion(version string) string {
	vers := strings.Split(version, ".")
	marjor := "0"

	if len(vers) > 0 {
		marjor = vers[0]
	}

	return marjor
}
