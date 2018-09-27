package guard

import "github.com/kamilsk/guard/pkg/service/request"

// New TODO issue#docs
func New(storage Storage) *Guard {
	return &Guard{&licenseManager{storage}}
}

// Guard TODO issue#docs
type Guard struct {
	licenseManager *licenseManager
}

// CheckLicense TODO issue#docs
func (service Guard) CheckLicense(license request.License) error {
	return service.licenseManager.Check(license)
}
