package grpc

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/kamilsk/guard/pkg/transport"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"google.golang.org/grpc"
)

var secret = "secret"

// New TODO issue#docs
func New() transport.Server {
	return &server{}
}

type server struct{}

// Serve TODO issue#docs
func (*server) Serve(listener net.Listener) error {
	defer listener.Close()
	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(middleware.TokenInjector)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(middleware.TokenInjector)),
	)
	RegisterLicenseServer(srv, NewLicenseServer())
	return srv.Serve(listener)
}
