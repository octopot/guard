package grpc

import "context"

// NewLicenseServer TODO issue#docs
func NewLicenseServer() LicenseServer {
	return &licenseServer{}
}

type licenseServer struct{}

// Register TODO issue#docs
func (*licenseServer) Register(context.Context, *RegisterLicenseRequest) (*RegisterLicenseResponse, error) {
	return &RegisterLicenseResponse{}, nil
}

// Extend TODO issue#docs
func (*licenseServer) Extend(context.Context, *ExtendLicenseRequest) (*ExtendLicenseResponse, error) {
	return &ExtendLicenseResponse{}, nil
}

// Check TODO issue#docs
func (*licenseServer) Check(context.Context, *CheckLicenseRequest) (*CheckLicenseResponse, error) {
	return &CheckLicenseResponse{}, nil
}
