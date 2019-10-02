package server

import "net"

type channelPool struct {
	conn net.Conn
}
