package request

import domain "github.com/kamilsk/guard/pkg/service/types"

// License TODO issue#docs
type License struct {
	ID        domain.ID
	Employee  domain.ID
	Workplace domain.ID
}
