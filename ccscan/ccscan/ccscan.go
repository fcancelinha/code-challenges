package ccscan

import "net"

type Connection struct {
	host string
	port int
}

func (c *Connection) ScanPort(adress string) error {
	conn, err := net.Dial("tcp", adress)
	if err != nil {
		return err
	}

	conn.Close()

	return nil
}
