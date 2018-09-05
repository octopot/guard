package transport

import "net"

// Server TODO issue#docs
type Server interface {
	Serve(net.Listener) error
}
