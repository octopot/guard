package guard

import (
	"context"

	"go.octolab.org/ecosystem/guard/internal/config"
	"go.octolab.org/ecosystem/guard/internal/service/guard/internal"
	"go.octolab.org/ecosystem/guard/internal/service/types/request"
	"go.octolab.org/ecosystem/guard/internal/service/types/response"
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
