package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]int

func (ip IPAddr) String() string {
	var result string
	for _, octet := range ip {
		result += strconv.Itoa(octet) + "."
	}
	return strings.TrimSuffix(result, ".")
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
