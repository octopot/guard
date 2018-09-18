package types

import (
	"time"

	"github.com/kamilsk/guard/pkg/service/types"
)

// User TODO issue#docs
type User struct {
	ID        types.ID   `db:"id"`
	AccountID types.ID   `db:"account_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Account   *Account   `db:"-"`
	Tokens    []*Token   `db:"-"`
}
