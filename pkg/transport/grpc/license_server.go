package grpc

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"

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
	license, createErr := server.storage.CreateLicense(ctx, token, query.CreateLicense{
		ID:       ptrToID(req.Id),
		Contract: convertToDomainContract(req.Contract),
	})
	if createErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", createErr) // TODO issue#6
	}
	return &CreateLicenseResponse{Id: license.ID.String(), CreatedAt: Timestamp(&license.CreatedAt)}, nil
}

// Read TODO issue#docs
func (server *licenseServer) Read(ctx context.Context, req *ReadLicenseRequest) (*ReadLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, readErr := server.storage.ReadLicense(ctx, token, query.ReadLicense{ID: domain.ID(req.Id)})
	if readErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", readErr) // TODO issue#6
	}
	return &ReadLicenseResponse{
		Id:        license.ID.String(),
		Contract:  convertFromDomainContract(license.Contract),
		CreatedAt: Timestamp(&license.CreatedAt),
		UpdatedAt: Timestamp(license.UpdatedAt),
		DeletedAt: Timestamp(license.DeletedAt),
	}, nil
}

// Update TODO issue#docs
func (server *licenseServer) Update(ctx context.Context, req *UpdateLicenseRequest) (*UpdateLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, updateErr := server.storage.UpdateLicense(ctx, token, query.UpdateLicense{
		ID:       domain.ID(req.Id),
		Contract: convertToDomainContract(req.Contract),
	})
	if updateErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", updateErr) // TODO issue#6
	}
	return &UpdateLicenseResponse{UpdatedAt: Timestamp(license.UpdatedAt)}, nil
}

// Delete TODO issue#docs
func (server *licenseServer) Delete(ctx context.Context, req *DeleteLicenseRequest) (*DeleteLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, deleteErr := server.storage.DeleteLicense(ctx, token, query.DeleteLicense{ID: domain.ID(req.Id)})
	if deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", deleteErr) // TODO issue#6
	}
	return &DeleteLicenseResponse{DeletedAt: Timestamp(license.DeletedAt)}, nil
}

// ---

// Register TODO issue#docs
func (server *licenseServer) Register(ctx context.Context, req *RegisterLicenseRequest) (*RegisterLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	_, registerErr := server.storage.RegisterLicense(ctx, token, query.RegisterLicense{
		ID:       domain.ID(req.Id),
		Contract: convertToDomainContract(req.Contract),
	})
	if registerErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", registerErr) // TODO issue#6
	}
	return &RegisterLicenseResponse{}, nil
}
