package domainutil

import (
	"strings"
)

func ParseDomainNames(domainnames string) []string {
	domain_names := strings.Fields(domainnames)
	return domain_names
}
