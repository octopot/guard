package types

import (
	"time"

	"github.com/kamilsk/guard/pkg/service/types"
)

// Account TODO issue#docs
type Account struct {
	ID        types.ID   `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Users     []*User    `db:"-"`
}
