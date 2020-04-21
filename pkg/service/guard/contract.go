package guard

import (
	"context"

	domain "github.com/kamilsk/guard/pkg/service/types"
	"github.com/kamilsk/guard/pkg/storage/query"
	repository "github.com/kamilsk/guard/pkg/storage/types"
)

// Storage TODO issue#docs
type Storage interface {
	accountStorage
	licenseStorage
}

type accountStorage interface {
	// RegisterAccount TODO issue#docs
	RegisterAccount(context.Context, *query.RegisterAccount) (*repository.Account, error)
}

type licenseStorage interface {
	// LicenseByID TODO issue#docs
	LicenseByID(context.Context, domain.ID) (repository.License, error)
	// LicenseByEmployee TODO issue#docs
	LicenseByEmployee(context.Context, domain.ID) (repository.License, error)
}
