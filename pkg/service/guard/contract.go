package guard

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"
	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// Storage TODO issue#docs
type Storage interface {
	// LicenseByID TODO issue#docs
	LicenseByID(context.Context, domain.ID) (repository.License, error)
	// LicenseByEmployee TODO issue#docs
	LicenseByEmployee(context.Context, domain.ID) (repository.License, error)

	// RegisterAccount TODO issue#docs
	RegisterAccount(context.Context, *query.RegisterAccount) (*repository.Account, error)
}
