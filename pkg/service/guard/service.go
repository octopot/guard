package guard

import "github.com/kamilsk/guard/pkg/service/types"

// New TODO issue#docs
func New() *Guard {
	return &Guard{&licenseManager{}}
}

// Guard TODO issue#docs
type Guard struct {
	licenseManager *licenseManager
}

// CheckLicense TODO issue#docs
func (service Guard) CheckLicense(license types.License) error {
	return service.licenseManager.Check(license)
}
