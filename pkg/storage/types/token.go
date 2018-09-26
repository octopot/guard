package types

import (
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

// Token TODO issue#docs
type Token struct {
	ID        domain.Token `db:"id"`
	UserID    domain.ID    `db:"user_id"`
	ExpiredAt *time.Time   `db:"expired_at"`
	CreatedAt time.Time    `db:"created_at"`
	User      *User        `db:"-"`
}
