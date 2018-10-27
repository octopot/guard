//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package rpc_test -destination $PWD/pkg/transport/grpc/rpc/mock_maintenance_test.go github.com/kamilsk/guard/pkg/transport/grpc/rpc Maintenance
//go:generate mockgen -package rpc_test -destination $PWD/pkg/transport/grpc/rpc/mock_storage_test.go github.com/kamilsk/guard/pkg/transport/grpc/rpc Storage
package rpc_test
