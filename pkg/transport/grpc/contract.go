package grpc

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// Maintenance TODO issue#docs
type Maintenance interface {
	// RegisterAccount TODO issue#docs
	RegisterAccount(context.Context, *query.RegisterAccount) (repository.Account, error)
}

// ProtectedStorage TODO issue#docs
type ProtectedStorage interface {
	// RegisterLicense TODO issue#docs
	RegisterLicense(context.Context, domain.Token, query.RegisterLicense) (repository.License, error)
	// CreateLicense TODO issue#docs
	CreateLicense(context.Context, domain.Token, query.CreateLicense) (repository.License, error)
	// ReadLicense TODO issue#docs
	ReadLicense(context.Context, domain.Token, query.ReadLicense) (repository.License, error)
	// UpdateLicense TODO issue#docs
	UpdateLicense(context.Context, domain.Token, query.UpdateLicense) (repository.License, error)
	// DeleteLicense TODO issue#docs
	DeleteLicense(context.Context, domain.Token, query.DeleteLicense) (repository.License, error)
	// RestoreLicense TODO issue#docs
	RestoreLicense(context.Context, domain.Token, query.RestoreLicense) (repository.License, error)
}
