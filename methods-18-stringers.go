/*
https://tour.golang.org/methods/18
*/
package main

import "fmt"

type IPAddr [4]byte

func (this IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", this[0], this[1], this[2], this[3])
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
