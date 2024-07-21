package ccscan_test

import (
	"testing"

	"github.com/fcancelinha/code-challenges/ccscan/ccscan"
	"gotest.tools/v3/assert"
)

const address = "scanme.nmap.org"

func TestPortScanning(t *testing.T) {
	err := ccscan.ScanPort(address + ":80")
	assert.NilError(t, err)
}
