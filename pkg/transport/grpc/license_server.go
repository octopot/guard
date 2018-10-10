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
func NewLicenseServer(storage ProtectedStorage, draft DraftStorage) LicenseServer {
	return &licenseServer{storage, draft}
}

type licenseServer struct {
	storage ProtectedStorage
	draft   DraftStorage
}

// Register TODO issue#docs
func (server *licenseServer) Register(ctx context.Context, req *RegisterLicenseRequest) (*RegisterLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, registerErr := server.storage.RegisterLicense(ctx, token, query.RegisterLicense{
		ID:       domain.ID(req.Id),
		Contract: convertToDomainContract(req.Contract),
	})
	if registerErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", registerErr) // TODO issue#6
	}
	return &RegisterLicenseResponse{Id: license.ID.String()}, nil
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
	return &UpdateLicenseResponse{Id: license.ID.String(), UpdatedAt: Timestamp(license.UpdatedAt)}, nil
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
	return &DeleteLicenseResponse{Id: license.ID.String(), DeletedAt: Timestamp(license.DeletedAt)}, nil
}

// Restore TODO issue#docs
func (server *licenseServer) Restore(ctx context.Context, req *RestoreLicenseRequest) (*RestoreLicenseResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	license, restoreErr := server.storage.RestoreLicense(ctx, token, query.RestoreLicense{ID: domain.ID(req.Id)})
	if restoreErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", restoreErr) // TODO issue#6
	}
	return &RestoreLicenseResponse{Id: license.ID.String(), UpdatedAt: Timestamp(license.UpdatedAt)}, nil
}

// TODO issue#draft {

// AddEmployee TODO issue#docs
func (server *licenseServer) AddEmployee(ctx context.Context, req *EmployeeRequest) (*EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if addErr := server.draft.AddEmployee(ctx, token, query.LicenseEmployee{
		ID:       domain.ID(req.Id),
		Employee: domain.ID(req.Employee),
	}); addErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", addErr) // TODO issue#6
	}
	return new(EmptyResponse), nil
}

// DeleteEmployee TODO issue#docs
func (server *licenseServer) DeleteEmployee(ctx context.Context, req *EmployeeRequest) (*EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if deleteErr := server.draft.DeleteEmployee(ctx, token, query.LicenseEmployee{
		ID:       domain.ID(req.Id),
		Employee: domain.ID(req.Employee),
	}); deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", deleteErr) // TODO issue#6
	}
	return new(EmptyResponse), nil
}

// AddWorkplace TODO issue#docs
func (server *licenseServer) AddWorkplace(ctx context.Context, req *WorkplaceRequest) (*EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if addErr := server.draft.AddWorkplace(ctx, token, query.LicenseWorkplace{
		ID:        domain.ID(req.Id),
		Workplace: domain.ID(req.Workplace),
	}); addErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", addErr) // TODO issue#6
	}
	return new(EmptyResponse), nil
}

// DeleteWorkplace TODO issue#docs
func (server *licenseServer) DeleteWorkplace(ctx context.Context, req *WorkplaceRequest) (*EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if deleteErr := server.draft.DeleteWorkplace(ctx, token, query.LicenseWorkplace{
		ID:        domain.ID(req.Id),
		Workplace: domain.ID(req.Workplace),
	}); deleteErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", deleteErr) // TODO issue#6
	}
	return new(EmptyResponse), nil
}

// PushWorkplace TODO issue#docs
func (server *licenseServer) PushWorkplace(ctx context.Context, req *WorkplaceRequest) (*EmptyResponse, error) {
	token, authErr := middleware.TokenExtractor(ctx)
	if authErr != nil {
		return nil, authErr
	}
	if pushErr := server.draft.PushWorkplace(ctx, token, query.LicenseWorkplace{
		ID:        domain.ID(req.Id),
		Workplace: domain.ID(req.Workplace),
	}); pushErr != nil {
		return nil, status.Errorf(codes.Internal, "something happen: %v", pushErr) // TODO issue#6
	}
	return new(EmptyResponse), nil
}

// issue#draft }
