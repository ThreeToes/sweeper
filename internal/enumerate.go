package internal

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
)

func EnumerateCidr(cidr string) ([]string, error) {
	ipaddr, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	var addresses []string
	for ip := ipaddr.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		addresses = append(addresses, ip.String())
	}
	return addresses, nil
}

func inc(ip net.IP) {
	for j := len(ip)-1; j>=0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func EnumeratePorts(portSpec string) ([]int, error) {
	blocks := strings.Split(portSpec, ",")
	var ports []int
	present := map[int]bool {}
	for _, b := range blocks {
		newPorts, err := enumeratePortBlock(b)
		if err != nil {
			return nil, err
		}
		for _, p := range newPorts {
			_, ok := present[p]
			if !ok {
				ports = append(ports, p)
				present[p] = true
			}
		}
	}
	sort.Ints(ports)
	return ports, nil
}

func enumeratePortBlock(portBlock string) ([]int, error) {
	blockSplit := strings.Split(portBlock, "-")
	var ports []int
	if len(blockSplit) == 1 {
		port, err := strconv.Atoi(blockSplit[0])
		if err != nil {
			return nil, err
		}
		return []int{port}, nil
	}else if len(blockSplit) == 2 {
		var start, end int
		start, err := strconv.Atoi(blockSplit[0])
		if err != nil {
			return nil, err
		}
		end, err = strconv.Atoi(blockSplit[1])
		if err != nil {
			return nil, err
		}
		for i := start; i <= end; i++ {
			ports = append(ports, i)
		}
	} else {
		return nil, fmt.Errorf("invalid port spec %s", portBlock)
	}
	return ports, nil
}
