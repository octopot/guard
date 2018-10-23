package request

import domain "github.com/kamilsk/guard/pkg/service/types"

// CheckLicense TODO issue#docs
type CheckLicense struct {
	ID        domain.ID
	Employee  domain.ID
	Workplace domain.ID
}
