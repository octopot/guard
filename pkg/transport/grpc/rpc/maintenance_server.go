package rpc

import (
	"context"

	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/storage/query"
	"github.com/kamilsk/guard/pkg/transport/grpc/protobuf"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
func (server *maintenanceServer) Install(ctx context.Context, req *protobuf.InstallRequest) (*protobuf.InstallResponse, error) {
	type tokenSetter func(token *query.RegisterToken) *query.RegisterUser
	setTokens := func(tokens []*protobuf.InstallRequest_Token, set tokenSetter) {
		for _, token := range tokens {
			_ = set(&query.RegisterToken{
				ID: ptrToToken(token.Id),
			})
		}
	}

	type userSetter func(user *query.RegisterUser) *query.RegisterAccount
	setUsers := func(users []*protobuf.InstallRequest_User, set userSetter) {
		for _, user := range users {
			q := &query.RegisterUser{
				ID:   ptrToID(user.Id),
				Name: user.Name,
			}
			setTokens(user.Tokens, q.AddToken)
			_ = set(q)
		}
	}

	setAccount := func(account *protobuf.InstallRequest_Account) *query.RegisterAccount {
		q := &query.RegisterAccount{
			ID:   ptrToID(account.Id),
			Name: account.Name,
		}
		setUsers(account.Users, q.AddUser)
		return q
	}

	resp := server.service.Install(ctx, request.Install{Account: setAccount(req.Account)})
	if resp.HasError() {
		return nil, status.Errorf(codes.Internal, "something happen: %v", errors.Cause(&resp)) // TODO issue#6
	}

	convertTokens := func(tokens []*repository.Token) []*protobuf.InstallResponse_Token {
		out := make([]*protobuf.InstallResponse_Token, 0, len(tokens))
		for _, token := range tokens {
			out = append(out, &protobuf.InstallResponse_Token{
				Id:        token.ID.String(),
				Revoked:   token.Revoked,
				ExpiredAt: Timestamp(token.ExpiredAt),
				CreatedAt: Timestamp(&token.CreatedAt),
				UpdatedAt: Timestamp(token.UpdatedAt),
			})
		}
		return out
	}

	convertUsers := func(users []*repository.User) []*protobuf.InstallResponse_User {
		out := make([]*protobuf.InstallResponse_User, 0, len(users))
		for _, user := range users {
			out = append(out, &protobuf.InstallResponse_User{
				Id:        user.ID.String(),
				Name:      user.Name,
				CreatedAt: Timestamp(&user.CreatedAt),
				UpdatedAt: Timestamp(user.UpdatedAt),
				DeletedAt: Timestamp(user.DeletedAt),
				Tokens:    convertTokens(user.Tokens),
			})
		}
		return out
	}

	account := resp.Account
	return &protobuf.InstallResponse{
		Account: &protobuf.InstallResponse_Account{
			Id:        account.ID.String(),
			Name:      account.Name,
			CreatedAt: Timestamp(&account.CreatedAt),
			UpdatedAt: Timestamp(account.UpdatedAt),
			DeletedAt: Timestamp(account.DeletedAt),
			Users:     convertUsers(account.Users),
		},
	}, nil
}
