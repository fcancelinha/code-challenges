package ccscan_test

import (
	"net"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fcancelinha/code-challenges/ccscan/ccscan"
)

func TestPortScanning(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:6553")
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewUnstartedServer(nil)
	ts.Listener = l

	ts.Start()
	defer ts.Close()

	c := ccscan.Connection{
		Host:             "localhost",
		Port:             6553,
		Network:          "tcp",
		Timeout:          100 * time.Millisecond,
		ConcurrencyLevel: 100,
	}

	opnprt := c.ScanPorts()

	if len(opnprt) != 1 {
		t.Error("Port was unsuccessfully scanned")
	}

	if opnprt[0] != c.Port {
		t.Errorf("Wanted port: 6553, got: %d ", opnprt[0])
	}
}
