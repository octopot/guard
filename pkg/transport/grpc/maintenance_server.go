package grpc

import (
	"github.com/kamilsk/guard/pkg/storage/query"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewMaintenanceServer TODO issue#docs
func NewMaintenanceServer(service Maintenance) MaintenanceServer {
	return &maintenanceServer{service}
}

type maintenanceServer struct {
	service Maintenance
}

// AuthFuncOverride implements ServiceAuthFuncOverride interface
// and allows to ignore the need for a user access token.
func (server *maintenanceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	// for best accuracy: fullMethodName == /grpc.Maintenance/Install
	return ctx, nil
}

// Install TODO issue#docs
func (server *maintenanceServer) Install(ctx context.Context, req *InstallRequest) (*InstallResponse, error) {
	account, registerErr := server.service.RegisterAccount(ctx, query.RegisterAccount{})
	if registerErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", registerErr) // TODO issue#6
	}
	return &InstallResponse{
		Account: &InstallResponse_Account{
			Id:        account.ID.String(),
			Name:      account.Name,
			CreatedAt: Timestamp(&account.CreatedAt),
			UpdatedAt: Timestamp(account.UpdatedAt),
			DeletedAt: Timestamp(account.DeletedAt),
		},
	}, nil
}
