package main

import (
	"strconv"
	"fmt"
)

type IPAddr [4]byte

type Stringer interface {
    String() string
}
func (r IPAddr)String() string {
	var myString string
	for i,x := range r {
		t:= strconv.Itoa(int(x))
		//fmt.Println(x,t)
		myString += string(t)
		if i<3 {
			myString +="."
		}
	}
	return myString
}

func main() {
	
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		var s Stringer
		s=ip
		s.String()
		fmt.Printf("%v: %v\n", name, ip)
	}
}
