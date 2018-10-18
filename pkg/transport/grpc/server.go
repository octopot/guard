package grpc

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"google.golang.org/grpc"
)

// New TODO issue#docs
func New(_ config.GRPCConfig, service Maintenance, storage ProtectedStorage, draft DraftStorage) transport.Server {
	return &server{service, storage, draft}
}

type server struct {
	service Maintenance
	storage ProtectedStorage
	draft   DraftStorage
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
	RegisterLicenseServer(srv, NewLicenseServer(server.storage, server.draft))
	RegisterMaintenanceServer(srv, NewMaintenanceServer(server.service))
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
func (gateway *gateway) Serve(listener net.Listener) error {
	defer listener.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	conn, err := grpc.DialContext(ctx, gateway.config.Interface, grpc.WithInsecure())
	if err != nil {
		return err
	}
	if err = RegisterLicenseHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err = RegisterMaintenanceHandler(ctx, mux, conn); err != nil {
		return err
	}
	// TODO issue#configure
	return (&http.Server{Handler: mux,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
	}).Serve(listener)
}
