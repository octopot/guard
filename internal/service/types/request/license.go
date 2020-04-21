package request

import domain "go.octolab.org/ecosystem/guard/internal/service/types"

// CheckLicense TODO issue#docs
type CheckLicense struct {
	License   domain.ID
	Employee  domain.ID
	Workplace domain.ID
}
