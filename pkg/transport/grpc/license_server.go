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

// Create TODO issue#docs
func (server *licenseServer) Create(ctx context.Context, req *CreateLicenseRequest) (*CreateLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if _, err := server.storage.CreateLicense(ctx, token, query.CreateLicense{}); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token: %s; something happen: %v", token, err)
	}
	log.Printf("LicenseServer.Create was called with token %q\n", token)
	return &CreateLicenseResponse{}, nil
}

// Read TODO issue#docs
func (server *licenseServer) Read(ctx context.Context, req *ReadLicenseRequest) (*ReadLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if _, err := server.storage.ReadLicense(ctx, token, query.ReadLicense{}); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token: %s; something happen: %v", token, err)
	}
	log.Printf("LicenseServer.Read was called with token %q\n", token)
	return &ReadLicenseResponse{}, nil
}

// Update TODO issue#docs
func (server *licenseServer) Update(ctx context.Context, req *UpdateLicenseRequest) (*UpdateLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if _, err := server.storage.UpdateLicense(ctx, token, query.UpdateLicense{}); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token: %s; something happen: %v", token, err)
	}
	log.Printf("LicenseServer.Update was called with token %q\n", token)
	return &UpdateLicenseResponse{}, nil
}

// Delete TODO issue#docs
func (server *licenseServer) Delete(ctx context.Context, req *DeleteLicenseRequest) (*DeleteLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if _, err := server.storage.DeleteLicense(ctx, token, query.DeleteLicense{}); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token: %s; something happen: %v", token, err)
	}
	log.Printf("LicenseServer.Update was called with token %q\n", token)
	return &DeleteLicenseResponse{}, nil
}

// ---

// Register TODO issue#docs
func (server *licenseServer) Register(ctx context.Context, req *RegisterLicenseRequest) (*RegisterLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if _, err := server.storage.RegisterLicense(ctx, token, query.RegisterLicense{}); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token: %s; something happen: %v", token, err)
	}
	log.Printf("LicenseServer.Register was called with token %q\n", token)
	return &RegisterLicenseResponse{}, nil
}
