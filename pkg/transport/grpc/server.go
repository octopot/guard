package grpc

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"google.golang.org/grpc"
)

var secret = "secret"

// New TODO issue#docs
func New(_ config.GRPCConfig) transport.Server {
	return &server{}
}

type server struct{}

// Serve TODO issue#docs
func (*server) Serve(listener net.Listener) error {
	defer listener.Close()
	srv := grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_auth.StreamServerInterceptor(middleware.TokenInjector),
			grpc_prometheus.StreamServerInterceptor,
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_auth.UnaryServerInterceptor(middleware.TokenInjector),
			grpc_prometheus.UnaryServerInterceptor,
		),
	)
	RegisterLicenseServer(srv, NewLicenseServer())
	return srv.Serve(listener)
}

// Gateway TODO issue#docs
func Gateway(cnf config.GRPCConfig) transport.Server {
	return &gateway{config: cnf}
}

type gateway struct {
	config config.GRPCConfig
}

// Serve TODO issue#docs
func (gw *gateway) Serve(listener net.Listener) error {
	defer listener.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	conn, err := grpc.DialContext(ctx, gw.config.Interface, grpc.WithInsecure())
	if err != nil {
		return err
	}
	if err = RegisterLicenseHandler(ctx, mux, conn); err != nil {
		return err
	}
	return http.Serve(listener, mux)
}
