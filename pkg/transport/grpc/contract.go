package grpc

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// ProtectedStorage TODO issue#docs
type ProtectedStorage interface {
	// ExtendLicense TODO issue#docs
	ExtendLicense(context.Context, domain.Token, query.ExtendLicense) (repository.License, error)
	// ReadLicense TODO issue#docs
	ReadLicense(context.Context, domain.Token, query.ReadLicense) (repository.License, error)
	// RegisterLicense TODO issue#docs
	RegisterLicense(context.Context, domain.Token, query.RegisterLicense) (repository.License, error)
}
