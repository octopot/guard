package types

import (
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

// Contract TODO issue#docs
type Contract struct {
	Since      time.Time `db:"since"`
	Until      time.Time `db:"until"`
	Workplaces uint      `db:"workplace_limits"`
	Limits
}

// Limits TODO issue#docs
type Limits struct {
	Rate    domain.Rate `db:"rate_limits"`
	Request uint        `db:"request_limits"`
}

// License TODO issue#docs
type License struct {
	Number domain.ID
}
