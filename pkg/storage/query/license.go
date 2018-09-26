package query

import domain "github.com/kamilsk/guard/pkg/service/types"

// ExtendLicense TODO issue#docs
type ExtendLicense struct {
	Number domain.ID
}

// ReadLicense TODO issue#docs
type ReadLicense struct {
	Number domain.ID
}

// RegisterLicense TODO issue#docs
type RegisterLicense struct {
	Number domain.ID
}
