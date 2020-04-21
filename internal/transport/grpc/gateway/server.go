package gateway

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"go.octolab.org/ecosystem/guard/internal/config"
	"go.octolab.org/ecosystem/guard/internal/transport"
)

// New TODO issue#docs
func New(cnf config.GRPCConfig) transport.Server {
	return &server{config: cnf}
}

type server struct {
	config config.GRPCConfig
}

// Serve TODO issue#docs
func (server *server) Serve(listener net.Listener) error {
	defer listener.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	conn, err := grpc.DialContext(ctx, server.config.RPC.Interface, grpc.WithInsecure())
	if err != nil {
		return err
	}
	if err = RegisterLicenseHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err = RegisterMaintenanceHandler(ctx, mux, conn); err != nil {
		return err
	}
	return (&http.Server{Handler: mux,
		ReadTimeout:       server.config.Gateway.ReadTimeout,
		ReadHeaderTimeout: server.config.Gateway.ReadHeaderTimeout,
		WriteTimeout:      server.config.Gateway.WriteTimeout,
		IdleTimeout:       server.config.Gateway.IdleTimeout,
	}).Serve(listener)
}
