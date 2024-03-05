package domainutil

import (
	"strings"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

func ParseDomainNames(domainnames string) []string {
	domain_names := strings.Fields(domainnames)
	var result []string
	for _, dm := range domain_names {
		if d, err := publicsuffix.Parse(dm); err == nil {
			result = append(result, d.String())
		}
	}
	return result
}

func ParseDomainName(domainname string) string {
	if d, err := publicsuffix.Parse(domainname); err == nil {
		return d.String()
	}
	return ""
}

func Domain(name string) (string, error) {
	return publicsuffix.Domain(name)
}

func Labels(name string) []string {
	return publicsuffix.Labels(name)
}

func Parse(name string) (*publicsuffix.DomainName, error) {
	return publicsuffix.Parse(name)
}

func ParseHostingDomain(domainname string, domain_names string) (string, []string) {
	domainNames := ParseDomainNames(domain_names)
	d, err := publicsuffix.Parse(domainname)

	if err != nil && len(domainNames) > 1 {
		return domainNames[0], domainNames[1:]
	}
	return d.String(), domainNames
}
