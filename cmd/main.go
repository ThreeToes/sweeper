package main

import (
	"github.com/sirupsen/logrus"
	"sweeper/internal"
)

func main() {
	addresses, err := internal.EnumerateCidr("172.16.0.0/24")
	if err != nil {
		logrus.Fatalf("Error enumerating CIDR: %v", err)
	}
	for _, addr := range addresses {
		logrus.Infof(addr)
	}
}
