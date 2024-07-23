package ccscan

import (
	"fmt"
	"math"
	"net"
	"time"
)

type PortScanner interface {
	Scan() (openPorts []int)
}

type Connection struct {
	Host             string
	Network          string
	Port             uint16
	Timeout          time.Duration
	ConcurrencyLevel int
}

// TODO: find way to remove the results dependency, it's inefficient to send a result for each port
// especially when it is an error or it's closed
func (c *Connection) scan(ports, results chan uint16) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", c.Host, p)
		conn, err := net.DialTimeout(c.Network, address, c.Timeout)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func (c *Connection) ScanPorts() (openPorts []uint16) {
	start, end := uint16(1), uint16(math.MaxUint16)
	ports := make(chan uint16, c.ConcurrencyLevel)
	results := make(chan uint16)
	total := (end - start)

	if c.Port > 0 {
		start, end = c.Port, c.Port
		total = 1
	}

	for i := 0; i < cap(ports); i++ {
		go c.scan(ports, results)
	}

	go func() {
		for i := start; i <= end; i++ {
			ports <- i
		}
	}()

	for i := uint16(1); i <= total; i++ {
		p := <-results
		progressBar(i, total)
		if p != 0 {
			openPorts = append(openPorts, p)
		}
	}

	close(ports)
	close(results)
	fmt.Println()

	return
}

func progressBar(current, total uint16) {
	bar := 100
	percentage := (current * 100) / total
	progress := int(float64(bar) * (float64(current) / float64(total)))

	fmt.Printf("\r[")
	for i := 0; i < bar; i++ {
		if i < progress {
			fmt.Print("=")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("] %d%%", percentage)
}
