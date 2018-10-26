package rpc

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"github.com/kamilsk/guard/pkg/transport/grpc/protobuf"
	"google.golang.org/grpc"
)

// New TODO issue#docs
func New(_ config.GRPCConfig, service Maintenance, storage ProtectedStorage) transport.Server {
	return &server{service, storage}
}

type server struct {
	service Maintenance
	storage ProtectedStorage
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
