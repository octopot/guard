package grpc

import (
	"context"
	"log"

	"github.com/kamilsk/guard/pkg/storage/query"
	"github.com/kamilsk/guard/pkg/transport/grpc/middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewLicenseServer TODO issue#docs
func NewLicenseServer(storage ProtectedStorage) LicenseServer {
	return &licenseServer{storage}
}

type licenseServer struct {
	storage ProtectedStorage
}

// Extend TODO issue#docs
func (server *licenseServer) Extend(ctx context.Context, req *ExtendLicenseRequest) (*ExtendLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if _, err := server.storage.ExtendLicense(ctx, token, query.ExtendLicense{}); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %s", token)
	}
	log.Printf("LicenseServer.Extend was called with token %q\n", token)
	return &ExtendLicenseResponse{}, nil
}

// Read TODO issue#docs
func (server *licenseServer) Read(ctx context.Context, req *ReadLicenseRequest) (*ReadLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if _, err := server.storage.ReadLicense(ctx, token, query.ReadLicense{}); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %s", token)
	}
	log.Printf("LicenseServer.Read was called with token %q\n", token)
	return &ReadLicenseResponse{}, nil
}

// Register TODO issue#docs
func (server *licenseServer) Register(ctx context.Context, req *RegisterLicenseRequest) (*RegisterLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if _, err := server.storage.RegisterLicense(ctx, token, query.RegisterLicense{}); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %s", token)
	}
	log.Printf("LicenseServer.Register was called with token %q\n", token)
	return &RegisterLicenseResponse{}, nil
}
