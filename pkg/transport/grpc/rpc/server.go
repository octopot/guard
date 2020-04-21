package rpc

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"

	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"github.com/kamilsk/guard/pkg/transport/grpc/protobuf"
)

// New TODO issue#docs
func New(_ config.GRPCConfig, service Maintenance, storage Storage) transport.Server {
	return &server{service, storage}
}

type server struct {
	service Maintenance
	storage Storage
}

// Serve TODO issue#docs
func (server *server) Serve(listener net.Listener) error {
	defer listener.Close()
	srv := grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_prometheus.StreamServerInterceptor,
			grpc_auth.StreamServerInterceptor(middleware.TokenInjector),
			grpc_recovery.StreamServerInterceptor(),
		),
		grpc_middleware.WithUnaryServerChain(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_auth.UnaryServerInterceptor(middleware.TokenInjector),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	protobuf.RegisterLicenseServer(srv, &licenseServer{server.storage})
	protobuf.RegisterMaintenanceServer(srv, &maintenanceServer{server.service})
	return srv.Serve(listener)
}
