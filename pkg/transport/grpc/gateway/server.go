package gateway

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport"
	"google.golang.org/grpc"
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
	conn, err := grpc.DialContext(ctx, server.config.Interface, grpc.WithInsecure())
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
