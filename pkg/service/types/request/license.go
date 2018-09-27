package request

import domain "github.com/kamilsk/guard/pkg/service/types"

// License TODO issue#docs
type License struct {
	Number    domain.ID
	Employee  domain.ID
	Workplace domain.ID
	Metadata
}
