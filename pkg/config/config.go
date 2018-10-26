package config

import (
	"net/url"
	"time"

	"github.com/kamilsk/go-kit/pkg/config"
)

// ApplicationConfig holds all configurations of the application.
type ApplicationConfig struct {
	Union struct {
		DatabaseConfig   `json:"db"         xml:"db"         yaml:"db"`
		GRPCConfig       `json:"grpc"       xml:"grpc"       yaml:"grpc"`
		MigrationConfig  `json:"migration"  xml:"migration"  yaml:"migration"`
		MonitoringConfig `json:"monitoring" xml:"monitoring" yaml:"monitoring"`
		ProfilingConfig  `json:"profiling"  xml:"profiling"  yaml:"profiling"`
		ServerConfig     `json:"server"     xml:"server"     yaml:"server"`
		ServiceConfig    `json:"service"    xml:"service"    yaml:"service"`
	} `json:"config" xml:"config" yaml:"config"`
}

// DatabaseConfig contains configuration related to a database.
type DatabaseConfig struct {
	DSN             config.Secret `json:"dsn"      xml:"dsn"      yaml:"dsn"`
	MaxIdleConns    int           `json:"idle"     xml:"idle"     yaml:"idle"`
	MaxOpenConns    int           `json:"open"     xml:"open"     yaml:"open"`
	ConnMaxLifetime time.Duration `json:"lifetime" xml:"lifetime" yaml:"lifetime"`

	dsn *url.URL
}

// DriverName returns database driver name.
// Error ignored, panic is possible (nil pointer).
// Not thread-safe call.
func (cnf *DatabaseConfig) DriverName() string {
	if cnf.dsn == nil {
		cnf.dsn, _ = url.Parse(string(cnf.DSN))
	}
	return cnf.dsn.Scheme
}

// GRPCConfig contains configuration related to the application gRPC server.
type GRPCConfig struct {
	RPC     RPCConfig     `json:"rpc"     xml:"rpc"     yaml:"rpc"`
	Gateway GatewayConfig `json:"gateway" xml:"gateway" yaml:"gateway"`
}

// RPCConfig contains configuration related to the application RPC server.
type RPCConfig struct {
	Interface string        `json:"interface" xml:"interface" yaml:"interface"`
	Timeout   time.Duration `json:"timeout"   xml:"timeout"   yaml:"timeout"`
	Token     config.Secret `json:"token"     xml:"token"     yaml:"token"`
}

// GatewayConfig contains configuration related to the application REST JSON API server above RPC server.
type GatewayConfig struct {
	Enabled           bool          `json:"enabled"             xml:"enabled"             yaml:"enabled"`
	Interface         string        `json:"interface"           xml:"interface"           yaml:"interface"`
	ReadTimeout       time.Duration `json:"read-timeout"        xml:"read-timeout"        yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" xml:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout"       xml:"write-timeout"       yaml:"write-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout"        xml:"idle-timeout"        yaml:"idle-timeout"`
}

// MigrationConfig contains configuration related to the application migrations.
type MigrationConfig struct {
	Table  string `json:"table"     xml:"table"     yaml:"table"`
	Schema string `json:"schema"    xml:"schema"    yaml:"schema"`
	Limit  uint   `json:"limit"     xml:"limit"     yaml:"limit"`
	DryRun bool   `json:"dry-run"   xml:"dry-run"   yaml:"dry-run"`
}

// MonitoringConfig contains configuration related to monitoring the application.
type MonitoringConfig struct {
	Enabled   bool   `json:"enabled"   xml:"enabled"   yaml:"enabled"`
	Interface string `json:"interface" xml:"interface" yaml:"interface"`
}

// ProfilingConfig contains configuration related to profiling the application.
type ProfilingConfig struct {
	Enabled   bool   `json:"enabled"   xml:"enabled"   yaml:"enabled"`
	Interface string `json:"interface" xml:"interface" yaml:"interface"`
}

// ServerConfig contains configuration related to the application web server.
type ServerConfig struct {
	Interface         string        `json:"interface"           xml:"interface"           yaml:"interface"`
	CPUCount          uint          `json:"cpus"                xml:"cpus"                yaml:"cpus"`
	ReadTimeout       time.Duration `json:"read-timeout"        xml:"read-timeout"        yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" xml:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout"       xml:"write-timeout"       yaml:"write-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout"        xml:"idle-timeout"        yaml:"idle-timeout"`
}

// ServiceConfig contains configuration related to the main application service.
type ServiceConfig struct {
	Disabled bool `json:"disabled" xml:"disabled" yaml:"disabled"`
}

/*
 *
 * TODO issue#future
 *
 */

// EtcdConfig contains configuration related to a Etcd cluster.
type EtcdConfig struct{}

// RedisConfig contains configuration related to a Redis cluster.
type RedisConfig struct{}
