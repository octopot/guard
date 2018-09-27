package types

import (
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

// Action TODO issue#docs
type Action string

// Create, Update, Delete defines Action set.
const (
	Create Action = "create"
	Update Action = "update"
	Delete Action = "delete"
)

// LicenseAudit TODO issue#docs
type LicenseAudit struct {
	ID       uint64          `db:"id"`
	Number   domain.ID       `db:"number"`
	Contract domain.Contract `db:"contract"`
	What     Action          `db:"what"`
	Who      domain.ID       `db:"who"`
	When     time.Time       `db:"when"`
	With     domain.Token    `db:"with"`
}
