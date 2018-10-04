package cmd

import (
	"log"
	"net"
	"runtime"
	"strconv"

	"github.com/kamilsk/go-kit/pkg/fn"
	"github.com/kamilsk/go-kit/pkg/strings"
	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/service/guard"
	"github.com/kamilsk/guard/pkg/storage"
	"github.com/kamilsk/guard/pkg/transport/grpc"
	"github.com/kamilsk/guard/pkg/transport/http"
	"github.com/kamilsk/guard/pkg/transport/http/monitor"
	"github.com/kamilsk/guard/pkg/transport/http/profiler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Run starts HTTP server.
var Run = &cobra.Command{
	Use:   "run",
	Short: "Start HTTP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		runtime.GOMAXPROCS(int(cnf.Union.ServerConfig.CPUCount))
		var (
			repository = storage.Must(
				storage.Database(cnf.Union.DatabaseConfig),
			)
			service = guard.New(cnf.Union.ServiceConfig, repository)
		)
		if err := startGRPCServer(cnf.Union.GRPCConfig, repository); err != nil {
			return err
		}
		if err := startMonitoring(cnf.Union.MonitoringConfig); err != nil {
			return err
		}
		if err := startProfiler(cnf.Union.ProfilingConfig); err != nil {
			return err
		}
		return startHTTPServer(cnf.Union.ServerConfig, service)
	},
}

func init() {
	v := viper.New()
	fn.Must(
		func() error { return v.BindEnv("max_cpus") },
		func() error { return v.BindEnv("bind") },
		func() error { return v.BindEnv("http_port") },
		func() error { return v.BindEnv("profiling_port") },
		func() error { return v.BindEnv("monitoring_port") },
		func() error { return v.BindEnv("grpc_port") },
		func() error { return v.BindEnv("grpc_gateway_port") },
		func() error { return v.BindEnv("read_timeout") },
		func() error { return v.BindEnv("read_header_timeout") },
		func() error { return v.BindEnv("write_timeout") },
		func() error { return v.BindEnv("idle_timeout") },
		func() error {
			v.SetDefault("max_cpus", defaults["max_cpus"])
			v.SetDefault("bind", defaults["bind"])
			v.SetDefault("http_port", defaults["http_port"])
			v.SetDefault("profiling_port", defaults["profiling_port"])
			v.SetDefault("monitoring_port", defaults["monitoring_port"])
			v.SetDefault("grpc_port", defaults["grpc_port"])
			v.SetDefault("grpc_gateway_port", defaults["grpc_gateway_port"])

			bind := v.GetString("bind")
			v.SetDefault("host", strings.Concat(bind, ":", strconv.Itoa(v.GetInt("http_port"))))
			v.SetDefault("profiling_host", strings.Concat(bind, ":", strconv.Itoa(v.GetInt("profiling_port"))))
			v.SetDefault("monitoring_host", strings.Concat(bind, ":", strconv.Itoa(v.GetInt("monitoring_port"))))
			v.SetDefault("grpc_host", strings.Concat(bind, ":", strconv.Itoa(v.GetInt("grpc_port"))))
			v.SetDefault("grpc_gateway_host", strings.Concat(bind, ":", strconv.Itoa(v.GetInt("grpc_gateway_port"))))

			v.SetDefault("read_timeout", defaults["read_timeout"])
			v.SetDefault("read_header_timeout", defaults["read_header_timeout"])
			v.SetDefault("write_timeout", defaults["write_timeout"])
			v.SetDefault("idle_timeout", defaults["idle_timeout"])
			return nil
		},
		func() error {
			flags := Run.Flags()
			flags.UintVarP(&cnf.Union.ServerConfig.CPUCount,
				"cpus", "C", uint(v.GetInt("max_cpus")), "maximum number of CPUs that can be executing simultaneously")
			flags.StringVarP(&cnf.Union.ServerConfig.Interface,
				"host", "H", v.GetString("host"), "web server host")
			flags.DurationVarP(&cnf.Union.ServerConfig.ReadTimeout,
				"read-timeout", "", v.GetDuration("read_timeout"),
				"maximum duration for reading the entire request, including the body")
			flags.DurationVarP(&cnf.Union.ServerConfig.ReadHeaderTimeout,
				"read-header-timeout", "", v.GetDuration("read_header_timeout"),
				"amount of time allowed to read request headers")
			flags.DurationVarP(&cnf.Union.ServerConfig.WriteTimeout,
				"write-timeout", "", v.GetDuration("write_timeout"),
				"maximum duration before timing out writes of the response")
			flags.DurationVarP(&cnf.Union.ServerConfig.IdleTimeout,
				"idle-timeout", "", v.GetDuration("idle_timeout"),
				"maximum amount of time to wait for the next request when keep-alive is enabled")
			flags.BoolVarP(&cnf.Union.ProfilingConfig.Enabled,
				"with-profiling", "", false, "enable pprof on /pprof/* and /debug/pprof/")
			flags.StringVarP(&cnf.Union.ProfilingConfig.Interface,
				"profiling-host", "", v.GetString("profiling_host"), "profiling host")
			flags.BoolVarP(&cnf.Union.MonitoringConfig.Enabled,
				"with-monitoring", "", false, "enable prometheus on /monitoring and expvar on /vars")
			flags.StringVarP(&cnf.Union.MonitoringConfig.Interface,
				"monitoring-host", "", v.GetString("monitoring_host"), "monitoring host")
			flags.StringVarP(&cnf.Union.GRPCConfig.Interface,
				"grpc-host", "", v.GetString("grpc_host"), "gRPC server host")
			flags.BoolVarP(&cnf.Union.GRPCConfig.Gateway.Enabled,
				"with-grpc-gateway", "", false, "enable RESTful JSON API above gRPC")
			flags.StringVarP(&cnf.Union.GRPCConfig.Gateway.Interface,
				"grpc-gateway-host", "", v.GetString("grpc_gateway_host"), "gRPC gateway server host")
			flags.BoolVarP(&cnf.Union.ServiceConfig.Disabled,
				"disabled", "", false, "disable any service barriers, only logging")
			return nil
		},
	)
	db(Run)
}

func startHTTPServer(cnf config.ServerConfig, service *guard.Guard) error {
	listener, err := net.Listen("tcp", cnf.Interface)
	if err != nil {
		return err
	}
	log.Println("start HTTP server at", listener.Addr())
	return http.New(cnf, service).Serve(listener)
}

func startGRPCServer(cnf config.GRPCConfig, repository *storage.Storage) error {
	listener, err := net.Listen("tcp", cnf.Interface)
	if err != nil {
		return err
	}
	gateway, cascade := net.Listener(nil), make(chan struct{})
	if cnf.Gateway.Enabled {
		gateway, err = net.Listen("tcp", cnf.Gateway.Interface)
		if err != nil {
			return err
		}
	}
	go func(listener net.Listener) {
		close(cascade)
		log.Println("start gRPC server at", listener.Addr())
		_ = grpc.New(cnf, repository, repository).Serve(listener)
	}(listener)
	go func(listener net.Listener) {
		if listener == nil {
			return
		}
		<-cascade
		log.Println("start gRPC gateway server at", listener.Addr())
		_ = grpc.Gateway(cnf).Serve(listener)
	}(gateway)
	return nil
}

func startMonitoring(cnf config.MonitoringConfig) error {
	if cnf.Enabled {
		listener, err := net.Listen("tcp", cnf.Interface)
		if err != nil {
			return err
		}
		go func(listener net.Listener) {
			log.Println("start monitoring server at", listener.Addr())
			_ = monitor.New(cnf).Serve(listener)
		}(listener)
	}
	return nil
}

func startProfiler(cnf config.ProfilingConfig) error {
	if cnf.Enabled {
		listener, err := net.Listen("tcp", cnf.Interface)
		if err != nil {
			return err
		}
		go func(listener net.Listener) {
			log.Println("start profiling server at", listener.Addr())
			_ = profiler.New(cnf).Serve(listener)
		}(listener)
	}
	return nil
}
