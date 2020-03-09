package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (address IPAddr) String() string {
	res := make([]string, len(address))
	for i, octet := range address {
		res[i] = strconv.Itoa(int(octet))
	}
	return strings.Join(res, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
