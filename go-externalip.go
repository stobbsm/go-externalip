package externalip

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type IP struct {
	IpV6    string
	ErrIpV6 error
	IpV4    string
	ErrIpV4 error
}

var localhost = &IP{IpV4: "127.0.0.1", IpV6: "::1"}

func New() *IP {
	var ip = new(IP)

	ip.setV4()
	ip.setV6()

	fmt.Println("Contents of IP:", *ip)
	return ip
}

func (i *IP) setV4() {
	var b bytes.Buffer
	i.IpV4 = "127.0.0.1"

	resp, err := http.Get("http://ipv4.myexternalip.com/raw")
	if err != nil {
		i.ErrIpV4 = err
	}
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		i.ErrIpV4 = err
	}
	resp.Body.Close()

	i.IpV4 = strings.TrimSpace(b.String())
}

func (i *IP) setV6() {
	i.IpV6 = "::1"
}
