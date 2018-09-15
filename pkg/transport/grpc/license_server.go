package grpc

import (
	"context"
	"log"

	"github.com/kamilsk/go-kit/pkg/service/types"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewLicenseServer TODO issue#docs
func NewLicenseServer() LicenseServer {
	return &licenseServer{}
}

type licenseServer struct{}

// Extend TODO issue#docs
func (*licenseServer) Extend(ctx context.Context, req *ExtendLicenseRequest) (*ExtendLicenseResponse, error) {
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	if tokenID != types.ID(secret) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %s", tokenID)
	}
	log.Printf("LicenseServer.Extend was called with token %q\n", tokenID)
	return &ExtendLicenseResponse{}, nil
}

// Read TODO issue#docs
func (*licenseServer) Read(ctx context.Context, req *ReadLicenseRequest) (*ReadLicenseResponse, error) {
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	if tokenID != types.ID(secret) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %s", tokenID)
	}
	log.Printf("LicenseServer.Read was called with token %q\n", tokenID)
	return &ReadLicenseResponse{}, nil
}

// Register TODO issue#docs
func (*licenseServer) Register(ctx context.Context, req *RegisterLicenseRequest) (*RegisterLicenseResponse, error) {
	tokenID, err := middleware.TokenExtractor(ctx)
	if err != nil {
		return nil, err
	}
	if tokenID != types.ID(secret) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %s", tokenID)
	}
	log.Printf("LicenseServer.Register was called with token %q\n", tokenID)
	return &RegisterLicenseResponse{}, nil
}
