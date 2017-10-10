package main

import (
	"fmt"

	externalip "github.com/stobbsm/go-externalip"
)

func main() {
	MyIp := externalip.New()
	if MyIp.ErrIpV4 != nil {
		fmt.Println(MyIp.ErrIpV4.Error())
	}
	fmt.Println("IPv4 address:", MyIp.IpV4)
}
