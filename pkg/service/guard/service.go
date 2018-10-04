package guard

import (
	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/service/types/request"
	"github.com/kamilsk/guard/pkg/service/types/response"
)

// New TODO issue#docs
func New(cnf config.ServiceConfig, storage Storage) *Guard {
	return &Guard{&licenseManager{cnf.Disabled, storage}}
}

// Guard TODO issue#docs
type Guard struct {
	licenseManager *licenseManager
}

// CheckLicense TODO issue#docs
func (service Guard) CheckLicense(request request.License) response.License {
	return service.licenseManager.Check(request)
}
