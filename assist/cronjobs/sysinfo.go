package main

import (
	"fmt"
	"github.com/elastic/go-sysinfo"
)

func main() {
	host, _ := sysinfo.Host()
	m, _ := host.Memory()
	fmt.Println(m.Total)

}
