package protect

import (
	"regexp"
	"strings"
)

func IsBinary(mimeType string) bool {
	switch mimeType {
	case "application/octet-stream":
		return true
	}
	return false
}

func IsImage(mimeType string) bool {
	return strings.HasPrefix(mimeType, "image/")
}

func IsTextFile(mimeType string) bool {
	// var validText = regexp.MustCompile(`\w+/(txt|xml)$`)

	// fmt.Println(validID.MatchString(""))

	if strings.HasPrefix(mimeType, "text/") {
		return true
	} else if validText := regexp.MustCompile(`\w+/.*(txt|xml|sh|json|javascript).*$`); validText.MatchString(mimeType) {
		return true
	}

	return false
}
