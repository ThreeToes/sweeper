package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"sweeper/internal"
	"sync"
)

type cidrFlags []string

func (c *cidrFlags) String() string {
	return fmt.Sprintf("%+v", *c)
}

func (c *cidrFlags) Set(value string) error {
	*c = append(*c, value)
	return nil
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		DisableLevelTruncation:	   true,
		DisableTimestamp:          true,
	})
	cidrs := &cidrFlags{}
	flag.Var(cidrs, "cidr",  "CIDR blocks to scan. Specify multiple times for multiple CIDR blocks")
	timeout := flag.Int("timeout", 100, "Timeout in milliseconds")
	workers := flag.Int("workers", 10, "Number of worker routines")
	portsF := flag.String("ports", "80,443", "Ports to scan")
	flag.Parse()
	if len(*cidrs) == 0 {
		logrus.Fatalf("Must specify at least one CIDR block with -cidr")
	}
	var addresses []string
	logrus.Debugf("Scanning %+v on ports %s with %d workers", *cidrs, *portsF, *workers)
	for _, c := range *cidrs {
		newAddrs, err := internal.EnumerateCidr(c)
		if err != nil {
			logrus.Fatalf("Error enumerating %s: %v", c, err)
		}
		addresses = append(addresses, newAddrs...)
	}
	logrus.Printf("Scanning %d IPs on ports %s", len(addresses), *portsF)
	ports, err := internal.EnumeratePorts(*portsF)
	if err != nil {
		logrus.Fatalf("Error enumerating port string %s: %v", *portsF, err)
	}
	input := make(chan *internal.DialSpec)
	results := make(chan *internal.WorkerResult)
	wg := &sync.WaitGroup{}
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go internal.Worker(wg, input, results)
	}
	var collectedResults []*internal.WorkerResult
	//collector thread
	go func() {
		for r := range results {
			collectedResults = append(collectedResults, r)
		}
	}()
	for _, addr := range addresses {
		for _, port := range ports {
			input<-&internal.DialSpec{
				Ip:      addr,
				Port:    port,
				Timeout: *timeout,
			}
		}
	}
	close(input)
	wg.Wait()
	close(results)
	for _, r := range collectedResults {
		if r.Error != nil {
			logrus.Debugf("Error scanning %s:%d: %v", r.IP, r.Port, r.Error)
		}
		if r.Up {
			logrus.Printf("%s up on port %d", r.IP, r.Port)
		}
	}
}