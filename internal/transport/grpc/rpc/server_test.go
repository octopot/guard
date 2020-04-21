//go:generate echo $PWD/$GOPACKAGE/$GOFILE
//go:generate mockgen -package rpc_test -destination mock_maintenance_test.go go.octolab.org/ecosystem/guard/internal/transport/grpc/rpc Maintenance
//go:generate mockgen -package rpc_test -destination mock_storage_test.go go.octolab.org/ecosystem/guard/internal/transport/grpc/rpc Storage
package rpc_test
