package guard

import (
	"context"

	repository "github.com/kamilsk/guard/pkg/storage/types"

	"github.com/kamilsk/guard/pkg/storage/query"
)

// Storage TODO issue#docs
type Storage interface {
	// RegisterAccount TODO issue#docs
	RegisterAccount(context.Context, *query.RegisterAccount) (*repository.Account, error)
}
