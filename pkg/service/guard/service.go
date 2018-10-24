package guard

import (
	"context"

	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/service/guard/internal"
	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/service/types/response"
)

// New TODO issue#docs
func New(cnf config.ServiceConfig, storage Storage) *Guard {
	return &Guard{
		&licenseService{cnf.Disabled, internal.NewLicenseCache(storage)},
		&maintenanceService{storage},
	}
}

// Guard TODO issue#docs
type Guard struct {
	license     *licenseService
	maintenance *maintenanceService
}

// CheckLicense TODO issue#docs
func (service Guard) CheckLicense(ctx context.Context, req request.CheckLicense) response.CheckLicense {
	return service.license.Check(ctx, req)
}

// Install TODO issue#docs
func (service Guard) Install(ctx context.Context, req request.Install) response.Install {
	return service.maintenance.Install(ctx, req)
}
