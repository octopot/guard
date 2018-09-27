package query

import domain "github.com/kamilsk/guard/pkg/service/types"

// CreateLicense TODO issue#docs
type CreateLicense struct {
	Number   *domain.ID
	Contract domain.Contract
}

// DeleteLicense TODO issue#docs
type DeleteLicense struct {
	Number domain.ID
}

// ReadLicense TODO issue#docs
type ReadLicense struct {
	Number domain.ID
}

// UpdateLicense TODO issue#docs
type UpdateLicense struct {
	Number   domain.ID
	Contract domain.Contract
}
