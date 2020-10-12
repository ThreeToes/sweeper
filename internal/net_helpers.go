package internal

import (
	"net"
	"strings"
)

// GetLocalNetworks queries connected interfaces and returns
// connected networks in CIDR notation
func GetLocalNetworks() ([]string, error) {
	interfaces, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	var ret []string
	for _, i := range interfaces {
		// Stupidly huge blocks that we
		// probably don't want in the "local" addresses
		if i.String() == "127.0.0.1/8" {
			ret = append(ret, "127.0.0.1/32")
			continue
		}
		if i.String() == "::1/64" {
			ret = append(ret, "::1/128")
		}
		ret = append(ret, i.String())
	}

	return ret, nil
}

// Ipv4Only takes a list of CIDRs and returns only IPv4 addresses
func Ipv4Only(addresses []string) []string {
	var ret []string
	for _, addr := range addresses {
		split := strings.Split(addr, "/")
		_, err := net.ResolveIPAddr("ip4", split[0])
		if err != nil {
			continue
		}
		ret = append(ret, addr)
	}
	return ret
}

// Ipv6Only takes a list of CIDRs and returns only IPv4 addresses
func Ipv6Only(addresses []string) []string {
	var ret []string
	for _, addr := range addresses {
		split := strings.Split(addr, "/")
		_, err := net.ResolveIPAddr("ip6", split[0])
		if err != nil {
			continue
		}
		ret = append(ret, addr)
	}
	return ret
}