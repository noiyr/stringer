package main

import (
	"fmt"
	"strings"
	_"strconv"
)

type IPAddr [4]byte
type IPv6 [6]byte

func (ip IPAddr) String() string {
	s := make([]string, len(ip))
	
	
	return strings.Join(s, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}


	sex := map[string]IPv6{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range sex {
		fmt.Printf("%v: %v\n", name, ip)
	}
}