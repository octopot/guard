package http

import (
	"net"
	"net/http"

	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport"
	"github.com/kamilsk/guard/pkg/transport/http/router/chi"
)

// New TODO issue#docs
func New(cnf config.ServerConfig, service Service) transport.Server {
	return &server{config: cnf, service: service}
}

type server struct {
	config  config.ServerConfig
	service Service
}

// Serve TODO issue#docs
func (srv *server) Serve(listener net.Listener) error {
	defer listener.Close()
	return (&http.Server{Addr: srv.config.Interface, Handler: chi.Configure(srv),
		ReadTimeout:       srv.config.ReadTimeout,
		ReadHeaderTimeout: srv.config.ReadHeaderTimeout,
		WriteTimeout:      srv.config.WriteTimeout,
		IdleTimeout:       srv.config.IdleTimeout,
	}).Serve(listener)
}
