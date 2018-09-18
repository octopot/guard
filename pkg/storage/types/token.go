package types

import (
	"time"

	"github.com/kamilsk/guard/pkg/service/types"
)

// Token TODO issue#docs
type Token struct {
	ID        types.ID   `db:"id"`
	UserID    types.ID   `db:"user_id"`
	ExpiredAt *time.Time `db:"expired_at"`
	CreatedAt time.Time  `db:"created_at"`
	User      *User      `db:"-"`
}
