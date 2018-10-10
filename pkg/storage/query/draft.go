package query

import domain "github.com/kamilsk/guard/pkg/service/types"

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
