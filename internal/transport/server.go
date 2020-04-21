package transport

import "net"

// Server TODO issue#docs
type Server interface {
	// Serve TODO issue#docs
	Serve(net.Listener) error
}
