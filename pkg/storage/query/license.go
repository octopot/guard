package query

import "github.com/kamilsk/guard/pkg/service/types"

// ExtendLicense TODO issue#docs
type ExtendLicense struct {
	Number types.ID
}

// ReadLicense TODO issue#docs
type ReadLicense struct {
	Number types.ID
}

// RegisterLicense TODO issue#docs
type RegisterLicense struct {
	Number types.ID
}
