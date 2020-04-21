package types

import (
	"time"

	domain "go.octolab.org/ecosystem/guard/internal/service/types"
)

// User TODO issue#docs
type User struct {
	ID        domain.ID  `db:"id"`
	AccountID domain.ID  `db:"account_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
	Account   *Account   `db:"-"`
	Tokens    []*Token   `db:"-"`
}
