package query

import domain "github.com/kamilsk/guard/pkg/service/types"

// CreateLicense TODO issue#docs
type CreateLicense struct {
	ID       *domain.ID
	Contract domain.Contract
}

// DeleteLicense TODO issue#docs
type DeleteLicense struct {
	ID domain.ID
}

// ReadLicense TODO issue#docs
type ReadLicense struct {
	ID domain.ID
}

// RegisterLicense TODO issue#docs
type RegisterLicense struct {
	ID       domain.ID
	Contract domain.Contract
}

// RestoreLicense TODO issue#docs
type RestoreLicense struct {
	ID domain.ID
}

// UpdateLicense TODO issue#docs
type UpdateLicense struct {
	ID       domain.ID
	Contract domain.Contract
}

// GetLicenseWithID TODO issue#docs
type GetLicenseWithID struct {
	ID domain.ID
}

// GetEmployeeLicense TODO issue#docs
type GetEmployeeLicense struct {
	Employee domain.ID
}

// TODO issue#draft {

// LicenseEmployee TODO issue#docs
type LicenseEmployee struct {
	ID       domain.ID
	Employee domain.ID
}

// LicenseWorkplace TODO issue#docs
type LicenseWorkplace struct {
	ID        domain.ID
	Workplace domain.ID
}

// issue#draft }

/*
 *
 * TODO issue#future
 *
 */

// ExtendLicense TODO issue#docs
type ExtendLicense struct {
	ID    domain.ID
	Patch interface{}
}
