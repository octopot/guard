package grpc

import (
	"net"

	"github.com/kamilsk/guard/pkg/transport"
	"google.golang.org/grpc"
)

// New TODO issue#docs
func New() transport.Server {
	return &server{}
}

type server struct{}

// Serve TODO issue#docs
func (*server) Serve(listener net.Listener) error {
	defer listener.Close()
	srv := grpc.NewServer()
	RegisterLicenseServer(srv, NewLicenseServer())
	return srv.Serve(listener)
}
