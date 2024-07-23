package main

import (
	"flag"
	"fmt"
	"sort"
	"time"

	"github.com/fcancelinha/code-challenges/ccscan/ccscan"
)

var (
	host    = flag.String("host", "", "the host/address for port scanning")
	port    = flag.Uint("port", 0, "the port for scannning")
	cc      = flag.Int("concurrency", 100, "the number of parallel workers scanning")
	timeout = flag.Float64("timout", 700, "the timout for each scan in milliseconds")
	help    = flag.Bool("help", false, "CLI usage")
)

func main() {
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	printInfo()

	c := ccscan.Connection{
		Host:             *host,
		Port:             uint16(*port),
		Network:          "tcp",
		Timeout:          700 * time.Millisecond,
		ConcurrencyLevel: 1000,
	}

	oprts := c.ScanPorts()

	if len(oprts) == 0 {
		fmt.Println("No open port(s)")
	}

	sort.Slice(oprts, func(i, j int) bool { return oprts[i] < oprts[j] })

	for _, p := range oprts {
		fmt.Printf("Port: %d is open\n", p)
	}
}

func printInfo() {
	fmt.Print("Scanning ")
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "host" || f.Name == "port" {
			fmt.Printf(" %s: %v ", f.Name, f.Value)
		}
	})
	fmt.Println()
}
