package middleware

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"

	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// AuthHeader defines authorization header.
	AuthHeader = "authorization"
	// AuthScheme defines authorization scheme.
	AuthScheme = "bearer"
)

// TokenInjector TODO issue#docs
func TokenInjector(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, AuthScheme)
	if err != nil {
		return nil, err
	}
	value := domain.Token(token)
	if !value.IsValid() {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %s", token)
	}
	return context.WithValue(ctx, tokenKey{}, value), nil
}

// TokenExtractor TODO issue#docs
func TokenExtractor(ctx context.Context) (domain.Token, error) {
	value, found := ctx.Value(tokenKey{}).(domain.Token)
	if !found {
		return value, status.Error(codes.Unauthenticated, "auth token not found")
	}
	return value, nil
}
