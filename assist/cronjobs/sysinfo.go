package main

import (
	"fmt"

	"github.com/elastic/go-sysinfo"
)

func main() {
	host, _ := sysinfo.Host()
	fmt.Println(host.Info().OS.Platform)
	m, _ := host.Memory()
	fmt.Println(m.Total)

}
