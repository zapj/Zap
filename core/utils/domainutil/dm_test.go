package domainutil_test

import (
	"fmt"
	"testing"

	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zapj/zap/core/utils/domainutil"
)

var domainnames = `test.com www.test.com
test1.com
test2
www.test1.com`

func TestParseDomainNames(t *testing.T) {
	domainnames := domainutil.ParseDomainNames(domainnames)
	t.Logf("ParseDomainNames(%s) = %v", domainnames, domainnames)
}

func TestDomain(t *testing.T) {
	fmt.Println(publicsuffix.Domain("example.com"))       // example.com
	fmt.Println(publicsuffix.Domain("www.example.com"))   // example.com
	fmt.Println(publicsuffix.Domain("example.co.uk"))     // example.co.uk
	fmt.Println(publicsuffix.Domain("www.example.co.uk")) // example.co.uk
	var err error
	// Parse the domain from a string
	// using the default list
	fmt.Println(publicsuffix.Parse("example.com"))       // &DomainName{"com", "example", ""}
	fmt.Println(publicsuffix.Parse("www.example.com"))   // &DomainName{"com", "example", "www"}
	fmt.Println(publicsuffix.Parse("example.co.uk"))     // &DomainName{"co.uk", "example", ""}
	fmt.Println(publicsuffix.Parse("www.example.co.uk")) // &DomainName{"co.uk", "example", "www"}
	d, _ := publicsuffix.Parse("www.web.example.com.cn")
	fmt.Println("TLD", d.TLD, "SLD", d.SLD, "TRD", d.TRD)
	d, err = publicsuffix.Parse("hostnamme")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("TLD", d.TLD, "SLD", d.SLD, "TRD", d.TRD)
	}

}
