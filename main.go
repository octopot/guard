// +build !ctl

package main

import (
	"io"
	"os"

	"github.com/kamilsk/guard/cmd"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
	_ "github.com/mailru/easyjson"
	_ "go.uber.org/zap"

	_ "github.com/grpc-ecosystem/go-grpc-prometheus"
	_ "go.etcd.io/etcd/clientv3"
	_ "go.etcd.io/etcd/clientv3/clientv3util"
	_ "go.etcd.io/etcd/clientv3/concurrency"
	_ "go.etcd.io/etcd/clientv3/integration"
	_ "go.etcd.io/etcd/clientv3/leasing"
	_ "go.etcd.io/etcd/clientv3/mirror"
	_ "go.etcd.io/etcd/clientv3/namespace"
	_ "go.etcd.io/etcd/clientv3/naming"
	_ "go.etcd.io/etcd/clientv3/ordering"
	_ "go.etcd.io/etcd/clientv3/yaml"
)

var service cli = func(executor commander, output io.Writer, shutdown func(code int)) {
	defer func() {
		if r := recover(); r != nil {
			shutdown(failure)
		}
	}()
	executor.AddCommand(cmd.Completion, cmd.Run, cmd.Version)
	if err := executor.Execute(); err != nil {
		shutdown(failure)
	}
	shutdown(success)
}

func main() { service(&cobra.Command{Use: "guard", Short: "Guard Service"}, os.Stderr, os.Exit) }
